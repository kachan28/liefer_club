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
	firmaBasTable         = "firma_bas"
	niederLassungBasTable = "niederlassung_bas"
	//menu tables
	ArtikelBas                               = "artikel_bas"
	ArtikelGroessePackungPreisDat            = "artikel_groesse_packung_preis_dat"
	ArtGruppenOpValuesBas                    = "art_gruppen_op_values_bas"
	ArtGruppenOpValuesGroessePackungPreisDat = "art_gruppen_op_values_groesse_packung_preis_dat"
	ArtikelAllowedGroesseDat                 = "artikel_allowed_groesse_dat"
	OptionsAllowedGroesseDat                 = "options_allowed_groesse_dat"
)

type protocolService struct{}

func MakeProtocolService() *protocolService {
	return &protocolService{}
}

var (
	dbsAndValues = map[string]map[string][]string{
		services.FirmaPizzaNovaDBName: {
			firmaBasTable: {
				"name",
				"steuer_nr",
				"strasse",
				"haus_nr",
				"plz",
				"ort",
				"bilanzierer",
			},
			niederLassungBasTable: {
				"niederlassung",
				"vat_id",
				"strasse",
				"haus_nu",
				"plz",
				"ort",
			},
		},
		menuDBName: {
			ArtikelBas: {
				"artikel",
				"artikel_nu",
				"u_st_id",
				"deleted",
			},
			ArtikelGroessePackungPreisDat: {
				"artikel",
				"groesse",
				"packung",
				"preis",
				"pfandaufschlag",
			},
			ArtGruppenOpValuesBas: {
				"art_gruppen_op_values",
				"art_gruppen_op_values_nu",
				"u_st_id",
				"deleted",
			},
			ArtGruppenOpValuesGroessePackungPreisDat: {
				"art_gruppen_op_values",
				"groesse",
				"packung",
				"preis",
				"pfandaufschlag",
			},
			ArtikelAllowedGroesseDat: {
				"groesse",
			},
			OptionsAllowedGroesseDat: {
				"groesse",
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
			firma, err := conn.GetFirma(firmaBasTable, dbsAndValues[services.FirmaPizzaNovaDBName][firmaBasTable])
			if err != nil {
				return err
			}
			nieder, err := conn.GetNiederlassung(niederLassungBasTable, dbsAndValues[services.FirmaPizzaNovaDBName][niederLassungBasTable])
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

	err = file.FileService{}.WriteProtokol(result)
	if err != nil {
		return err
	}

	return nil
}
