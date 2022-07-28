package db

import (
	"fmt"

	"github.com/kachan28/liefer_club/internal/models"
)

const (
	TaxesTable = "umsatzsteuer_bassis_gruppen"
)

var (
	TaxesMap map[int64]int64

	getTaxesQuery = fmt.Sprintf(
		"select id, procent from %s",
		TaxesTable,
	)
)

func (c *Connection) GetTaxes() (*models.Taxes, error) {
	var taxesCount int64
	taxes := new(models.Taxes)
	err := c.db.QueryRow(c.prepareEntitiesCountQuery("id", TaxesTable, nil)).Scan(&taxesCount)
	if err != nil {
		return nil, err
	}
	taxes.TaxList = make([]models.Tax, taxesCount)
	taxesRows, err := c.db.Query(getTaxesQuery)
	if err != nil {
		return nil, err
	}
	taxIndex := 0
	for taxesRows.Next() {
		taxesRows.Scan(&taxes.TaxList[taxIndex].Id, &taxes.TaxList[taxIndex].Procent)
		taxIndex++
	}
	return taxes, nil
}
