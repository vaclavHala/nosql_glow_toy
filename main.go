package main

import (
	"flag"
	"fmt"
	"math"

	"github.com/chrislusf/glow/flow"
)

type totCnt struct {
	total float32
	count int
}

func main() {
	flag.Parse()

	flow := flow.New().Channel(MongoProducts())
	flow.Map(func(p Product) (int, totCnt) {
		ratingRound := math.Floor(float64(p.Rating.Value) + 0.5)
		return int(ratingRound), totCnt{p.Offer.Price, 1}
	}).ReduceByKey(func(a totCnt, b totCnt) totCnt {
		return totCnt{a.total + b.total, a.count + b.count}
	}).Map(func(rating int, acc totCnt) (int, float32) {
		return rating, acc.total / float32(acc.count)
	}).Map(func(rating int, avg float32) {
		fmt.Printf("rating: %d average price: %f\n", rating, avg)
	})

	flow.Run()

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
