package service

import (
	"encoding/json"
	"os"
	"strings"

	"github.com/kachan28/liefer_club/app"
	"github.com/kachan28/liefer_club/internal/db"
	"github.com/kachan28/liefer_club/internal/models"
)

const (
	firmaPizzaNovaDBName = "firma_pizzanova_db"
	menuDBName           = "firma_pizzanova_db_menu"
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

var dbsAndValues = map[string]map[string][]string{
	firmaPizzaNovaDBName: {
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
			"preispfandaufschlag",
		},
		ArtikelAllowedGroesseDat: {
			"groesse",
		},
		OptionsAllowedGroesseDat: {
			"groesse",
		},
	},
}

func (p *protocolService) MakeProtocol(conf *app.Conf) error {
	result := models.ResultModel{}

	for _, dbName := range conf.Dbs.Dbs {
		var conn *db.Connection
		var err error
		if dbName.Name == firmaPizzaNovaDBName {
			conn, err = db.MakeConnection(conf, firmaPizzaNovaDBName)
			if err != nil {
				return err
			}
			firma, err := conn.GetFirma(firmaBasTable, dbsAndValues[firmaPizzaNovaDBName][firmaBasTable])
			if err != nil {
				return err
			}
			nieder, err := conn.GetNiederlassung(niederLassungBasTable, dbsAndValues[firmaPizzaNovaDBName][niederLassungBasTable])
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
			conn, err = db.MakeConnection(conf, dbName.Name)
			if err != nil {
				return err
			}
			menu, err := conn.GetMenu(dbsAndValues[menuDBName])
			if err != nil {
				return err
			}
			result.Menus = append(result.Menus, menu)
			err = conn.CloseConnection()
			if err != nil {
				return err
			}
		}
	}

	jsonResult, err := json.Marshal(&result)
	if err != nil {
		return err
	}
	err = os.WriteFile("protokol.json", jsonResult, 0644)
	if err != nil {
		return err
	}
	return nil
}
