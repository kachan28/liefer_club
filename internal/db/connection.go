package db

import (
	"database/sql"
	"fmt"
	"strings"

	_ "github.com/go-sql-driver/mysql"
	"github.com/kachan28/liefer_club/app"
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

func (c *Connection) prepareEntitiesCountQuery(column, table string, filter *string) string {
	query := fmt.Sprintf("select count(%s) from %s", column, table)
	if filter != nil {
		query = c.prepareQuery(query, filter)
	}
	return query
}

func (c *Connection) prepareQuery(query string, filter *string) string {
	if filter != nil {
		query = fmt.Sprintf("%s where %s", query, *filter)
	}
	return query
}
