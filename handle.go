package golang808

import (
	"fmt"
	"net"
	"golang808/message"
	// "golang808/message/extra"
)

func Handle(conn net.Conn){
	// 检验标志位
	dataByte:=prefix(conn)
	var m message.Message
	if err:=m.Decode(dataByte); err != nil {
	// 	// return
	}
	entity := m.Body.(*message.T808_0x0200)
	fmt.Println("entity",entity)
	//最终获取到以下的信息
	// "IccID": message.Header.IccID,
	// "警告":    fmt.Sprintf("0x%x", entity.Alarm),
	// "状态":    fmt.Sprintf("0x%x", entity.Status),
	// "纬度":    entity.Lat,
	// "经度":    entity.Lng,
	// "海拔":    entity.Altitude,
	// "速度":    entity.Speed,
	// "方向":    entity.Direction,
	// "时间":    entity.Time,

	// for _, ext := range entity.Extras {
	// 	switch ext.ID() {
	// 	case extra.Extra_0x01{}.ID():
	// 		fields["行驶里程"] = ext.(*extra.Extra_0x01).Value()
	// 	case extra.Extra_0x02{}.ID():
	// 		fields["剩余油量"] = ext.(*extra.Extra_0x02).Value()
	// 	}
	// }
}

//检验标志位
func prefix(conn net.Conn) []byte{
	var dataByte []byte
	var buff = make([]byte,1)
	var readTag = false
	for{
		//Read方法从conn中读取最多len(buff)字节数据并写入buff
		_,err:=conn.Read(buff)
		if err != nil{
			// return nil,err
		}
		if !readTag {
			// 寻找数据头
			if buff[0] == 126 {
				readTag = true
			} else {
				continue
			}
		} else {
			// 寻找数据尾
			if buff[0] == 126 {
				break
			}
			dataByte = append(dataByte, buff[0])
		}
	}
	return dataByte
}
