package productHandler

import (
	"log/slog"

	"github.com/kourai55k/store/internal/models"
)

type ProductService interface {
	GetProducts() ([]models.Product, error)
	GetProductsByCategory(category string) ([]models.Product, error)
	GetProductByID(id uint) (models.Product, error)
	SaveProduct(product models.Product) error
	UpdateProduct(product models.Product) error
	DeleteProduct(id uint) error
}

type ProductHandler struct {
	log            *slog.Logger
	productService ProductService
}

func NewProductHandler(log *slog.Logger, productService ProductService) *ProductHandler {
	return &ProductHandler{log: log, productService: productService}
}
