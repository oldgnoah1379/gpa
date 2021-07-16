package gpa

import (
	"context"
	"gorm.io/gorm"
)

type Repository struct {
	DB *gorm.DB
}

func New(db *gorm.DB) *Repository {
	return &Repository{db}
}

func (r Repository) Find(ctx context.Context, models interface{}, filters []Filter) error {
	return ApplyFilters(r.DB.WithContext(ctx), filters).Find(models).Error
}

func (r Repository) First(ctx context.Context, model interface{}, filters []Filter) error {
	return ApplyFilters(r.DB.WithContext(ctx), filters).First(model).Error
}

func (r Repository) Take(ctx context.Context, model interface{}, filters []Filter) error {
	return ApplyFilters(r.DB.WithContext(ctx), filters).Take(model).Error
}

func (r Repository) Create(ctx context.Context, model interface{}, opts []TransactionOption) error {
	return MakeTransaction(Create, model, opts)(r.DB.WithContext(ctx))
}

func (r Repository) Update(ctx context.Context, model interface{}, opts []TransactionOption) error {
	return MakeTransaction(Updates, model, opts)(r.DB.WithContext(ctx))
}

func (r Repository) Save(ctx context.Context, model interface{}, opts []TransactionOption) error {
	return MakeTransaction(Save, model, opts)(r.DB.WithContext(ctx))
}

func (r Repository) Delete(ctx context.Context, model interface{}, opts []TransactionOption) error {
	return MakeTransaction(Delete, model, opts)(r.DB.WithContext(ctx))
}
