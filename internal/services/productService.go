package services

import "github.com/kourai55k/store/internal/models"

type ProductRepository interface {
	GetProducts() []models.Product
	GetProductsByCategory(category string) []models.Product
	GetProductByID(id uint) models.Product
	AddProduct(product models.Product)
	UpdateProduct(product models.Product)
	DeleteProduct(id uint)
}

type ProductService struct {
	productRepo ProductRepository
}

func NewProductService(productRepo ProductRepository) *ProductService {
	return &ProductService{productRepo: productRepo}
}

func (svc *ProductService) GetProducts() []models.Product {
	return svc.productRepo.GetProducts()
}

func (svc *ProductService) GetProductsByCategory(category string) []models.Product {
	return svc.productRepo.GetProductsByCategory(category)
}

func (svc *ProductService) GetProductByID(id uint) models.Product {
	return svc.productRepo.GetProductByID(id)
}

func (svc *ProductService) AddProduct(product models.Product) {
	svc.productRepo.AddProduct(product)
}

func (svc *ProductService) UpdateProduct(product models.Product) {
	svc.productRepo.UpdateProduct(product)
}

func (svc *ProductService) DeleteProduct(id uint) {
	svc.productRepo.DeleteProduct(id)
}
