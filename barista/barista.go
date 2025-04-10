package barista

import (
	"fmt"
	"math/rand"
	"time"

	"github.com/Felix-kuang/coffee-concurrency/order"
)

func StartBarista(id int, orderChan <-chan order.Order) {
	for order := range orderChan {
		fmt.Printf("👨‍🍳 Barista %d mulai membuat kopi untuk %s(Order #%d)...\n", id, order.CustomerName, order.OrderID)
		sleepTime := time.Duration(rand.Intn(3)+1) * time.Second
		time.Sleep(sleepTime)
		fmt.Printf("✅ Barista %d: Kopi untuk %s(Order #%d) sudah selesai dibuat!\n", id, order.CustomerName, order.OrderID)
	}
}
