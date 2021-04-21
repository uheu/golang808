package message

import (
	"time"
	"errors"
)

// BCD转字符串
func bcdToString(data []byte) string {
	for {
		if len(data) == 0 {
			return ""
		}
		if data[0] != 0 {
			break
		}
		data = data[1:]
	}

	buf := make([]byte, 0, len(data)*2)
	for i := 0; i < len(data); i++ {
		buf = append(buf, data[i]&0xf0>>4+'0')
		buf = append(buf, data[i]&0x0f+'0')
	}

	return string(buf)
}

// 转为time.Time
func fromBCDTime(bcd []byte) (time.Time, error) {
	if len(bcd) != 6 {
		return time.Time{},errors.New("invalid BCD time")
	}
	t, err := time.ParseInLocation(
		"20060102150405", "20"+bcdToString(bcd), time.Local)
	if err != nil {
		return time.Time{}, err
	}
	return t, nil
}