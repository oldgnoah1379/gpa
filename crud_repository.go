package gpa

import (
	"context"
	"gorm.io/gorm"
)

type CrudRepository struct {
	DB *gorm.DB
}

func New(db *gorm.DB) *CrudRepository {
	return &CrudRepository{db}
}

func (r CrudRepository) TableName(ctx context.Context) string {
	panic("implement me")
}

func (r CrudRepository) Count(ctx context.Context, filters []Filter) (int, error) {
	var result int64
	err := ApplyFilters(r.DB.WithContext(ctx).Table(r.TableName(ctx)), filters).Count(&result).Error
	return int(result), err
}

func (r CrudRepository) Find(ctx context.Context, models interface{}, filters []Filter) error {
	return ApplyFilters(r.DB.WithContext(ctx).Table(r.TableName(ctx)), filters).Find(models).Error
}

func (r CrudRepository) First(ctx context.Context, model interface{}, filters []Filter) error {
	return ApplyFilters(r.DB.WithContext(ctx).Table(r.TableName(ctx)), filters).First(model).Error
}

func (r CrudRepository) Take(ctx context.Context, model interface{}, filters []Filter) error {
	return ApplyFilters(r.DB.WithContext(ctx).Table(r.TableName(ctx)), filters).Take(model).Error
}

func (r CrudRepository) Create(ctx context.Context, model interface{}, opts []TransactionOption) error {
	return MakeTransaction(Create, model, opts)(r.DB.WithContext(ctx).Table(r.TableName(ctx)))
}

func (r CrudRepository) Update(ctx context.Context, model interface{}, opts []TransactionOption) error {
	return MakeTransaction(Updates, model, opts)(r.DB.WithContext(ctx).Table(r.TableName(ctx)))
}

func (r CrudRepository) Save(ctx context.Context, model interface{}, opts []TransactionOption) error {
	return MakeTransaction(Save, model, opts)(r.DB.WithContext(ctx).Table(r.TableName(ctx)))
}

func (r CrudRepository) Delete(ctx context.Context, model interface{}, opts []TransactionOption) error {
	return MakeTransaction(Delete, model, opts)(r.DB.WithContext(ctx).Table(r.TableName(ctx)))
}
