package main

import (
	"log"

	"database/sql"
	_ "github.com/lib/pq"
)

func PostgresProducts() chan Product {
	sink := make(chan Product, 100)
	go fetchPostgres(sink)
	return sink
}

func fetchPostgres(sink chan Product) {
	db, err := sql.Open("postgres",
		"host=localhost port=54321 dbname=eshop user=postgres")
	if err != nil {
		log.Fatal(err)
		return
	}
	defer db.Close()

	rows, err := db.Query(
		"SELECT id, name, description, price, availability, currency, rating, ratingCount FROM product;")
	if err != nil {
		log.Fatal(err)
		return
	}
	defer rows.Close()

	for rows.Next() {
		var id int
		var name string
		var description string
		var rating float32
		var ratingCount int
		var price float32
		var currency string
		var availability string
		err = rows.Scan(&id, &name, &description, &price, &availability, &currency, &rating, &ratingCount)
		if err != nil {
			log.Println("Error", err)
			continue
		}
		p := Product{
			Id:          id,
			Name:        name,
			Description: description,
			Offer: Offer{
				Availability: availability,
				Currency:     currency,
				Price:        price},
			Rating: Rating{
				Value: rating,
				Count: ratingCount}}
		sink <- p
	}
	close(sink)
}
