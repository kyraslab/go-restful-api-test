package service

import (
	"context"

	"github.com/kyraslab/go-restful-api-test/model/web"
)

type ProductService interface {
	Create(ctx context.Context, request web.ProductCreateRequest) (web.ProductResponse, error)
	Update(ctx context.Context, request web.ProductUpdateRequest) (web.ProductResponse, error)
	Delete(ctx context.Context, productId uint64) error
	FindById(ctx context.Context, productId uint64) (web.ProductResponse, error)
	FindAll(ctx context.Context) ([]web.ProductResponse, error)
}
