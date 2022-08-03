package models

type SideDishGroup struct {
	Name       string     `json:"sideDish group name"`
	ID         int64      `json:"sideDish group id"`
	SideDishes []SideDish `json:"sideDishes,omitempty"`
}

type SideDish struct {
	Name           string          `json:"sideDish name"`
	ID             int64           `json:"sideDish id"`
	Number         string          `json:"sideDish number"`
	UStId          int64           `json:"-"`
	TaxValue       int64           `json:"sideDish tax value"`
	SideDishPrices []SideDishPrice `json:"sideDish prices,omitempty"`
}

type SideDishPrice struct {
	SizeOrPackage    string  `json:"sideDish size or package"`
	SizeOrPackageId  int64   `json:"sideDish size or package id"`
	Price            float64 `json:"sideDish price"`
	BottleDepositFee float64 `json:"bottle deposit fee"`
}
