package main

import "gorm.io/gorm"

type Attribute[Belongings any] interface {
	Find() Belongings // can cast to belongings
	Kvs() map[string]any
}

type Condition[Cond any, Belongings any] struct {
	ov Cond // ptr...
	Attribute[Belongings]
}

type GormRepository[Belongings any] struct {
	db *gorm.DB // db instance
}

func NewGormRepository[Belongings any](db *gorm.DB) *GormRepository[Belongings] {
	return &GormRepository[Belongings]{
		db: db,
	}
}
