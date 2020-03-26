package usecase

import (
	"context"
	"lab/productLab/internal/entity"
)

// ProductStore interface to store product
type ProductStore interface {
	Create(ctx context.Context, product *entity.Product) (*entity.Product, error)
	ByID(ctx context.Context, id int) (*entity.Product, error)
	Update(ctx context.Context, id int, product *entity.Product) error
	Delete(ctx context.Context, id int) error
	List(ctx context.Context, start, count int) ([]*entity.Product, error)
}

// ProductUC struct for store
type ProductUC struct {
	Store ProductStore
}

// NewProductUC new usecase products
func NewProductUC(ps ProductStore) ProductUC {
	return ProductUC{
		Store: ps,
	}
}

// Create usecase create product
func (p ProductUC) Create(ctx context.Context, product *entity.Product) (*entity.Product, error) {
	pr, err := p.Store.Create(ctx, product)
	if err != nil {
		return nil, err
	}
	return pr, nil
}

// ByID usecase gets by id
func (p ProductUC) ByID(ctx context.Context, id int) (*entity.Product, error) {
	pr, err := p.Store.ByID(ctx, id)
	if err != nil {
		return nil, err
	}
	return pr, err
}

// Update usecase to update a product
func (p ProductUC) Update(ctx context.Context, id int, product *entity.Product) error {
	err := p.Store.Update(ctx, id, product)
	if err != nil {
		return err
	}
	return nil
}

// Delete usecase delete a product
func (p ProductUC) Delete(ctx context.Context, id int) error {
	err := p.Store.Delete(ctx, id)
	if err != nil {
		return err
	}
	return nil
}

//List usecase list all products
func (p ProductUC) List(ctx context.Context, start, count int) ([]*entity.Product, error) {
	products, err := p.Store.List(ctx, start, count)
	if err != nil {
		return products, err
	}
	return products, err
}
