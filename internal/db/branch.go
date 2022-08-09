package db

import (
	"fmt"

	"github.com/kachan28/liefer_club/internal/models"
)

var queryCheckIsBranchHead = "select hauptsitz from niederlassung_bas where id = %d"

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

func (c *Connection) GetBranchHeadParam(branchID int64) (int, error) {
	var isHead int
	queryCheckIsBranchHead = fmt.Sprintf(queryCheckIsBranchHead, branchID)
	err := c.db.QueryRow(queryCheckIsBranchHead).Scan(&isHead)
	if err != nil {
		return 0, err
	}
	return isHead, nil
}
