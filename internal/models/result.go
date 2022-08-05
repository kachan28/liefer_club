package models

type ResultModel struct {
	CreationDate string         `json:"creation date"`
	Company      *FirmaBas      `json:"company,omitempty"`
	Branch       *NiederLassung `json:"branch,omitempty"`
	Menus        []*Menu        `json:"menus,omitempty"`
}
