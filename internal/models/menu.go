package models

import (
	"reflect"
	"strings"
)

type Menu struct {
	ArtikelartikelNuuStIddeleted                        string `json:"artikelartikel_nuu_st_iddeleted"`
	Artikelgroessepackungpreispfandaufschlag            string `json:"artikelgroessepackungpreispfandaufschlag"`
	ArtGruppenOpValuesartGruppenOpValuesNuuStIddeleted  string `json:"art_gruppen_op_valuesart_gruppen_op_values_nuu_st_iddeleted"`
	ArtGruppenOpValuesgroessepackungpreispfandaufschlag string `json:"art_gruppen_op_valuesgroessepackungpreispfandaufschlag"`
	Groesse                                             string `json:"groesse"`
}

func (m *Menu) ScanField(column string, value string) {
	const jsonKey = "json"
	fieldIndex := -1
	rt := reflect.TypeOf(*m)
	for i := 0; i < rt.NumField(); i++ {
		field := rt.Field(i)
		tagValue := strings.Split(field.Tag.Get(jsonKey), ",")[0]
		if tagValue == column {
			fieldIndex = i
		}
	}
	if fieldIndex != -1 {
		reflect.ValueOf(m).Elem().Field(fieldIndex).SetString(value)
	}
}
