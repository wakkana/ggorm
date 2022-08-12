package ggorm

import (
	"gorm.io/gorm"
)

type CommonConfig struct {
	Orders map[string]string
	Offset *int
	Limit  *int
}

func (cfg *CommonConfig) Build(gm *gorm.DB) {
	if cfg == nil {
		// empty limit
		return
	}
	if len(cfg.Orders) != 0 {
		for k, v := range cfg.Orders {
			gm = gm.Order(k + v)
		}
	}

	if cfg.Offset != nil {
		gm = gm.Offset(*cfg.Offset)
	}

	if cfg.Limit != nil {
		gm = gm.Limit(*cfg.Limit)
	}
}
