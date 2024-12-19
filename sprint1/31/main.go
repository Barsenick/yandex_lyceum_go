package main

import (
	"log"
)

type Order struct {
	OrderNumber  int
	CustomerName string
	OrderAmount  float64
}

type OrderLogger struct {
	log.Logger
}

func NewOrderLogger() OrderLogger {
	return OrderLogger{*log.Default()}
}

func (logger *OrderLogger) AddOrder(order Order) {
	logger.Printf("Добавлен заказ #%d, Имя клиента: %s, Сумма заказа: $%.2f", order.OrderNumber, order.CustomerName, order.OrderAmount)
}
