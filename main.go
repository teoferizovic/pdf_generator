package main

import (
	"fmt"
	"github.com/BurntSushi/toml"
	"github.com/go-redis/redis"
	"pdf_generator/model"
	"pdf_generator/processor"
)

var conf model.Config
var (
	RedisClient *redis.Client
)


func init(){

	if _, err := toml.DecodeFile("./config.toml", &conf); err != nil {
		fmt.Println(err)
	}

}

func main() {

	RedisClient := redis.NewClient(&redis.Options{
		Addr:     conf.RedisServer+":"+conf.RedisPort,
		Password: "", // no password set
		DB:       0,  // use default DB
	})

	pong, err := RedisClient.Ping().Result()
	fmt.Println(pong, err)

	psNewMessage := RedisClient.Subscribe(conf.RedisChannel1)

	for {
		msg, _ := psNewMessage.ReceiveMessage()

			//fmt.Println(msg.Payload)
			processor.GenerateMsg(msg.Payload)

	}

}
