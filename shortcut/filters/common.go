package filters

import (
	"github.com/gnoah1379/gpa"
	"gorm.io/gorm"
	"time"
)

func WithID(id interface{}) gpa.Filter {
	return func(query *gorm.DB) *gorm.DB {
		return query.Where("id=?", id)
	}
}

func WithIDs(ids ...interface{}) gpa.Filter {
	return func(query *gorm.DB) *gorm.DB {
		return query.Where("id IN ?", ids)
	}
}

func WithCreateBy(user interface{}) gpa.Filter {
	return func(query *gorm.DB) *gorm.DB {
		return query.Where("create_by = ?", user)
	}
}

func WithUpdateBy(user interface{}) gpa.Filter {
	return func(query *gorm.DB) *gorm.DB {
		return query.Where("update_by = ?", user)
	}
}

func WithCreateAtAfter(t time.Time) gpa.Filter {
	return func(query *gorm.DB) *gorm.DB {
		return query.Where("create_at > ?", t)
	}
}

func WithCreateAtBefore(t time.Time) gpa.Filter {
	return func(query *gorm.DB) *gorm.DB {
		return query.Where("create_at < ?", t)
	}
}

func WithUpdateAtAfter(t time.Time) gpa.Filter {
	return func(query *gorm.DB) *gorm.DB {
		return query.Where("create_at > ?", t)
	}
}

func WithUpdateAtBefore(t time.Time) gpa.Filter {
	return func(query *gorm.DB) *gorm.DB {
		return query.Where("create_at < ?", t)
	}
}

func WithName(name interface{}) gpa.Filter {
	return func(query *gorm.DB) *gorm.DB {
		return query.Where("name = ?", name)
	}
}

func WithCode(code interface{}) gpa.Filter {
	return func(query *gorm.DB) *gorm.DB {
		return query.Where("code = ?", code)
	}
}

func WithTitle(title interface{}) gpa.Filter {
	return func(query *gorm.DB) *gorm.DB {
		return query.Where("title = ?", title)
	}
}
