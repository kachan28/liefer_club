package db

import "github.com/kachan28/liefer_club/internal/models"

func (c *Connection) GetBranch(table string, columns []string) (*models.Branch, error) {
	q := c.prepareQueryForSelect(table, columns)
	res := c.db.QueryRow(q)
	nieder := new(models.Branch)
	err := res.Scan(&nieder.Name, &nieder.Id, &nieder.Address.Street, &nieder.Address.HouseNumber, &nieder.Address.PostalCode, &nieder.Address.Location)
	if err != nil {
		return nil, err
	}
	return nieder, nil
}
