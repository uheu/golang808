package message

import (
	// "fmt"
	"time"
	"strconv"
	"errors"
	"golang808/message/extra"
)

// 纬度类型
type LatitudeType int

const (
	_ LatitudeType = iota
	// 北纬
	NorthLatitudeType = 0
	// 南纬
	SouthLatitudeType = 1
)

// 经度类型
type LongitudeType int

const (
	_ LongitudeType = iota
	// 东经
	EastLongitudeType = 0
	// 西经
	WestLongitudeType = 1
)

// 位置状态
type T808_0x0200_Status uint32

type T808_0x0200 struct{
	// 警告
	Alarm uint32
	// 状态
	Status T808_0x0200_Status
	// 纬度
	Lat string
	// 经度
	Lng string
	// 海拔高度
	// 单位：米
	Altitude uint16
	// 速度
	// 单位：1/10km/h
	Speed uint16
	// 方向
	// 0-359，正北为 0，顺时针
	Direction uint16
	// 时间
	Time time.Time
	// 附加信息
	Extras []extra.Entity
}

func (t *T808_0x0200) MsgID() MsgID {
	return MsgT808_0x0200
}

func (t *T808_0x0200) Decode(data []byte) (int, error) {
	read := DecodeReader(data)

	var err error
	// 读取警告标志
	t.Alarm,err=read.ReadUint32()
	if err != nil {
		return 0, err
	}
	
	// 读取状态信息
	status,err:= read.ReadUint32()
	if err != nil {
		return 0, err
	}
	t.Status = T808_0x0200_Status(status)

	// 读取纬度信息
	latitude, err := read.ReadUint32()
	if err != nil {
		return 0, err
	}

	// 读取经度信息
	longitude, err := read.ReadUint32()
	if err != nil {
		return 0, err
	}
	
	t.Lat=strconv.FormatFloat(float64(latitude)/1000000,'f',6,64)
	t.Lng=strconv.FormatFloat(float64(longitude)/1000000,'f',6,64)

	// 读取海拔高度
	t.Altitude, err = read.ReadUint16()
	if err != nil {
		return 0, err
	}

	// 读取行驶速度
	t.Speed, err = read.ReadUint16()
	if err != nil {
		return 0, err
	}

	// 读取行驶方向
	t.Direction, err = read.ReadUint16()
	if err != nil {
		return 0, err
	}

	// 读取上报时间
	t.Time,err= read.ReadBcdTime()
	if err != nil {
		return 0, err
	}

	// 解码附加信息
	extras := make([]extra.Entity, 0)
	buffer := data[len(data)-read.read.Len():]
	for {
		if len(buffer) < 2 {
			break
		}
		id,length := buffer[0], int(buffer[1])
		buffer = buffer[2:]
		if len(buffer) < length {
			return 0,errors.New("invalid extra length")
		}
	
		extraEntity, count, err := extra.Decode(id, buffer[:length])
		if err != nil {
			return 0, errors.New("[JT/T808] unknown T808_0x0200 extra type")
		}
		if count != length {
			return 0, errors.New("invalid extra length")
		}
		extras = append(extras, extraEntity)
		buffer = buffer[length:]
	}
	if len(extras) > 0 {
		t.Extras = extras
	}
	return len(data) - read.read.Len(), nil
}
