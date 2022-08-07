package models

import (
	timeService "github.com/kachan28/liefer_club/internal/services/time"
)

type ResultModel struct {
	CreationDate string   `json:"creation date"`
	Company      *Company `json:"company,omitempty"`
	Branch       *Branch  `json:"branch,omitempty"`
	Menus        []*Menu  `json:"menus,omitempty"`
}

func (r *ResultModel) SetCreationTime() {
	r.CreationDate = timeService.TimeService{}.GetTimeForProtocol()
}

func (r *ResultModel) GetCreationTimeForFile() (string, error) {
	dt, err := timeService.TimeService{}.GetCreationTimeForFile(r.CreationDate)
	if err != nil {
		return "", err
	}
	return dt, err
}
