package db

import "github.com/kachan28/liefer_club/internal/models"

func (c *Connection) GetCompany(table string, columns []string) (*models.Company, error) {
	q := c.prepareQueryForSelect(table, columns)
	res := c.db.QueryRow(q)
	firma := new(models.Company)
	err := res.Scan(&firma.Name, &firma.TaxNumber, &firma.Address.Street, &firma.Address.HouseNumber, &firma.Address.PostalCode, &firma.Address.Location, &firma.TypeOfTaxation)
	if err != nil {
		return nil, err
	}
	return firma, nil
}
