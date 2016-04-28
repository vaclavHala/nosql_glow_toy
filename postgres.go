package main

import (
	"log"
	"sync"

	"database/sql"

	_ "github.com/lib/pq"
)

func FetchPostgres(sink chan Product, wait *sync.WaitGroup) {
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
		p := Product{}
		rows.Scan(&p.Id, &p.Name, &p.Description,
			&p.Offer.Price, &p.Offer.Availability,
			&p.Offer.Currency,
			&p.Rating.Value, &p.Rating.Count)
		sink <- p
	}
	wait.Done()
}
