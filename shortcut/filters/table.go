package filters

import (
	"fmt"
	"github.com/gnoah1379/gpa"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
)

func WithSchema(schema, table schema.Tabler) gpa.Filter {
	return func(query *gorm.DB) *gorm.DB {
		return query.Table(fmt.Sprintf("%s.%s", schema, table))
	}
}



