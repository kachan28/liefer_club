package db

import (
	"fmt"

	_ "github.com/go-sql-driver/mysql"
	"github.com/kachan28/liefer_club/internal/models"
)

const (
	menuTable = "menu_bas"
)

var (
	getMenuQuery = fmt.Sprintf("select menu, id from %s where db = ?", menuTable)
)

func (c *Connection) GetMenu(menu *models.Menu) error {
	row := c.db.QueryRow(getMenuQuery, menu.Db)
	err := row.Scan(&menu.Name, &menu.Id)
	if err != nil {
		return err
	}
	return nil
}
