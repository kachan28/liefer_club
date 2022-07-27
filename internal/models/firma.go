package models

type FirmaBas struct {
	Name        string `json:"company name"`
	SteuerNr    string `json:"company tax number"`
	Strasse     string `json:"company adress"`
	HausNr      string `json:"company house number"`
	Plz         string `json:"plz"`
	Ort         string `json:"ort"`
	Bilanrierer string `json:"bilanrierer"`
}
