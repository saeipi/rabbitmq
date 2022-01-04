package main

import (
	"fmt"
	"rabbitmq/RabbitMQ"
)

func main() {
	rabbitmq := RabbitMQ.NewRabbitMQSimple("" + "mySimple")
	rabbitmq.PublishSimple("Hello saeipi!")
	fmt.Println("发送成功！")
}
