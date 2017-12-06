package redisClient

import (
	"fmt"
	"github.com/garyburd/redigo/redis"
	"io"
	"sync"
	"time"
)

type RedisClient struct {
	Connection   redis.Conn
	IsAvaliable  bool
	Mux          sync.Mutex
	redisAddress string
	password     string
}

const (
	ConnectionTimeout = 1 * time.Second
	ReadTimeout       = 1 * time.Second
	WriteTimeout      = 1 * time.Second
	PingWaitTime      = 1 * time.Second
	KeepAliveTime     = 60 * time.Second
)

func (c *RedisClient) doConnect() error {
	var err error
	c.Connection, err = redis.Dial("tcp", c.redisAddress, redis.DialConnectTimeout(ConnectionTimeout),
		redis.DialReadTimeout(ReadTimeout), redis.DialWriteTimeout(WriteTimeout), redis.DialKeepAlive(KeepAliveTime))
	if err != nil {
		fmt.Println("Connect to redis error", err)
		c.Connection = nil
		return err
	}

	if len(c.password) > 0 {
		_, err = c.Connection.Do("AUTH", c.password)
		if err != nil {
			fmt.Println("AUTH failed:", err)
			c.Connection.Close()
			return err
		}
	}

	c.IsAvaliable = true
	return nil
}

func (c *RedisClient) ConnectTo(address, password string) error {
	c.redisAddress = address
	c.password = password

	err := c.doConnect()
	return err
}

func (c *RedisClient) EnableAutoReconnect() {
	go func() {
		for {
			if c.IsAvaliable {
				c.Mux.Lock()
				_, err := c.Connection.Do("PING")
				c.Mux.Unlock()
				if err != nil {
					c.IsAvaliable = false
					fmt.Println("ping error", err)
					c.Close()
					c.doConnect()
				}
			} else {
				c.doConnect()
			}

			time.Sleep(PingWaitTime)
		}
	}()
}

func (c *RedisClient) Close() {
	c.Mux.Lock()
	if c.Connection != nil {
		c.Connection.Close()
		c.Connection = nil
		c.IsAvaliable = false
	}
	c.Mux.Unlock()
}

func (c *RedisClient) GetString(key string) string {
	if c.IsAvaliable {
		c.Mux.Lock()
		result, err := c.Connection.Do("GET", key)
		c.Mux.Unlock()
		if err != nil {
			if err == io.EOF {
				c.Close()
			}
			return ""
		} else {
			str, err := redis.String(result, err)
			if err != nil {
				return ""
			} else {
				return str
			}
		}
	} else {
		return ""
	}
}
