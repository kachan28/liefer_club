package models

import "fmt"

type SpecialOfferAndSetMenu struct {
	Components []Component `json:"components of special offers and menus"`
	Offers     []Offer     `json:"special offers and menus dishes"`
}

type Component struct {
	ComponentName        string  `json:"component name"`
	ComponentId          int64   `json:"component id"`
	ComponentNumber      string  `json:"component number"`
	ConfiguringPrinciple string  `json:"configuring principle"`
	Quantity             int64   `json:"quantity"`
	PaidQuantity         int64   `json:"paid quantity"`
	PricingPrinciple     string  `json:"pricing principle"`
	PriceOrDiscount      float64 `json:"price or discount"`
	PercentageDiscount   float64 `json:"percentage discount"`
}

type Offer struct {
	DishName             string          `json:"dish name"`
	DishID               int64           `json:"dish id"`
	DishNumber           string          `json:"dish number"`
	ConfiguringPrinciple string          `json:"configuring principle"`
	Quantity             int64           `json:"quantity"`
	PaidQuantity         int64           `json:"paid quantity"`
	PricingPrinciple     string          `json:"pricing principle"`
	PriceOrDiscount      float64         `json:"price or discount"`
	PercentageDiscount   float64         `json:"percentage discount"`
	DishComponents       []DishComponent `json:"dish components"`
}

type DishComponent struct {
	ComponentName string `json:"component name"`
}

func (c Component) ToString() string {
	return fmt.Sprintf(
		"%s: Nr. - %s; Ergänzungsprinzip - %s; Anzahl - %d; Bezahlte Anzahl - %d; Preisprinzip - %s; Preis/Reduziert - %.2f; Rabatt in Prozent - %.2f",
		c.ComponentName, c.ComponentNumber, c.ConfiguringPrinciple,
		c.Quantity, c.PaidQuantity, c.PricingPrinciple,
		c.PriceOrDiscount, c.PercentageDiscount,
	)
}

func (o Offer) ToString() string {
	return fmt.Sprintf(
		"%s: Nr. - %s; Ergänzungsprinzip - %s; Anzahl - %d; Bezahlte Anzahl - %d; Preisprinzip - %s; Preis/Reduziert - %.2f; Rabatt in Prozent - %.2f",
		o.DishName, o.DishNumber, o.ConfiguringPrinciple,
		o.Quantity, o.PaidQuantity, o.PricingPrinciple,
		o.PriceOrDiscount, o.PercentageDiscount,
	)
}
