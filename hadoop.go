package main

import (
	"fmt"

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
		fmt.Println(line)
		// sink <- line
	}).Run()

	close(sink)
}
