package models

type Product struct {
	ID           uint
	Price        float64
	CategoryName string
	ImageURL     string
	Params       string
	Name         string
	Stock        bool
}
