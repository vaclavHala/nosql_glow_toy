package main

import (
	"log"
"sync"

	"gopkg.in/mgo.v2"
)

func FetchMongo(sink chan Product, wait *sync.WaitGroup) {
	session, err := mgo.Dial("localhost")
	if err != nil {
		log.Fatal("Can't connect")
		close(sink)
		return
	}
	defer session.Close()

	// Optional. Switch the session to a monotonic behavior.
	session.SetMode(mgo.Monotonic, true)

	products := session.DB("eshop").C("products")
	iter := products.Find(nil).Iter()
	var p Product
	for iter.Next(&p) {
		sink <- p
	}
	wait.Done()
}
