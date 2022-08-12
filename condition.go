package ggorm

import "gorm.io/gorm"

type Condition[Belongings any] interface {
	Find() Belongings // can cast to belongings
	Kvs() (map[string]any, error)
}

type GormRepository[Belongings any] struct {
	db *gorm.DB // db instance
}
