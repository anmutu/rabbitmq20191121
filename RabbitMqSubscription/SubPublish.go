package main

import (
	"fmt"
	"rabbitmq20181121/RabbitMq"
	"strconv"
	"time"
)

func main() {
	rabbitmq := RabbitMq.NewRabbitMqSubscription("duexchangeName")
	for i := 0; i < 100; i++ {
		rabbitmq.PublishSubscription("订阅模式生产第" + strconv.Itoa(i) + "条数据")
		fmt.Printf("订阅模式生产第" + strconv.Itoa(i) + "条数据\n")
		time.Sleep(1 * time.Second)
	}
}
