package db

import (
	"fmt"

	"github.com/kachan28/liefer_club/internal/models"
)

const (
	OffersTable                 = "komplex_artikel_bas"
	OffersComplectPrincipeTable = "komplettieren_prinzip_dat"
	OffersPricePrincipeTable    = "komplex_artikel_preis_prinzip_dat"
	OfferToComponentsTable      = "`komplex_artikel-komplex_artikel_rel`"
	IsOffer                     = 0
)

var (
	GetOffersQuery = fmt.Sprintf(
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
		OffersTable, OffersTable, OffersTable,
		OffersComplectPrincipeTable,
		OffersTable, OffersTable,
		OffersPricePrincipeTable,
		OffersTable, OffersTable, OffersTable,
		OffersComplectPrincipeTable, OffersComplectPrincipeTable,
		OffersTable,
		OffersPricePrincipeTable, OffersPricePrincipeTable,
		OffersTable,
	)
	GetOffersToComponents = fmt.Sprintf(
		"select "+
			"%s.parent_komplex_artikel_id, "+
			"%s.komplex_artikel "+
			"from %s "+
			"join %s on %s.child_komplex_artikel_id = %s.id",
		OfferToComponentsTable,
		OffersTable,
		OfferToComponentsTable,
		OffersTable,
		OfferToComponentsTable,
		OffersTable,
	)
)

func (c *Connection) GetOffers(menu *models.Menu) error {
	offers := make([]models.Offer, 0)
	filter := fmt.Sprintf("%s.komponent=%d", OffersTable, IsOffer)
	offersRows, err := c.db.Query(c.prepareQuery(GetOffersQuery, &filter))
	if err != nil {
		return err
	}
	offerToComponentsMap, err := c.GetOfferToComponentsMap()
	if err != nil {
		return err
	}
	for offersRows.Next() {
		offer := models.Offer{}
		err = offersRows.Scan(
			&offer.DishName,
			&offer.DishID,
			&offer.DishNumber,
			&offer.ConfiguringPrinciple,
			&offer.Quantity,
			&offer.PaidQuantity,
			&offer.PricingPrinciple,
			&offer.PriceOrDiscount,
			&offer.PercentageDiscount,
		)
		if err != nil {
			return err
		}
		if _, ok := offerToComponentsMap[offer.DishID]; ok {
			offer.DishComponents = offerToComponentsMap[offer.DishID]
		}
		offers = append(offers, offer)
	}
	menu.SpecialOffersAndSetMenus[0].Offers = offers
	return nil
}

func (c *Connection) GetOfferToComponentsMap() (map[int64][]models.DishComponent, error) {
	offerToComponents := make(map[int64][]models.DishComponent, 0)
	offerToComponentsRows, err := c.db.Query(GetOffersToComponents)
	if err != nil {
		return nil, err
	}
	for offerToComponentsRows.Next() {
		var dishID int64
		var component models.DishComponent
		err = offerToComponentsRows.Scan(&dishID, &component.ComponentName)
		if _, ok := offerToComponents[dishID]; !ok {
			offerToComponents[dishID] = make([]models.DishComponent, 0)
		}
		offerToComponents[dishID] = append(offerToComponents[dishID], component)
	}
	return offerToComponents, nil
}
