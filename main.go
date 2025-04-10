package main

import (
	"fmt"
	"github.com/Felix-kuang/coffee-concurrency/barista"
	"github.com/Felix-kuang/coffee-concurrency/order"
	"math/rand"
	"time"
)

func main() {
	rand.New(rand.NewSource(time.Now().UnixMicro()))

	orderChan := make(chan order.Order)

	numBaristas := 3
	for i := 1; i <= numBaristas; i++ {
		go barista.StartBarista(i, orderChan)
	}

	customers := []string{"Aldi", "Budi", "Cici", "Dewi", "Eka", "Fajar", "Gita"}

	for i, name := range customers {
		fmt.Printf("📦 %s mengorder kopi! (Order #%d)\n", name, i+1)
		orderChan <- order.Order{CustomerName: name, OrderID: i + 1}
		time.Sleep(500 * time.Millisecond)
	}

	close(orderChan)
	time.Sleep(5 * time.Second)
}
