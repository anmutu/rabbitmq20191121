package main

import (
	"fmt"
	"rabbitmq20181121/RabbitMq"
	"strconv"
	"time"
)

func main() {
	rabbitmq := RabbitMq.NewRabbitMQSimple("duQueueName191224")
	for i := 0; i < 100; i++ {
		rabbitmq.PublishSimple("hello du message" + strconv.Itoa(i) + "---来自work模式")
		time.Sleep(1 * time.Second)
		fmt.Printf("work模式，共产生了%d条消息\n", i)
	}
}
