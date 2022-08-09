package models

type Taxes struct {
	TaxList []Tax
}

type Tax struct {
	Id      int64
	Procent int64
}

func (t *Taxes) ConvertToMap() map[int64]int64 {
	taxesMap := make(map[int64]int64)
	for _, tax := range t.TaxList {
		taxesMap[tax.Id] = tax.Procent
	}
	return taxesMap
}
