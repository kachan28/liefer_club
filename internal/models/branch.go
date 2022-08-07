package models

type Branch struct {
	Name    string        `json:"branch name"`
	Id      string        `json:"branch id"`
	Address BranchAddress `json:"branch adress"`
	// TaxNumber string        `json:"branch tax number"`
}

type BranchAddress struct {
	Street      string `json:"street"`
	HouseNumber string `json:"house number"`
	PostalCode  string `json:"postal code"`
	Location    string `json:"location"`
}

func (b BranchAddress) PrepareAddressForExport() string {
	return b.Street + " " + b.HouseNumber + ", " + b.PostalCode + ", " + b.Location
}
