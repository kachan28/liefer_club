package models

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

func (c CompanyAddress) PrepareAddressForExport() string {
	return c.Street + " " + c.HouseNumber + ", " + c.PostalCode + ", " + c.Location
}
