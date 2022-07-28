package models

type ResultModel struct {
	CreationData string         `json:"creation date"`
	Company      *FirmaBas      `json:"company,omitempty"`
	Branch       *NiederLassung `json:"branch,omitempty"`
	Menus        []*Menu        `json:"menus,omitempty"`
}
