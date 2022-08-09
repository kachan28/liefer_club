package db

import (
	"fmt"

	_ "github.com/go-sql-driver/mysql"
	"github.com/kachan28/liefer_club/internal/models"
)

const (
	dishNotDeleted       = 0
	dishGroupsTable      = "art_gruppen_bas"
	dishGroupToDishTable = "`artikel-art_gruppen_rel`"
	dishesTable          = "artikel_bas"
	dishToPricesTable    = "artikel_groesse_packung_preis_dat"
	pricesTable          = "artikel_allowed_groesse_dat"
)

var (
	getDishGroupsQuery        = fmt.Sprintf("select art_gruppen, id from %s", dishGroupsTable)
	getDishesByDishGroupQuery = fmt.Sprintf(
		"select"+
			"%s.artikel, "+
			"%s.label, "+
			"%s.artikel_nu, "+
			"%s.u_st_id "+
			"from %s "+
			"join %s on %s.id=%s.artikel",
		dishGroupToDishTable,
		dishesTable, dishesTable, dishesTable,
		dishGroupToDishTable,
		dishesTable, dishesTable,
		dishGroupToDishTable,
	)
	getDishToPricesWithSizesQuery = fmt.Sprintf(
		"select "+
			"%s.artikel, "+
			"%s.groesse, "+
			"%s.groesse, "+
			"%s.preis, "+
			"%s.pfandaufschlag "+
			"from %s "+
			"join %s on %s.id=%s.groesse "+
			"order by %s.artikel asc",
		dishToPricesTable,
		pricesTable,
		dishToPricesTable, dishToPricesTable, dishToPricesTable, dishToPricesTable,
		pricesTable, pricesTable,
		dishToPricesTable, dishToPricesTable,
	)
	getDishToPricesWithoutSizesQuery = fmt.Sprintf(
		"select "+
			"%s.artikel, "+
			"%s.preis "+
			"from %s ",
		dishToPricesTable,
		dishToPricesTable,
		dishToPricesTable,
	)
)

func (c *Connection) GetDishGroups(menu *models.Menu) error {
	var dishGroupsCount int
	dishGroupIndex := 0
	filter := "komplex_art_id IS NULL"
	err := c.db.QueryRow(c.prepareEntitiesCountQuery("id", dishGroupsTable, &filter)).Scan(&dishGroupsCount)
	if err != nil {
		return err
	}
	menu.DishGroups = make([]models.DishGroup, dishGroupsCount)
	rows, err := c.db.Query(c.prepareQuery(getDishGroupsQuery, &filter))
	for rows.Next() {
		rows.Scan(&menu.DishGroups[dishGroupIndex].Name, &menu.DishGroups[dishGroupIndex].ID)
		dishGroupIndex++
	}
	if err != nil {
		return err
	}
	return nil
}

func (c *Connection) GetDishes(dishGroups []models.DishGroup) error {
	var err error
	for dishGroupIndex := range dishGroups {
		var dishesCount int
		filter := fmt.Sprintf("art_gruppen=%d", dishGroups[dishGroupIndex].ID)
		countQuery := c.prepareEntitiesCountQuery("artikel", dishGroupToDishTable, &filter)
		err = c.db.QueryRow(countQuery).Scan(&dishesCount)
		if err != nil {
			return err
		}
		dishGroups[dishGroupIndex].Dishes = make([]models.Dish, dishesCount)
		dishIndex := 0
		filter += fmt.Sprintf(" and %s.deleted=%d", dishesTable, dishNotDeleted)
		getQuery := c.prepareQuery(getDishesByDishGroupQuery, &filter)
		dishesRows, err := c.db.Query(getQuery)
		if err != nil {
			return err
		}
		for dishesRows.Next() {
			dishesRows.Scan(
				&dishGroups[dishGroupIndex].Dishes[dishIndex].ID,
				&dishGroups[dishGroupIndex].Dishes[dishIndex].Name,
				&dishGroups[dishGroupIndex].Dishes[dishIndex].Number,
				&dishGroups[dishGroupIndex].Dishes[dishIndex].UStId,
			)
			dishIndex++
		}
		dishToPricesMap, err := c.GetDishPrices()
		if err != nil {
			return err
		}
		for dishIndex := range dishGroups[dishGroupIndex].Dishes {
			dishGroups[dishGroupIndex].Dishes[dishIndex].TaxValue = TaxesMap[dishGroups[dishGroupIndex].Dishes[dishIndex].UStId]
			dishGroups[dishGroupIndex].Dishes[dishIndex].DishPrices = dishToPricesMap[dishGroups[dishGroupIndex].Dishes[dishIndex].ID]
		}

	}
	return nil
}

func (c *Connection) GetDishPrices() (map[int64][]models.DishPrice, error) {
	dishPricesWithSizesRows, err := c.db.Query(getDishToPricesWithSizesQuery)
	if err != nil {
		return nil, err
	}
	dishToPriceList := make(map[int64][]models.DishPrice, 0)
	var dishID int64
	for dishPricesWithSizesRows.Next() {
		dishPrice := models.DishPrice{}
		err = dishPricesWithSizesRows.Scan(&dishID, &dishPrice.SizeOrPackage, &dishPrice.SizeOrPackageId, &dishPrice.Price, &dishPrice.BottleDepositFee)
		if err != nil {
			return nil, err
		}
		if _, ok := dishToPriceList[dishID]; !ok {
			dishToPriceList[dishID] = make([]models.DishPrice, 0)
		}
		dishToPriceList[dishID] = append(dishToPriceList[dishID], dishPrice)
	}
	filter := fmt.Sprintf("%s.groesse is null", dishToPricesTable)
	dishPricesWithoutSizesRows, err := c.db.Query(c.prepareQuery(getDishToPricesWithoutSizesQuery, &filter))
	if err != nil {
		return nil, err
	}
	for dishPricesWithoutSizesRows.Next() {
		dishPrice := models.DishPrice{}
		err = dishPricesWithoutSizesRows.Scan(&dishID, &dishPrice.Price)
		if err != nil {
			return nil, err
		}
		if _, ok := dishToPriceList[dishID]; !ok {
			dishToPriceList[dishID] = make([]models.DishPrice, 0)
		}
		dishToPriceList[dishID] = append(dishToPriceList[dishID], dishPrice)
	}
	return dishToPriceList, nil
}
