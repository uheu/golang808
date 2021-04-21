package message

import (
	// "fmt"
	"errors"
	"golang808/config"
)

type Message struct{
	Header Header
	Body   Entity
}

func (m *Message) Decode(dataByte []byte)error{
	data:=decodeMsg(dataByte) //relu文件的函数
	orgChksum :=data[len(data)-1] //获取包的效验码
	bin:=data[:len(data)-1]
	checkSum :=checkPacket(bin) //计算包的效验码
	if checkSum != orgChksum {
		return errors.New("效验码错误")
	}

	if len(bin) < config.MessageHeaderLen{
		return errors.New("无效的消息头")
	}
	
	var header Header  //header文件 结构体
	err:=header.Decode(bin)
	if err != nil {
		return errors.New("无效的消息头")
	}

	if !header.Proper.IsEnablePacket(){
		bin=bin[config.MessageHeaderLen:]
	}else{
		bin=bin[config.MessageHeaderLen+4:]
	}

	if uint16(len(bin)) != header.Proper.GetBodyLen(){
		return errors.New("[JT/T808] body length mismatch")
	}else{
		entity, _, err := m.decode(uint16(header.MsgID),bin)
		if err == nil {
			m.Body= entity
		}else{
			return errors.New("[JT/T808] failed to decode message")
		}
	}
	m.Header=header
	return nil
}

func (m *Message) decode(typ uint16, data []byte)(Entity, int, error) {
	creator, ok := entityMapper[typ]
	if !ok {
		return nil, 0, errors.New("entity not registered")
	}
	
	entity := creator() //返回&{0001-01-01 00:00:00 +0000 UTC}
	count, err := entity.Decode(data) //此处调用entity结构体
	if err != nil {
		return nil, 0, err
	}
	return entity, count, nil
}