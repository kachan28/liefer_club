package db

import (
	"fmt"

	"github.com/kachan28/liefer_club/internal/models"
)

const (
	ComponentsTable                 = "komplex_artikel_bas"
	ComponentsComplectPrincipeTable = "komplettieren_prinzip_dat"
	ComponentsPricePrincipeTable    = "komplex_artikel_preis_prinzip_dat"
	IsComponent                     = 1
)

var GetComponentsQuery = fmt.Sprintf(
	"select "+
		"%s.komplex_artikel, "+
		"%s.id, "+
		"%s.komplex_artikel_nu, "+
		"%s.label, "+
		"%s.item_menge, "+
		"%s.bezahlt_menge, "+
		"%s.preis_prinzip, "+
		"%s.preis_oder_nachzahlung_oder_reduziert, "+
		"%s.reduziertes_prozent "+
		"from %s "+
		"join %s on %s.id = %s.komplettieren_prinzip "+
		"join %s on %s.id = %s.preis_prinzip ",
	ComponentsTable, ComponentsTable, ComponentsTable,
	ComponentsComplectPrincipeTable,
	ComponentsTable, ComponentsTable,
	ComponentsPricePrincipeTable,
	ComponentsTable, ComponentsTable, ComponentsTable,
	ComponentsComplectPrincipeTable, ComponentsComplectPrincipeTable,
	ComponentsTable,
	ComponentsPricePrincipeTable, ComponentsPricePrincipeTable,
	ComponentsTable,
)

func (c *Connection) GetComponents(menu *models.Menu) error {
	components := make([]models.Component, 0)
	filter := fmt.Sprintf("%s.komponent=%d", ComponentsTable, IsComponent)
	componentsRows, err := c.db.Query(c.prepareQuery(GetComponentsQuery, &filter))
	if err != nil {
		return err
	}
	for componentsRows.Next() {
		component := models.Component{}
		err = componentsRows.Scan(
			&component.ComponentName,
			&component.ComponentId,
			&component.ComponentNumber,
			&component.ConfiguringPrinciple,
			&component.Quantity,
			&component.PaidQuantity,
			&component.PricingPrinciple,
			&component.PriceOrDiscount,
			&component.PercentageDiscount,
		)
		if err != nil {
			return err
		}
		components = append(components, component)
	}
	menu.SpecialOffersAndSetMenus[0].Components = components
	return nil
}
