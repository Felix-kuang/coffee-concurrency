package main

import (
	"fmt"
	"time"
)

type Order struct {
	CustomerName string
}

func barista(orderChan <-chan Order) {
	for order := range orderChan {
		fmt.Printf("Barista mulai membuat kopi untuk %s...\n", order.CustomerName)
		time.Sleep(2 * time.Second)
		fmt.Printf("Kopi untuk %s sudah selesai dibuat!\n", order.CustomerName)
	}
}

func main() {
	orderChan := make(chan Order)

	go barista(orderChan)

	customers := []string{"Jane Doe", "Aldi", "Budi"}

	for _, name := range customers {
		fmt.Printf("%s mengorder kopi!\n", name)
		orderChan <- Order{CustomerName: name}
		time.Sleep(1 * time.Second)
	}

	close(orderChan)
}
