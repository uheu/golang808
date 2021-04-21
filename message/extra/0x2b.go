package extra

import (
	"encoding/binary"
	"errors"
)

// 模拟量
type Extra_0x2B struct {
	serialized []byte
	value      uint16
}

func NewExtra_0x2B(val uint16) *Extra_0x2B {
	extra := Extra_0x2B{
		value: val,
	}

	var temp [2]byte
	binary.BigEndian.PutUint16(temp[:2], val)
	extra.serialized = temp[:2]
	return &extra
}

func (Extra_0x2B) ID() byte {
	return byte(TypeExtra_0x2b)
}

func (extra Extra_0x2B) Data() []byte {
	return extra.serialized
}

func (extra Extra_0x2B) Value() interface{} {
	return extra.value
}

func (extra *Extra_0x2B) Decode(data []byte) (int, error) {
	if len(data) < 2 {
		return 0,errors.New("invalid message header")
	}
	extra.value = binary.BigEndian.Uint16(data)
	return 2, nil
}
