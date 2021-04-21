package message

import (
	// "fmt"
	"io"
	"time"
	"bytes"
	"encoding/binary"
)

type Reader struct{
	dbyte []byte
	read  *bytes.Reader
}

func DecodeReader(data []byte) Reader {
	return Reader{dbyte:data,read:bytes.NewReader(data)}
}

func (read *Reader) ReadUint16() (uint16,error){
	var buf [2]byte
	le,err:=read.read.Read(buf[:])
	if err != nil{
		return 0, err
	}

	if le != len(buf){
		return 0, err
	}
	//binary.BigEndian.Uint16() 将字节串转换成整形
	return binary.BigEndian.Uint16(buf[:]),nil
}

func (read *Reader) Read(le ...int) ([]byte,error){
	//le 是[6]
	num :=read.read.Len()
	if len(le) > 0{
		num =le[0]
	}

	curr:=len(read.dbyte)-read.read.Len()
	buf:=read.dbyte[curr:curr+num]
	read.read.Seek(int64(num),io.SeekCurrent)
	// fmt.Println("Read:",buf)
	return buf,nil
}

func (read *Reader) ReadUint32() (uint32, error) {
	if read.read.Len() < 4 {
		return 0, io.ErrUnexpectedEOF
	}

	var buf [4]byte
	_,err := read.read.Read(buf[:])
	if err != nil {
		return 0, err
	}
	return binary.BigEndian.Uint32(buf[:]), nil
}

func (read *Reader) ReadBcdTime() (time.Time, error) {
	if read.read.Len() < 6 {
		return time.Time{}, io.ErrUnexpectedEOF
	}

	var buf [6]byte
	n, err := read.read.Read(buf[:])
	if err != nil {
		return time.Time{}, err
	}
	if n != len(buf) {
		return time.Time{}, io.ErrUnexpectedEOF
	}
	return fromBCDTime(buf[:])
}