package menu

import (
	"github.com/kachan28/liefer_club/app"
	"github.com/kachan28/liefer_club/internal/db"
	"github.com/kachan28/liefer_club/internal/models"
	"github.com/kachan28/liefer_club/internal/services"
)

type GetMenuService struct{}

func (g GetMenuService) GetMenu(menuDbName string, conf *app.Conf) (*models.Menu, error) {
	menu := new(models.Menu)
	menu.Db = menuDbName
	//get info about menu
	conn, err := db.MakeConnection(conf, services.FirmaPizzaNovaDBName)
	if err != nil {
		return nil, err
	}
	err = conn.GetMenu(menu)
	if err != nil {
		return nil, err
	}
	err = conn.CloseConnection()
	if err != nil {
		return nil, err
	}
	//get info about dish groups
	conn, err = db.MakeConnection(conf, menu.Db)
	if err != nil {
		return nil, err
	}
	err = conn.GetDishGroups(menu)
	if err != nil {
		return nil, err
	}
	//get info about dishes
	err = conn.GetDishes(menu.DishGroups)
	if err != nil {
		return nil, err
	}
	err = conn.GetSideDishGroups(menu)
	if err != nil {
		return nil, err
	}
	err = conn.GetSideDishes(menu.SideDishGroups)
	if err != nil {
		return nil, err
	}
	err = conn.CloseConnection()
	if err != nil {
		return nil, err
	}
	return menu, nil
}
