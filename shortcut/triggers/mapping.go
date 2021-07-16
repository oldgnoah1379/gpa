package triggers

import (
	"github.com/gnoah1379/gpa"
	"github.com/mitchellh/mapstructure"
	"gorm.io/gorm"
)

func Mapping(target interface{}) gpa.TransactionTrigger {
	return func(tx *gorm.DB, models interface{}) error {
		return mapstructure.Decode(models, target)
	}
}

