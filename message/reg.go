package message

//消息ID枚举
type MsgID uint16

const (
	// 汇报位置
	MsgT808_0x0200 MsgID = 0x0200
)

// 消息实体映射
var entityMapper = map[uint16]func() Entity{
	uint16(MsgT808_0x0200): func() Entity {
		return new(T808_0x0200)
	},
}

// 类型注册
//Entity是entity文件中个的结构体
func Register(typ uint16, creator func() Entity) {
	entityMapper[typ] = creator
}