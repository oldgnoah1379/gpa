package filters

import (
	"fmt"
	"github.com/gnoah1379/gpa"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
	"gorm.io/gorm/schema"
)

func WithSchema(schema, table schema.Tabler) gpa.Filter {
	return func(query *gorm.DB) *gorm.DB {
		return query.Table(fmt.Sprintf("%s.%s", schema, table))
	}
}

func WithPreload(tableName string, cond ...interface{}) gpa.Filter {
	return func(query *gorm.DB) *gorm.DB {
		return query.Preload(tableName, cond...)
	}
}

func WithPreloadAll() gpa.Filter {
	return func(query *gorm.DB) *gorm.DB {
		return query.Preload(clause.Associations)
	}
}

func WithJoin(modelName string, args ...interface{}) gpa.Filter {
	return func(query *gorm.DB) *gorm.DB {
		return query.Joins(modelName, args...)
	}
}

func WithSelect(query string, args ...interface{}) gpa.Filter {
	return func(query *gorm.DB) *gorm.DB {
		return query.Select(query, args...)
	}
}

func WithOmit(fields ...string) gpa.Filter {
	return func(query *gorm.DB) *gorm.DB {
		return query.Omit(fields...)
	}
}
