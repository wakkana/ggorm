package main

import (
	"context"
	"fmt"
)

type DBAction interface {
	TableName() string
	Condition() map[string]any
}

func (r *GormRepository[EntityType]) Query(ctx context.Context,
	condition Attribute[EntityType], /* cond is ptr */
	extraConfig *CommonConfig) ([]*EntityType, error) {
	var model *EntityType
	gm := r.db.WithContext(ctx).Model(model).Limit(1)

	res := make([]*EntityType, 0)
	err := gm.Scan(&res).Error
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (r *GormRepository[EntityType]) QueryT(ctx context.Context,
	cond Attribute[EntityType], /* cond is ptr */
	cfg *CommonConfig) ([]*EntityType, error) {
	kvs := cond.Kvs()
	for k, v := range kvs {
		fmt.Println(k, v)
	}

	cfg.BuildT(nil)

	res := make([]*EntityType, 3)
	for i := 0; i < 3; i++ {
		var a EntityType
		res = append(res, &a)
	}
	return res, nil
}
