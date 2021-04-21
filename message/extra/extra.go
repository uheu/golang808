package extra

import (
	"errors"
)

// 附加信息
type Entity interface {
	Decode(data []byte) (int, error)
}

func Decode(typ byte, data []byte) (Entity, int, error) {
	creator, ok := entityMapper[typ]
	if !ok {
		return nil, 0,errors.New("entity not registered")
	}
	entity := creator()
	count, err := entity.Decode(data)
	if err != nil {
		return nil, 0, err
	}
	return entity, count, nil
}