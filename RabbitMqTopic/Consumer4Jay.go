package main

import "rabbitmq20181121/RabbitMq"

func main() {
	jay := RabbitMq.NewRabbitMqTopic("exchangeNameTpoic1224", "SuperStar.*")
	jay.ConsumerTopic()
}
