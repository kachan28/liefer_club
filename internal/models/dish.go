package models

type DishGroup struct {
	Name   string `json:"dish group name"`
	ID     int64  `json:"dish group id"`
	Dishes []Dish `json:"dishes"`
}

type Dish struct {
	Name       string      `json:"dish name"`
	ID         int64       `json:"dish id"`
	Number     string      `json:"dish number"`
	UStId      int64       `json:"-"`
	TaxValue   int64       `json:"dish tax value"`
	DishPrices []DishPrice `json:"dish prices"`
}

type DishPrice struct {
	SizeOrPackage    string  `json:"dish size or package"`
	SizeOrPackageId  int64   `json:"dish size or package id"`
	Price            float64 `json:"dish price"`
	BottleDepositFee float64 `json:"bottle deposit fee"`
}
