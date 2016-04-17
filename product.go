package main

import (
	"fmt"
)

type (
	Product struct {
		Id          int    `bson:"_id"`
		Rating      Rating `bson:"aggregateRating"`
		Offer       Offer  `bson:"offers"`
		Description string `bson:"description"`
		Name        string `bson:"name"`
	}

	Rating struct {
		Value float32 `bson:"ratingValue"`
		Count int     `bson:"reviewCount"`
	}

	Offer struct {
		Availability string  `bson:"availability"`
		Price        float32 `bson:"price"`
		Currency     string  `bson:"priceCurrency"`
	}
)

func (p Product) String() string {
	return fmt.Sprintf("Product{id:%d, name:%s, offer:%s, rating:%s}",
		p.Id, p.Name, p.Offer, p.Rating)
}
func (o Offer) String() string {
	return fmt.Sprintf("Offer{availability:%s, price:%f, currency:%s}",
		o.Availability, o.Price, o.Currency)
}
func (r Rating) String() string {
	return fmt.Sprintf("Rating{rating:%f, count:%d}",
		r.Value, r.Count)
}
