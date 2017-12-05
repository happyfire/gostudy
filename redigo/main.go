package main

import (
	"fmt"
	"github.com/garyburd/redigo/redis"
)

func main() {
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
