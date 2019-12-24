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
		rabbitmq.PublishSimple("hello du message" + strconv.Itoa(i))
		time.Sleep(1 * time.Second)
		//fmt.Printf("生产了%s个消息", )
		fmt.Println(i)
	}

}
