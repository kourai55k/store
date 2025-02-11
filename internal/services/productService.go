package services

import (
	"fmt"
	"log/slog"

	"github.com/kourai55k/store/internal/models"
)

type ProductRepository interface {
	GetProducts() ([]models.Product, error)
	GetProductsByCategory(category string) ([]models.Product, error)
	GetProductByID(id uint) (models.Product, error)
	SaveProduct(product models.Product) error
	UpdateProduct(product models.Product) error
	DeleteProduct(id uint) error
}

type ProductService struct {
	log         *slog.Logger
	productRepo ProductRepository
}

func NewProductService(log *slog.Logger, productRepo ProductRepository) *ProductService {
	return &ProductService{log: log, productRepo: productRepo}
}

func (svc *ProductService) GetProducts() ([]models.Product, error) {

	products, err := svc.productRepo.GetProducts()
	if err != nil {
		return nil, fmt.Errorf("failed to get products from storage: %w", err)
	}
	return products, nil
}

func (svc *ProductService) GetProductsByCategory(category string) ([]models.Product, error) {

	products, err := svc.productRepo.GetProductsByCategory(category)
	if err != nil {
		return nil, fmt.Errorf("failed to get products from storage: %w", err)
	}
	return products, nil
}

func (svc *ProductService) GetProductByID(id uint) (models.Product, error) {
	product, err := svc.productRepo.GetProductByID(id)
	if err != nil {
		return models.Product{}, fmt.Errorf("failed to get product from storage: %w", err)
	}
	return product, nil
}

func (svc *ProductService) SaveProduct(product models.Product) error {
	err := svc.productRepo.SaveProduct(product)
	if err != nil {
		return fmt.Errorf("failed to save product to storage: %w", err)
	}
	return nil
}

func (svc *ProductService) UpdateProduct(product models.Product) error {
	err := svc.productRepo.UpdateProduct(product)
	if err != nil {
		return fmt.Errorf("failed to update product to storage: %w", err)
	}
	return nil
}

func (svc *ProductService) DeleteProduct(id uint) error {
	err := svc.productRepo.DeleteProduct(id)
	if err != nil {
		return fmt.Errorf("failed to delete product from storage: %w", err)
	}
	return nil
}
