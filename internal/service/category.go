package service

import "context"

type (
	CategoryService interface {
		ListCategory(ctx context.Context)
		GetCategoryByID(ctx context.Context)
		CreateCategory(ctx context.Context)
		UpdateCategory(ctx context.Context)
		DeleteCategory(ctx context.Context)
	}
)

func (s service) ListCategory(ctx context.Context)    {}
func (s service) GetCategoryByID(ctx context.Context) {}
func (s service) DeleteCategory(ctx context.Context)  {}
func (s service) CreateCategory(ctx context.Context)  {}
func (s service) UpdateCategory(ctx context.Context)  {}
