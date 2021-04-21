package message

import (
	// "fmt"
	"strconv"
)

// 封包信息
type Packet struct {
	Sum uint16
	Seq uint16
}

//消息头
type Header struct{
	MsgID  		MsgID //第二个MsgID则是reg文件中的配置
	Proper    	Proper
	IccID       uint64
	MsgSerialNo uint16
	Packet      *Packet
}

//解码
func (h *Header) Decode(data []byte) error{
	// DecodeReader(data) //buff文件中的函数
	read:=DecodeReader(data)

	//获取消息ID
	msgID,err:= read.ReadUint16()
	if err != nil {
		return err
	}

	// 读取消息体属性
	proper,err:=read.ReadUint16()
	if err != nil {
		return err
	}

	// 读取终端号码
	temp,err:=read.Read(6)
	if err != nil {
		return err
	}
	iccID,err:=strconv.ParseUint(bcdToString(temp),10,64)
	if err != nil {
		return err
	}

	// 读取消息流水号
	serialNo,err:=read.ReadUint16()
	if err != nil {
		return err
	}

	// 读取分包信息
	if Proper(proper).IsEnablePacket() {
		//结构体
		var p Packet

		// 读取分包总数
		p.Sum, err = read.ReadUint16()
		if err != nil {
			return err
		}
		
		// 读取分包序列号
		p.Seq, err = read.ReadUint16()
		if err != nil {
			return err
		}
		h.Packet = &p
	}

	h.MsgID 		= MsgID(msgID)
	h.Proper 		= Proper(proper)
	h.IccID			= iccID
	h.MsgSerialNo	= serialNo
	return nil 
}

