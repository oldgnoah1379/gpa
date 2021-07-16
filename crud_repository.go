package gpa

import (
	"context"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
)

type CrudRepository struct {
	*gorm.DB
	schema.Tabler
}

func New(db *gorm.DB, entity schema.Tabler) *CrudRepository {
	return &CrudRepository{db, entity}
}

func (r CrudRepository) Count(ctx context.Context, filters []Filter) (int, error) {
	var result int64
	err := ApplyFilters(r.DB.WithContext(ctx).Table(r.TableName()), filters).Count(&result).Error
	return int(result), err
}

func (r CrudRepository) Find(ctx context.Context, models interface{}, filters []Filter) error {
	return ApplyFilters(r.DB.WithContext(ctx).Table(r.TableName()), filters).Find(models).Error
}

func (r CrudRepository) First(ctx context.Context, model interface{}, filters []Filter) error {
	return ApplyFilters(r.DB.WithContext(ctx).Table(r.TableName()), filters).First(model).Error
}

func (r CrudRepository) Take(ctx context.Context, model interface{}, filters []Filter) error {
	return ApplyFilters(r.DB.WithContext(ctx).Table(r.TableName()), filters).Take(model).Error
}

func (r CrudRepository) Create(ctx context.Context, model interface{}, opts []TransactionOption) error {
	return MakeTransaction(Create, model, opts)(r.DB.WithContext(ctx).Table(r.TableName()))
}

func (r CrudRepository) Update(ctx context.Context, model interface{}, opts []TransactionOption) error {
	return MakeTransaction(Updates, model, opts)(r.DB.WithContext(ctx).Table(r.TableName()))
}

func (r CrudRepository) Save(ctx context.Context, model interface{}, opts []TransactionOption) error {
	return MakeTransaction(Save, model, opts)(r.DB.WithContext(ctx).Table(r.TableName()))
}

func (r CrudRepository) Delete(ctx context.Context, model interface{}, opts []TransactionOption) error {
	return MakeTransaction(Delete, model, opts)(r.DB.WithContext(ctx).Table(r.TableName()))
}
