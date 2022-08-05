package models

import "time"

const (
	formatForProtocol = "2006-01-02 15:04:05"
	formatForFile     = "2006.01.02_15-04-05"
)

type ResultModel struct {
	CreationDate string         `json:"creation date"`
	Company      *FirmaBas      `json:"company,omitempty"`
	Branch       *NiederLassung `json:"branch,omitempty"`
	Menus        []*Menu        `json:"menus,omitempty"`
}

func (r *ResultModel) SetCreationTime() {
	r.CreationDate = time.Now().Format(formatForProtocol)
}

func (r *ResultModel) GetCreationTimeForFile() (string, error) {
	protocolFileCreationTime, err := time.Parse(formatForProtocol, r.CreationDate)
	if err != nil {
		return "", err
	}
	return protocolFileCreationTime.Format(formatForFile), nil
}
