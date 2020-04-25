package main

import (
	"fmt"
	"github.com/gomodule/redigo/redis"
)

func main(){
	c,err := redis.Dial("tcp","127.0.0.1:6379")
	if err != nil{
		fmt.Println("redis connect fail")
	}
	defer c.Close()

	_,err = c.Do("SET","name","zhangsan")
	if err !=nil{
		fmt.Println("redis set failed")
	}

	username,err := redis.String(c.Do("GET","name"))
	if err != nil{
		fmt.Println("redis get failed")
	}else{
		fmt.Printf("get key:%v",username)
	}
}
