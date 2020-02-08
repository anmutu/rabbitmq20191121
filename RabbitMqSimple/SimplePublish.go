package main

import (
	"fmt"
	"rabbitmq20181121/RabbitMq"
)

func main() {
	rabbitmq := RabbitMq.NewRabbitMQSimple("duQueueName1912161843")
	rabbitmq.PublishSimple("他是客，你是心上人。 ---来自simple模式")
	fmt.Println("发送成功！")
}
