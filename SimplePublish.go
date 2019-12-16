package main

import (
	"fmt"
	"rabbitmq20181121/RabbitMq"
)

func main() {
	rabbitmq := RabbitMq.NewRabbitMQSimple("duQueueName")
	rabbitmq.ConsumeSimple()
	fmt.Println("发送成功！")
}
