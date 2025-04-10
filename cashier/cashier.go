package cashier

import (
	"fmt"
	"time"

	"github.com/Felix-kuang/coffee-concurrency/order"
)

func StartCashier(customers []string, orderChan chan<- order.Order) {
	go func() {
		for i, name := range customers {
			fmt.Printf("📦 %s mengorder kopi! (Order #%d)\n", name, i+1)
			orderChan <- order.Order{CustomerName: name, OrderID: i + 1}
			time.Sleep(500 * time.Millisecond)
		}
		close(orderChan)
	}()
}
