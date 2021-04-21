package message

// 消息实体
type Entity interface {
	MsgID() MsgID
	Decode([]byte) (int, error)
}
