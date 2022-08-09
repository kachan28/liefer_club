package models

import (
	"encoding/json"
	"log"
)

type Company struct {
	Name           string         `json:"company name"`
	TaxNumber      string         `json:"company tax number"`
	TypeOfTaxation string         `json:"type of taxation"`
	Address        CompanyAddress `json:"company adress"`
}

type CompanyAddress struct {
	Street      string `json:"street"`
	HouseNumber string `json:"house number"`
	PostalCode  string `json:"postal code"`
	Location    string `json:"location"`
}

func (c Company) TypeOfTaxationToString() string {
	typeSting := []struct {
		TType  string `json:"type"`
		TStart string `json:"start"`
	}{}
	err := json.Unmarshal([]byte(c.TypeOfTaxation), &typeSting)
	if err != nil {
		log.Println("can't unmarshal Type Of Taxation, returning original")
		return c.TypeOfTaxation
	}
	return typeSting[0].TType + " " + typeSting[0].TStart
}

func (c CompanyAddress) PrepareAddressForExport() string {
	return c.Street + " " + c.HouseNumber + ", " + c.PostalCode + ", " + c.Location
}
