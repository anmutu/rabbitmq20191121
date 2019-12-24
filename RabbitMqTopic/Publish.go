package main

import (
	"fmt"
	"rabbitmq20181121/RabbitMq"
	"strconv"
	"time"
)

func main() {
	one := RabbitMq.NewRabbitMqTopic("exchangeNameTpoic1224", "SuperStar.Jay")
	two := RabbitMq.NewRabbitMqTopic("exchangeNameTpoic1224", "Superstar.James")
	for i := 0; i < 100; i++ {
		one.PublishTopic("小杜同学，topic模式，Jay," + strconv.Itoa(i))
		two.PublishTopic("小杜同学，topic模式，All," + strconv.Itoa(i))
		time.Sleep(1 * time.Second)
		fmt.Printf("这是小杜同学发布的消息%v \n", i)
	}
}
