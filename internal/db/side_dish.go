package db

import (
	"fmt"

	"github.com/kachan28/liefer_club/internal/models"
)

const (
	sideDishNotDeleted           = 0
	sideDishGroupsTable          = "art_gruppen_options_bas"
	sideDishGroupToSideDishTable = "`art_gruppen-art_gruppen_options_rel`"
	sideDishesTable              = "art_gruppen_op_values_bas"
	sideDishToPricesTable        = "art_gruppen_op_values_groesse_packung_preis_dat"
	sideDishesPricesTable        = "options_allowed_groesse_dat"
)

var (
	getSideDishGroupsQuery            = fmt.Sprintf("select art_gruppen_options, id from %s", sideDishGroupsTable)
	getSideDishesBySideDishGroupQuery = fmt.Sprintf(
		"select"+
			"%s.art_gruppen_options, "+
			"%s.label, "+
			"%s.index_nu , "+
			"%s.kosten_politik "+
			"from %s "+
			"join %s on %s.id=%s.art_gruppen_options",
		sideDishGroupToSideDishTable,
		sideDishesTable, sideDishesTable, sideDishGroupToSideDishTable,
		sideDishGroupToSideDishTable,
		sideDishesTable, sideDishesTable,
		sideDishGroupToSideDishTable,
	)
	getSideDishToPricesWithSizesQuery = fmt.Sprintf(
		"select "+
			"%s.art_gruppen_op_values, "+
			"%s.groesse, "+
			"%s.groesse, "+
			"%s.preis, "+
			"%s.pfandaufschlag "+
			"from %s "+
			"join %s on %s.id=%s.groesse "+
			"order by %s.art_gruppen_op_values asc",
		sideDishToPricesTable,
		sideDishesPricesTable,
		sideDishToPricesTable, sideDishToPricesTable, sideDishToPricesTable, sideDishToPricesTable,
		sideDishesPricesTable, sideDishesPricesTable,
		sideDishToPricesTable, sideDishToPricesTable,
	)
	getSideDishToPriceWitoutSizesQuery = fmt.Sprintf(
		"select "+
			"%s.art_gruppen_op_values, "+
			"%s.preis "+
			"from %s",
		sideDishToPricesTable,
		sideDishToPricesTable,
		sideDishToPricesTable,
	)
)

func (c *Connection) GetSideDishGroups(menu *models.Menu) error {
	var sideDishGroupsCount int
	sideDishGroupIndex := 0
	err := c.db.QueryRow(c.prepareEntitiesCountQuery("id", sideDishGroupsTable, nil)).Scan(&sideDishGroupsCount)
	if err != nil {
		return err
	}
	menu.SideDishGroups = make([]models.SideDishGroup, sideDishGroupsCount)
	rows, err := c.db.Query(c.prepareQuery(getSideDishGroupsQuery, nil))
	for rows.Next() {
		rows.Scan(&menu.SideDishGroups[sideDishGroupIndex].Name, &menu.SideDishGroups[sideDishGroupIndex].ID)
		sideDishGroupIndex++
	}
	if err != nil {
		return err
	}
	return nil
}

func (c *Connection) GetSideDishes(sideDishGroups []models.SideDishGroup) error {
	var err error
	for sideDishGroupIndex := range sideDishGroups {
		var sideDishesCount int
		filter := fmt.Sprintf("art_gruppen_options=%d and %s.kosten_politik is not null", sideDishGroups[sideDishGroupIndex].ID, sideDishGroupToSideDishTable)
		countQuery := c.prepareEntitiesCountQuery("art_gruppen_options", sideDishGroupToSideDishTable, &filter)
		err = c.db.QueryRow(countQuery).Scan(&sideDishesCount)
		if err != nil {
			return err
		}
		sideDishGroups[sideDishGroupIndex].SideDishes = make([]models.SideDish, sideDishesCount)
		sideDishIndex := 0
		filter = fmt.Sprintf("%s.id=%d and %s.kosten_politik is not null and %s.deleted=%d", sideDishesTable, sideDishGroups[sideDishGroupIndex].ID, sideDishGroupToSideDishTable, sideDishesTable, sideDishNotDeleted)
		getQuery := c.prepareQuery(getSideDishesBySideDishGroupQuery, &filter)
		sideDishesRows, err := c.db.Query(getQuery)
		if err != nil {
			panic(err)
		}
		for sideDishesRows.Next() {
			err = sideDishesRows.Scan(
				&sideDishGroups[sideDishGroupIndex].SideDishes[sideDishIndex].ID,
				&sideDishGroups[sideDishGroupIndex].SideDishes[sideDishIndex].Name,
				&sideDishGroups[sideDishGroupIndex].SideDishes[sideDishIndex].Number,
				&sideDishGroups[sideDishGroupIndex].SideDishes[sideDishIndex].UStId,
			)
			if err != nil {
				return err
			}
			sideDishIndex++
		}
		sideDishToPricesMap, err := c.GetSideDishPrices()
		if err != nil {
			return err
		}
		for sideDishIndex := range sideDishGroups[sideDishGroupIndex].SideDishes {
			sideDishGroups[sideDishGroupIndex].SideDishes[sideDishIndex].TaxValue = TaxesMap[sideDishGroups[sideDishGroupIndex].SideDishes[sideDishIndex].UStId]
			sideDishGroups[sideDishGroupIndex].SideDishes[sideDishIndex].SideDishPrices = sideDishToPricesMap[sideDishGroups[sideDishGroupIndex].SideDishes[sideDishIndex].ID]
		}

	}
	return nil
}

func (c *Connection) GetSideDishPrices() (map[int64][]models.SideDishPrice, error) {
	sideDishPricesWithSizesRows, err := c.db.Query(getSideDishToPricesWithSizesQuery)
	if err != nil {
		return nil, err
	}
	sideDishToPriceList := make(map[int64][]models.SideDishPrice, 0)
	var sideDishID int64
	for sideDishPricesWithSizesRows.Next() {
		sideDishPrice := models.SideDishPrice{}
		err = sideDishPricesWithSizesRows.Scan(&sideDishID, &sideDishPrice.SizeOrPackage, &sideDishPrice.SizeOrPackageId, &sideDishPrice.Price, &sideDishPrice.BottleDepositFee)
		if err != nil {
			return nil, err
		}
		if _, ok := sideDishToPriceList[sideDishID]; !ok {
			sideDishToPriceList[sideDishID] = make([]models.SideDishPrice, 0)
		}
		sideDishToPriceList[sideDishID] = append(sideDishToPriceList[sideDishID], sideDishPrice)
	}
	filter := fmt.Sprintf("%s.groesse is null", sideDishToPricesTable)
	sideDishPricesWithoutSizesRows, err := c.db.Query(c.prepareQuery(getSideDishToPriceWitoutSizesQuery, &filter))
	if err != nil {
		return nil, err
	}
	for sideDishPricesWithoutSizesRows.Next() {
		sideDishPrice := models.SideDishPrice{}
		err = sideDishPricesWithoutSizesRows.Scan(&sideDishID, &sideDishPrice.Price)
		if err != nil {
			return nil, err
		}
		if _, ok := sideDishToPriceList[sideDishID]; !ok {
			sideDishToPriceList[sideDishID] = make([]models.SideDishPrice, 0)
		}
		sideDishToPriceList[sideDishID] = append(sideDishToPriceList[sideDishID], sideDishPrice)
	}
	return sideDishToPriceList, nil
}
