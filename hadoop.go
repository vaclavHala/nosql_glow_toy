package main

import (
	"log"
	"strconv"
	"strings"

	"github.com/chrislusf/glow/flow"
	"github.com/chrislusf/glow/source/hdfs"
)

func HadoopProducts() chan Product {
	sink := make(chan Product, 100)
	go fetchHadoop(sink)
	return sink
}

func fetchHadoop(sink chan Product) {
	f := flow.New()

	hdfs.Source(
		f,
		"hdfs://localhost:12300/data",
		3,
	).Map(func(line string) {
		split := strings.FieldsFunc(line,
			func(c rune) bool { return c == ',' })

		if len(split) != 8 {
			log.Printf("Bad line: <%s> error <Has %d parts, expected 8>",
				line, len(split))
			return
		}

		var err error
		id, err := strconv.Atoi(split[0])
		ratingCount, err := strconv.Atoi(split[6])
		price, err := strconv.ParseFloat(split[3], 32)
		ratingVal, err := strconv.ParseFloat(split[5], 32)

		if err != nil {
			log.Printf("Bad line: <%s> error <%s>", line, err)
			return
		}

		p := Product{
			Id:   id,
			Name: split[1],
			Offer: Offer{Availability: split[2],
				Price:    float32(price),
				Currency: split[4]},
			Rating: Rating{
				Value: float32(ratingVal),
				Count: ratingCount},
			Description: split[7]}
		sink <- p
	}).Run()

	close(sink)
}
