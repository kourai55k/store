package services

import (
	"fmt"

	"github.com/kourai55k/store/internal/domain/models"
)

type CategoryRepository interface {
	GetCategories() ([]models.Category, error)
	GetCategoryByID(id uint) (models.Category, error)
	SaveCategory(category models.Category) error
	UpdateCategory(category models.Category) error
	DeleteCategory(id uint) error
}

type CategoryService struct {
	categoryRepo CategoryRepository
}

func NewCategoryService(categoryRepo CategoryRepository) *CategoryService {
	return &CategoryService{
		categoryRepo: categoryRepo,
	}
}

func (svc *CategoryService) GetCategories() ([]models.Category, error) {
	categories, err := svc.categoryRepo.GetCategories()
	if err != nil {
		return nil, fmt.Errorf("failed to get categories from storage: %w", err)
	}
	return categories, nil
}

func (svc *CategoryService) GetCategoryByID(id uint) (models.Category, error) {
	category, err := svc.categoryRepo.GetCategoryByID(id)
	if err != nil {
		return models.Category{}, fmt.Errorf("failed to get category from storage: %w", err)
	}
	return category, nil
}

func (svc *CategoryService) SaveCategory(category models.Category) error {
	err := svc.categoryRepo.SaveCategory(category)
	if err != nil {
		return fmt.Errorf("failed to save category to storage: %w", err)
	}
	return nil
}

func (svc *CategoryService) UpdateCategory(category models.Category) error {
	err := svc.categoryRepo.UpdateCategory(category)
	if err != nil {
		return fmt.Errorf("failed to update category to storage: %w", err)
	}
	return nil
}

func (svc *CategoryService) DeleteCategory(id uint) error {
	err := svc.categoryRepo.DeleteCategory(id)
	if err != nil {
		return fmt.Errorf("failed to delete category from storage: %w", err)
	}
	return nil
}
