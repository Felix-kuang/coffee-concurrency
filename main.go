package main

import (
	"fmt"
	"github.com/Felix-kuang/coffee-concurrency/barista"
	"github.com/Felix-kuang/coffee-concurrency/cashier"
	"github.com/Felix-kuang/coffee-concurrency/order"
	"math/rand"
	"time"
)

func main() {
	rand.New(rand.NewSource(time.Now().UnixMicro()))

	orderChan := make(chan order.Order)
	doneChan := make(chan barista.DoneInfo)

	numBaristas := 3
	for i := 1; i <= numBaristas; i++ {
		go barista.StartBarista(i, orderChan, doneChan)
	}

	//ini kasir
	customers := []string{"Aldi", "Budi", "Cici", "Dewi", "Eka", "Fajar", "Gita"}

	cashier.StartCashier(customers, orderChan)
	totalOrders := 0
	baristaCount := make(map[int]int)
	var totalTime time.Duration

	for range customers {
		done := <-doneChan
		totalOrders++
		baristaCount[done.BaristaID]++
		totalTime += done.PrepTime
		fmt.Printf("📝 DoneInfo: Barista #%d nyelesain Order #%d dalam %.2f detik\n", done.BaristaID, done.Order.OrderID, done.PrepTime.Seconds())
	}

	fmt.Println("\n📊 Statistik Akhir: ")
	fmt.Printf("- Total order selesai: %d\n", totalOrders)
	fmt.Printf("- Rata rata waktu penyajian: %.2f detik\n", totalTime.Seconds()/float64(totalOrders))

	var topBaristaID, topCount int
	for id, count := range baristaCount {
		if count > topCount {
			topBaristaID = id
			topCount = count
		}
	}
	fmt.Printf("- Barista dengan order selesai terbanyak: %d (%d order)\n", topBaristaID, topCount)
}
