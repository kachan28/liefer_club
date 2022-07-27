package models

type ResultModel struct {
	Company *FirmaBas      `json:"company,omitempty"`
	Branch  *NiederLassung `json:"branch,omitempty"`
	Menus   []*Menu        `json:"menus,omitempty"`
}
