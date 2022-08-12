package ggorm

import (
	"context"
)

func (r *GormRepository[EntityType]) Query(ctx context.Context,
	condition Condition[EntityType], /* cond is ptr */
	extraConfig *CommonConfig) ([]*EntityType, error) {
	var model *EntityType
	gm := r.db.WithContext(ctx).Model(model)

	limit, err := condition.Kvs()
	if err != nil {
		return nil, err
	}
	for k, v := range limit {
		gm = gm.Where(k, v) // ignore_security_alert
	}
	extraConfig.Build(gm)

	var res []*EntityType
	err = gm.Find(&res).Error
	if err != nil {
		return nil, err
	}
	return res, nil
}
