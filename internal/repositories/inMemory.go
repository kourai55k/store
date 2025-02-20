package repositories

import (
	"sync"

	"github.com/kourai55k/store/internal/domain/models"
)

// InMemoryProductRepository is an in-memory implementation of ProductRepository.
type InMemoryProductRepository struct {
	mu       sync.RWMutex
	products map[uint]models.Product
	nextID   uint
}

// NewInMemoryProductRepository создает новый экземпляр in-memory репозитория.
func NewInMemoryProductRepository() *InMemoryProductRepository {
	return &InMemoryProductRepository{
		products: make(map[uint]models.Product),
		nextID:   1,
	}
}

// GetProducts возвращает список всех продуктов.
func (r *InMemoryProductRepository) GetProducts() ([]models.Product, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()

	var result []models.Product
	for _, p := range r.products {
		result = append(result, p)
	}
	return result, nil
}

// GetProductsByCategory возвращает продукты, фильтрованные по категории.
func (r *InMemoryProductRepository) GetProductsByCategory(category string) ([]models.Product, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()

	var result []models.Product
	for _, p := range r.products {
		if p.CategoryName == category {
			result = append(result, p)
		}
	}
	return result, nil
}

// GetProductByID возвращает продукт по его ID.
func (r *InMemoryProductRepository) GetProductByID(id uint) (models.Product, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()

	if p, ok := r.products[id]; ok {
		return p, nil
	}
	return models.Product{}, ErrProductNotFound
}

// SaveProduct сохраняет новый продукт. Присваивает продукту уникальный ID.
func (r *InMemoryProductRepository) SaveProduct(product models.Product) error {
	r.mu.Lock()
	defer r.mu.Unlock()

	product.ID = r.nextID
	r.nextID++
	r.products[product.ID] = product
	return nil
}

// UpdateProduct обновляет существующий продукт.
func (r *InMemoryProductRepository) UpdateProduct(product models.Product) error {
	r.mu.Lock()
	defer r.mu.Unlock()

	if _, exists := r.products[product.ID]; !exists {
		return ErrProductNotFound
	}
	r.products[product.ID] = product
	return nil
}

// DeleteProduct удаляет продукт по его ID.
func (r *InMemoryProductRepository) DeleteProduct(id uint) error {
	r.mu.Lock()
	defer r.mu.Unlock()

	if _, exists := r.products[id]; !exists {
		return ErrProductNotFound
	}
	delete(r.products, id)
	return nil
}
