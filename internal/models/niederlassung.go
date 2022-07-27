package models

type NiederLassung struct {
	Niederlassung string `json:"niederlassung"`
	VatId         string `json:"vatId"`
	Strasse       string `json:"strasse"`
	HausNu        string `json:"hausNu"`
	Plz           string `json:"plz"`
	Ort           string `json:"ort"`
}
