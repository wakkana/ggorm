package ggorm

import (
	"context"
)

func (r *GormRepository[EntityType]) Query(ctx context.Context,
	condition Condition[EntityType], /* cond is ptr */
	commonConfig *CommonConfig) ([]*EntityType, error) {
	var model *EntityType
	gm := r.db.WithContext(ctx).Model(model)

	kvs := condition.Kvs()
	dst, err := buildKvs(&kvs)
	if err != nil {
		return nil, err
	}
	for k, v := range *dst {
		gm = gm.Where(k, v) // ignore_security_alert
	}
	commonConfig.buildKvs(gm)

	var res []*EntityType
	err = gm.Find(&res).Error
	if err != nil {
		return nil, err
	}
	return res, nil
}
