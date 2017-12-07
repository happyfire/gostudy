package main

import (
	"fmt"
	"github.com/garyburd/redigo/redis"
	//"time"
	"github.com/happyfire/gostudy/redigo/redisClient"
	"time"
)

func main2() {
	c, err := redis.Dial("tcp", "127.0.0.1:6379")
	if err != nil {
		fmt.Println("Connect to redis error", err)
		return
	}
	defer c.Close()

	_, err = c.Do("AUTH", "dream2010")
	if err != nil {
		fmt.Println("AUTH failed:", err)
		return
	}

	reply, err := c.Do("SET", "mykey", "hello")
	if err != nil {
		fmt.Println("redis set failed:", err)
	} else {
		fmt.Printf("%v %T\n", reply, reply)
	}

	reply, err = redis.String(c.Do("GET", "mykey"))
	if err != nil {
		fmt.Println("redis get failed:", err)
	} else {
		fmt.Printf("Get mykey: %v\n", reply)
	}

	c.Send("SET", "foo", "bar")
	c.Send("GET", "foo")
	c.Flush()
	fmt.Println(c.Receive()) // reply from set
	v, err := c.Receive()    // reply from get
	fmt.Println(redis.String(v, err))

}

func main1() {
	var c redisClient.RedisClient
	err := c.ConnectTo("127.0.0.1:6379", "dream2010")
	if err != nil {
		fmt.Println("connect failed:", err)
	}
	defer c.Close()
	c.EnableAutoReconnect()

	var c2 redisClient.RedisClient
	err = c2.ConnectTo("192.168.1.19:6379", "")
	if err != nil {
		fmt.Println("c2 connect failed:", err)
	}
	defer c2.Close()

	//c.Mux.Lock()
	//c.Connection.Do("SET", "foo", "bar")
	//c.Mux.Unlock()

	for {
		var value string
		var err error
		if c.IsAvaliable {
			value, err = c.GetString("foo")
			if err != nil {
				value, err = c2.GetString("foo")
			}
		} else if c2.IsAvaliable {
			value, err = c2.GetString("foo")
		}

		if err != nil {
			panic("all redis down!")
		} else {
			fmt.Println("get value for foo:", value)
		}

		time.Sleep(1 * time.Second)
	}
}

func main() {
	var rm redisClient.RedisManager
	rm.SetRedisServer("127.0.0.1:6379", "dream2010", true)
	rm.SetRedisServer("192.168.1.19:6379", "", false)

	defer rm.Close()

	for {
		value, err := rm.GetString("foo")
		if err != nil {
			panic("all redis down!")
		} else {
			fmt.Println("get value:", value)
		}

		time.Sleep(1 * time.Second)
	}
}
