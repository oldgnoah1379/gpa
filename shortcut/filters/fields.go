package filters

import (
	"fmt"
	"github.com/gnoah1379/gpa"
	"gorm.io/gorm"
)

func WithFieldEqual(field string, value interface{}) gpa.Filter {
	return func(query *gorm.DB) *gorm.DB {
		return query.Where(fmt.Sprintf("%s = ?", field), value)
	}
}

func WithFieldNotEqual(field string, value interface{}) gpa.Filter {
	return func(query *gorm.DB) *gorm.DB {
		return query.Where(fmt.Sprintf("%s <> ?", field), value)
	}
}

func WithFieldGt(field string, value interface{}) gpa.Filter {
	return func(query *gorm.DB) *gorm.DB {
		return query.Where(fmt.Sprintf("%s > ?", field), value)
	}
}

func WithFieldLt(field string, value interface{}) gpa.Filter {
	return func(query *gorm.DB) *gorm.DB {
		return query.Where(fmt.Sprintf("%s < ?", field), value)
	}
}

func WithFieldGte(field string, value interface{}) gpa.Filter {
	return func(query *gorm.DB) *gorm.DB {
		return query.Where(fmt.Sprintf("%s >= ?", field), value)
	}
}

func WithFieldLte(field string, value interface{}) gpa.Filter {
	return func(query *gorm.DB) *gorm.DB {
		return query.Where(fmt.Sprintf("%s <= ?", field), value)
	}
}


func WithFieldIn(field string, value interface{}) gpa.Filter {
	return func(query *gorm.DB) *gorm.DB {
		return query.Where(fmt.Sprintf("%s IN ?", field), value)
	}
}

func WithFieldNotIn(field string, value interface{}) gpa.Filter {
	return func(query *gorm.DB) *gorm.DB {
		return query.Where(fmt.Sprintf("%s NOT IN ?", field), value)
	}
}

func WithFieldLike(field string, value interface{}) gpa.Filter {
	return func(query *gorm.DB) *gorm.DB {
		return query.Where(fmt.Sprintf("%s LIKE ?", field), value)
	}
}

func WithNotLIKE(field string, value interface{}) gpa.Filter {
	return func(query *gorm.DB) *gorm.DB {
		return query.Where(fmt.Sprintf("%s LIKE ?", field), value)
	}
}

