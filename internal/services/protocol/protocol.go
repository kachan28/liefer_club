package protocol

import (
	"strings"

	"github.com/kachan28/liefer_club/app"
	"github.com/kachan28/liefer_club/internal/db"
	"github.com/kachan28/liefer_club/internal/models"
	"github.com/kachan28/liefer_club/internal/services"
	"github.com/kachan28/liefer_club/internal/services/file"
	menuService "github.com/kachan28/liefer_club/internal/services/menu"
	"github.com/kachan28/liefer_club/internal/services/tax"
)

const (
	menuDBName = "firma_pizzanova_db_menu"
	//firma tables
	companyTable = "firma_bas"
	branchTable  = "niederlassung_bas"
)

type protocolService struct{}

func MakeProtocolService() *protocolService {
	return &protocolService{}
}

var (
	dbsAndValues = map[string]map[string][]string{
		services.FirmaPizzaNovaDBName: {
			companyTable: {
				"name",
				"steuer_nr",
				"strasse",
				"haus_nr",
				"plz",
				"ort",
				"bilanzierer",
			},
			branchTable: {
				"niederlassung",
				"vat_id",
				"strasse",
				"haus_nu",
				"plz",
				"ort",
			},
		},
	}
	TaxesMap map[int64]int64
)

func (p *protocolService) MakeProtocol(conf *app.Conf) error {
	result := models.ResultModel{}
	err := tax.TaxesService{}.GetTaxesMap(conf)
	if err != nil {
		return err
	}
	for _, dbName := range conf.Dbs.Dbs {
		var conn *db.Connection
		var err error
		if dbName.Name == services.FirmaPizzaNovaDBName {
			conn, err = db.MakeConnection(conf, services.FirmaPizzaNovaDBName)
			if err != nil {
				return err
			}
			firma, err := conn.GetCompany(companyTable, dbsAndValues[services.FirmaPizzaNovaDBName][companyTable])
			if err != nil {
				return err
			}
			nieder, err := conn.GetBranch(branchTable, dbsAndValues[services.FirmaPizzaNovaDBName][branchTable])
			if err != nil {
				return err
			}
			result.Company = firma
			result.Branch = nieder
			err = conn.CloseConnection()
			if err != nil {
				return err
			}
		}
		if strings.Contains(dbName.Name, menuDBName) {
			menu, err := menuService.GetMenuService{}.GetMenu(dbName.Name, conf)
			if err != nil {
				return err
			}
			result.Menus = append(result.Menus, menu)
		}
	}

	result.SetCreationTime()
	err = file.FileService{}.WriteProtokol(result)
	if err != nil {
		return err
	}

	return nil
}
