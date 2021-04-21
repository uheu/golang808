package message

type Proper uint16

//是否分包
func (p Proper) IsEnablePacket() bool{
	val:=uint16(p)
	return val&(1 << 13 ) > 0
}

//是否加密
func (p Proper) IsEnableEncrypt() bool{
	val:=uint16(p)
	return val&(1 << 10 ) > 0
}

// 获取消息体长度
func (p *Proper) GetBodyLen() uint16 {
	// 前十位表示消息体长度
	// 0x3ff == ‭001111111111‬
	val := uint16(*p)
	return ((val << 6) >> 6) & 0x3ff
}