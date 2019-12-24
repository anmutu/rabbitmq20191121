package main

import "rabbitmq20181121/RabbitMq"

func main() {
	rabbitmq := RabbitMq.NewRabbitMqSubscription("duexchangeName")
	rabbitmq.ConsumeSbuscription()
}
