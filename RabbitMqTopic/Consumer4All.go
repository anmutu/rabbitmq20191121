package main

import "rabbitmq20181121/RabbitMq"

func main() {
	all := RabbitMq.NewRabbitMqTopic("exchangeNameTpoic1224", "#")
	all.ConsumerTopic()
}
