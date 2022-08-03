package models

type Menu struct {
	Name           string          `json:"menu name"`
	Id             int64           `json:"menu id"`
	Db             string          `json:"-"`
	DishGroups     []DishGroup     `json:"dish groups"`
	SideDishGroups []SideDishGroup `json:"sideDishes groups"`
}
