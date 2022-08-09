package tax

import (
	"fmt"

	"github.com/kachan28/liefer_club/app"
	"github.com/kachan28/liefer_club/internal/db"
	"github.com/kachan28/liefer_club/internal/services"
)

type TaxesService struct{}

func (t TaxesService) GetTaxesMap(conf *app.Conf) error {
	//get taxes
	conn, err := db.MakeConnection(conf, services.FirmaPizzaNovaDBName)
	if err != nil {
		return err
	}
	taxes, err := conn.GetTaxes()
	if err != nil {
		return err
	}
	if len(taxes.TaxList) == 0 {
		return fmt.Errorf("empty tax list")
	}
	db.TaxesMap = taxes.ConvertToMap()
	err = conn.CloseConnection()
	if err != nil {
		return err
	}
	return nil
}
