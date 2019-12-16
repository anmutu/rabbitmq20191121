package main

import (
	"fmt"
	"rabbitmq20181121/RabbitMq"
)

func main() {
	rabbitmq := RabbitMq.NewRabbitMQSimple("duQueueName1912161843")
	rabbitmq.PublishSimple("杰伦的歌大约能治感冒")
	fmt.Println("发送成功！")
}
