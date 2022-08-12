package ggorm

import (
	"errors"
)

type Condition[Belongings any] interface {
	Find() Belongings // belongs to what model
	Kvs() map[string][]any
}

func buildKvs(m *map[string][]any) (*map[string]any, error) {
	if m == nil {
		return nil, errors.New("input is nil")
	}
	dst := make(map[string]any)
	for name, args := range *m {
		if len(args) == 1 {
			// fast path
			dst[name+" = ?"] = args[0]
			continue
		}

		if len(args)%2 == 1 {
			return nil, errors.New("invalid args")
		}
		for index := 0; index < len(args); index += 2 {
			op, v := args[index].(string)
			if !v {
				return nil, errors.New("operator must be string")
			}
			dst[name+" "+op+" ?"] = args[index+1]
		}
	}
	return &dst, nil
}
