package repository

import (
	"context"
	"errors"

	"github.com/kyraslab/go-restful-api-test/model/domain"
	"gorm.io/gorm"
)

type CategoryRepositoryImpl struct {
	db *gorm.DB
}

func NewCategoryRepository(db *gorm.DB) CategoryRepository {
	return &CategoryRepositoryImpl{db: db}
}

// Save category
func (repository *CategoryRepositoryImpl) Save(ctx context.Context, category domain.Category) (domain.Category, error) {
	if err := repository.db.WithContext(ctx).Create(&category).Error; err != nil {
		return domain.Category{}, err
	}
	return category, nil
}

// Update category
func (repository *CategoryRepositoryImpl) Update(ctx context.Context, category domain.Category) (domain.Category, error) {
	if err := repository.db.WithContext(ctx).Save(&category).Error; err != nil {
		return domain.Category{}, err
	}
	return category, nil
}

// Delete category
func (repository *CategoryRepositoryImpl) Delete(ctx context.Context, category domain.Category) error {
	if err := repository.db.WithContext(ctx).Delete(&category).Error; err != nil {
		return err
	}
	return nil
}

// FindById - Get category by ID
func (repository *CategoryRepositoryImpl) FindById(ctx context.Context, categoryId uint64) (domain.Category, error) {
	var category domain.Category
	err := repository.db.WithContext(ctx).First(&category, categoryId).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return category, errors.New("category is not found")
	}
	return category, err
}

// FindAll - Get all categories
func (repository *CategoryRepositoryImpl) FindAll(ctx context.Context) ([]domain.Category, error) {
	var categories []domain.Category
	err := repository.db.WithContext(ctx).Find(&categories).Error
	return categories, err
}
