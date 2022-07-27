package db

import (
	"database/sql"
	"fmt"
	"strings"

	_ "github.com/go-sql-driver/mysql"
	"github.com/kachan28/liefer_club/app"
	"github.com/kachan28/liefer_club/internal/models"
)

type Connection struct {
	db *sql.DB
}

type row struct {
	name string
}

func MakeConnection(conf *app.Conf, database string) (*Connection, error) {
	db, err := sql.Open("mysql", fmt.Sprintf("%s:%s@/%s", conf.Connection.User, conf.Connection.Password, database))
	if err != nil {
		return nil, err
	}
	return &Connection{
		db: db,
	}, nil
}

func (c *Connection) CloseConnection() error {
	return c.db.Close()
}

func (c *Connection) GetFirma(table string, columns []string) (*models.FirmaBas, error) {
	q := c.prepareQueryForSelect(table, columns)
	res := c.db.QueryRow(q)
	firma := new(models.FirmaBas)
	err := res.Scan(&firma.Name, &firma.SteuerNr, &firma.Strasse, &firma.HausNr, &firma.Plz, &firma.Ort, &firma.Bilanrierer)
	if err != nil {
		return nil, err
	}
	return firma, nil
}

func (c *Connection) GetNiederlassung(table string, columns []string) (*models.NiederLassung, error) {
	q := c.prepareQueryForSelect(table, columns)
	res := c.db.QueryRow(q)
	nieder := new(models.NiederLassung)
	err := res.Scan(&nieder.Niederlassung, &nieder.VatId, &nieder.Strasse, &nieder.HausNu, &nieder.Plz, &nieder.Ort)
	if err != nil {
		return nil, err
	}
	return nieder, nil
}

func (c *Connection) GetMenu(tablesAndColumns map[string][]string) (*models.Menu, error) {
	var q string
	menu := new(models.Menu)
	for table, columns := range tablesAndColumns {
		q = c.prepareQueryForSelect(table, columns)
		rows, err := c.db.Query(q)
		if err != nil {
			return nil, err
		}
		fmt.Println(rows)
		for rows.Next() {
			fmt.Println(rows.Columns())
		}
	}
	return menu, nil
}

func (c *Connection) prepareQueryForSelect(table string, columns []string) string {
	columnsString := ""
	for _, column := range columns {
		columnsString = columnsString + column + ", "
	}
	q := fmt.Sprintf("select %s from %s", columnsString, table)
	lastColonIndex := strings.LastIndex(q, ",")
	q = q[:lastColonIndex] + q[lastColonIndex+2:]
	return q
}
