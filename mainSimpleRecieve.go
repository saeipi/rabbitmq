package main

import "rabbitmq/RabbitMQ"

func main() {
	rabbitmq := RabbitMQ.NewRabbitMQSimple("" + "mySimple")
	rabbitmq.ConsumeSimple()
}
