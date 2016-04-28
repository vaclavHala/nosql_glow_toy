package main

import (
	"flag"
	"fmt"
	"math"
	"sync"

	_ "github.com/chrislusf/glow/driver"
	"github.com/chrislusf/glow/flow"
)

type totCnt struct {
	total float32
	count int
}

func main() {
	flag.Parse()
	sink := make(chan Product, 100)
	var wait sync.WaitGroup
	wait.Add(3)

	go FetchHadoop(sink, &wait)
	go FetchMongo(sink, &wait)
	go FetchPostgres(sink, &wait)
	go func (s chan Product, w *sync.WaitGroup){
		w.Wait()
		close(s)
	}(sink, &wait)
	f := flow.New().Channel(sink)

	flow.Ready()

	f.Map(func(p Product) (int, totCnt) {
		ratingRound := math.Floor(float64(p.Rating.Value) + 0.5)
		return int(ratingRound), totCnt{p.Offer.Price, 1}
	}).ReduceByKey(func(a totCnt, b totCnt) totCnt {
		return totCnt{a.total + b.total, a.count + b.count}
	}).Map(func(rating int, acc totCnt) string {
		avg := acc.total/float32(acc.count)
		fmt.Printf("rating: %d average price: %f count: %d\n",
			rating, avg, acc.count)
		return fmt.Sprintf("%d;%f;%d\n", rating, avg, acc.count)
	}).AddOutput(SendOverNet("localhost", 12345))

	f.Run()

	//        Map(func(line string, ch chan string) {
	// 	for _, token := range strings.Split(line, ":") {
	// 		ch <- token
	// 	}
	// }).Map(func(key string) int {
	// 	return 1
	// }).Reduce(func(x int, y int) int {
	// 	return x + y
	// }).Map(func(x int) {
	// 	println("count:", x)
	// }).Run()
}
