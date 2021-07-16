package gpa

import "gorm.io/gorm"

type Transaction func(tx *gorm.DB) error
type Action func(tx *gorm.DB, models interface{}) *gorm.DB
type TransactionTrigger func(tx *gorm.DB, models interface{}) error
type TransactionOption func(optional *optional)
type optional struct {
	filters       []Filter
	beforeActions []TransactionTrigger
	afterActions  []TransactionTrigger
}

func WithFilters(filters ...Filter) TransactionOption {
	return func(opt *optional) {
		opt.filters = append(opt.filters, filters...)
	}
}

func TriggerBeforeAction(triggers ...TransactionTrigger) TransactionOption {
	return func(opt *optional) {
		opt.beforeActions = append(opt.beforeActions, triggers...)
	}
}
func TriggerAfterAction(triggers ...TransactionTrigger) TransactionOption {
	return func(opt *optional) {
		opt.afterActions = append(opt.afterActions, triggers...)
	}
}

func MakeTransaction(action Action, models interface{}, opts []TransactionOption) Transaction {
	var options optional
	for _, opt := range opts {
		opt(&options)
	}
	return func(tx *gorm.DB) (err error) {
		for _, filter := range options.filters {
			tx = filter(tx)
		}
		for _, trigger := range options.beforeActions {
			if err = trigger(tx, models); err != nil {
				return err
			}
		}
		if err = action(tx, models).Error; err != nil {
			return err
		}
		for _, trigger := range options.afterActions {
			if err = trigger(tx, models); err != nil {
				return err
			}
		}
		return
	}
}
