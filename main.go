package main

import (
	"fmt"
	"math/rand"
	"time"
)

type Order struct {
	CustomerName string
	OrderID      int
}

func barista(id int, orderChan <-chan Order) {
	for order := range orderChan {
		fmt.Printf("Barista %d mulai membuat kopi untuk %s(Order #%d)...\n", id, order.CustomerName, order.OrderID)
		sleepTime := time.Duration(rand.Intn(3)+1) * time.Second
		time.Sleep(sleepTime)
		fmt.Printf("Barista %d: Kopi untuk %s(Order #%d) sudah selesai dibuat!\n", id, order.CustomerName, order.OrderID)
	}
}

func main() {
	rand.New(rand.NewSource(time.Now().UnixMicro()))

	orderChan := make(chan Order)

	numBaristas := 3
	for i := 1; i <= numBaristas; i++ {
		go barista(i, orderChan)
	}

	customers := []string{"Aldi", "Budi", "Cici", "Dewi", "Eka", "Fajar", "Gita"}

	for i, name := range customers {
		fmt.Printf("%s mengorder kopi! (Order #%d)\n", name, i+1)
		orderChan <- Order{CustomerName: name, OrderID: i + 1}
		time.Sleep(500 * time.Millisecond)
	}

	close(orderChan)
	time.Sleep(5 * time.Second)
}
