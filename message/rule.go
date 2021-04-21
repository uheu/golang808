package message



func decodeMsg(dataByte []byte) ([]byte) {
	// 转换规则:
	// 如果 0x7d(125) 后面是 0x02(2) 转义为 0x7e(126)
	// 如果 0x7d(125) 后面是 0x01(1) 转义为 0x7d(125)
	var result []byte
	var len=len(dataByte)
	tag := false
	// 还原数据
	for i:=0;i<len;i++{
		if tag {
			tag = false
			continue
		}
		if dataByte[i] == 125 && dataByte[i+1] == 1{
			result=append(result,125)
			tag = true
		}else if dataByte[i] == 125 && dataByte[i+1] == 2{
			result=append(result,126)
			tag = true
		}else{
			result=append(result,dataByte[i])
		}
	}
	return result
}

func checkPacket(data []byte) (ret byte){
	// 取第0个，从第1个开始异或
	ret =data[0]
	len:=len(data)
	for i:=1;i<len;i++{
		ret = ret ^ data[i]
	}
	return
}