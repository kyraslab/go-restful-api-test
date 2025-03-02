package repository

import (
	"context"
	"errors"
	"testing"

	"github.com/kyraslab/go-restful-api-test/model/domain"
	"github.com/kyraslab/go-restful-api-test/repository/mocks"
	"github.com/stretchr/testify/assert"
	"go.uber.org/mock/gomock"
)

func TestProductRepository(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	repo := mocks.NewMockProductRepository(ctrl)
	ctx := context.Background()

	tests := []struct {
		name      string
		mock      func()
		method    func() (interface{}, error)
		expect    interface{}
		expectErr bool
	}{
		{
			name: "Save Success",
			mock: func() {
				product := domain.Product{ProductID: 1, Name: "Laptop Asus", Description: "This is a cutting-edge Laptop", Price: 32000000, StockQty: 128, CategoryId: 1, SKU: "ASUS-VIV-001", TaxRate: 0.15}
				repo.EXPECT().Save(ctx, product).Return(product, nil)
			},
			method: func() (interface{}, error) {
				return repo.Save(ctx, domain.Product{ProductID: 1, Name: "Laptop Asus", Description: "This is a cutting-edge Laptop", Price: 32000000, StockQty: 128, CategoryId: 1, SKU: "ASUS-VIV-001", TaxRate: 0.15})
			},
			expect:    domain.Product{ProductID: 1, Name: "Laptop Asus", Description: "This is a cutting-edge Laptop", Price: 32000000, StockQty: 128, CategoryId: 1, SKU: "ASUS-VIV-001", TaxRate: 0.15},
			expectErr: false,
		},
		{
			name: "Save Failure",
			mock: func() {
				repo.EXPECT().Save(ctx, gomock.Any()).Return(domain.Product{}, errors.New("error saving"))
			},
			method: func() (interface{}, error) {
				return repo.Save(ctx, domain.Product{Name: "Invalid"})
			},
			expect:    domain.Product{},
			expectErr: true,
		},
		{
			name: "Update Success",
			mock: func() {
				product := domain.Product{ProductID: 1, Name: "Updated Laptop Asus", Description: "This is a cutting-edge Laptop", Price: 32000000, StockQty: 128, CategoryId: 1, SKU: "ASUS-VIV-001", TaxRate: 0.15}
				repo.EXPECT().Update(ctx, product).Return(product, nil)
			},
			method: func() (interface{}, error) {
				return repo.Update(ctx, domain.Product{ProductID: 1, Name: "Updated Laptop Asus", Description: "This is a cutting-edge Laptop", Price: 32000000, StockQty: 128, CategoryId: 1, SKU: "ASUS-VIV-001", TaxRate: 0.15})
			},
			expect:    domain.Product{ProductID: 1, Name: "Updated Laptop Asus", Description: "This is a cutting-edge Laptop", Price: 32000000, StockQty: 128, CategoryId: 1, SKU: "ASUS-VIV-001", TaxRate: 0.15},
			expectErr: false,
		},
		{
			name: "FindById Success",
			mock: func() {
				repo.EXPECT().FindById(ctx, uint64(1)).Return(domain.Product{ProductID: 1, Name: "Laptop Asus", Description: "This is a cutting-edge Laptop", Price: 32000000, StockQty: 128, CategoryId: 1, SKU: "ASUS-VIV-001", TaxRate: 0.15}, nil)
			},
			method: func() (interface{}, error) {
				return repo.FindById(ctx, 1)
			},
			expect:    domain.Product{ProductID: 1, Name: "Laptop Asus", Description: "This is a cutting-edge Laptop", Price: 32000000, StockQty: 128, CategoryId: 1, SKU: "ASUS-VIV-001", TaxRate: 0.15},
			expectErr: false,
		},
		{
			name: "FindById Not Found",
			mock: func() {
				repo.EXPECT().FindById(ctx, uint64(999)).Return(domain.Product{}, errors.New("not found"))
			},
			method: func() (interface{}, error) {
				return repo.FindById(ctx, 999)
			},
			expect:    domain.Product{},
			expectErr: true,
		},
		{
			name: "FindAll Success",
			mock: func() {
				repo.EXPECT().FindAll(ctx).Return([]domain.Product{{ProductID: 1, Name: "Laptop Asus", Description: "This is a cutting-edge Laptop", Price: 32000000, StockQty: 128, CategoryId: 1, SKU: "ASUS-VIV-001", TaxRate: 0.15}}, nil)
			},
			method: func() (interface{}, error) {
				return repo.FindAll(ctx)
			},
			expect:    []domain.Product{{ProductID: 1, Name: "Laptop Asus", Description: "This is a cutting-edge Laptop", Price: 32000000, StockQty: 128, CategoryId: 1, SKU: "ASUS-VIV-001", TaxRate: 0.15}},
			expectErr: false,
		},
		{
			name: "Delete Success",
			mock: func() {
				repo.EXPECT().Delete(ctx, domain.Product{ProductID: 1}).Return(nil)
			},
			method: func() (interface{}, error) {
				return nil, repo.Delete(ctx, domain.Product{ProductID: 1})
			},
			expect:    nil,
			expectErr: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.mock()
			result, err := tt.method()

			if tt.expectErr {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
				assert.Equal(t, tt.expect, result)
			}
		})
	}
}
