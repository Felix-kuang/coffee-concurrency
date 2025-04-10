package barista

import (
	"fmt"
	"math/rand"
	"time"

	"github.com/Felix-kuang/coffee-concurrency/order"
)

type DoneInfo struct {
	BaristaID int
	Order     order.Order
	PrepTime  time.Duration
}

func StartBarista(id int, orderChan <-chan order.Order, doneChan chan<- DoneInfo) {
	for order := range orderChan {
		start := time.Now()

		fmt.Printf("🧑‍🍳 Barista #%d mulai membuat kopi untuk customer %s (Order #%d)...\n", id, order.CustomerName, order.OrderID)
		sleepTime := time.Duration(rand.Intn(3)+1) * time.Second
		time.Sleep(sleepTime)
		fmt.Printf("✅ Barista #%d: Kopi untuk customer %s (Order %d) udah selesai dibuat! \n", id, order.CustomerName, order.OrderID)

		doneChan <- DoneInfo{
			BaristaID: id,
			Order:     order,
			PrepTime:  time.Since(start),
		}
	}
}
