package ggorm

import "gorm.io/gorm"

type GormRepository[Belongings any] struct {
	db *gorm.DB // db instance
}

func NewGormRepository[Belongings any](db *gorm.DB) *GormRepository[Belongings] {
	return &GormRepository[Belongings]{
		db: db,
	}
}
