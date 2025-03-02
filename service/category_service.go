package service

import (
	"context"

	"github.com/kyraslab/go-restful-api-test/model/web"
)

type CategoryService interface {
	Create(ctx context.Context, request web.CategoryCreateRequest) (web.CategoryResponse, error)
	Update(ctx context.Context, request web.CategoryUpdateRequest) (web.CategoryResponse, error)
	Delete(ctx context.Context, categoryId uint64) error
	FindById(ctx context.Context, categoryId uint64) (web.CategoryResponse, error)
	FindAll(ctx context.Context) ([]web.CategoryResponse, error)
}
