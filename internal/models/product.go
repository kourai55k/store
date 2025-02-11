package models

type Product struct {
	ID           uint
	Price        float64
	Measure      string // В чем измеряется товар: метр, штука, упаковка, кг.
	CategoryName string
	ImageURL     string
	Params       string
	Name         string
	Stock        bool
}
