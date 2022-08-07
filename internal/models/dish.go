package models

import (
	"fmt"
)

type DishGroup struct {
	Name   string `json:"dish group name"`
	ID     int64  `json:"dish group id"`
	Dishes []Dish `json:"dishes,omitempty"`
}

type Dish struct {
	Name       string      `json:"dish name"`
	ID         int64       `json:"dish id"`
	Number     string      `json:"dish number"`
	UStId      int64       `json:"-"`
	TaxValue   int64       `json:"dish tax value"`
	DishPrices []DishPrice `json:"dish prices,omitempty"`
}

type DishPrice struct {
	SizeOrPackage    *string  `json:"dish size or package,omitempty"`
	SizeOrPackageId  *int64   `json:"dish size or package id,omitempty"`
	Price            *float64 `json:"dish price,omitempty"`
	BottleDepositFee *float64 `json:"bottle deposit fee,omitempty"`
}

func (d Dish) ToString() string {
	return fmt.Sprintf("%s: Nr. - %s; MwSt. - %d; Preis - %s", d.Name, d.Number, d.TaxValue, d.pricesToString())
}

func (d Dish) pricesToString() string {
	pricesString := ""
	for _, price := range d.DishPrices {
		if price.SizeOrPackage != nil {
			pricesString += *price.SizeOrPackage + " - "
		}
		pricesString += fmt.Sprintf("%.2f€", *price.Price)
		pricesString += "; "
	}
	return pricesString
}
