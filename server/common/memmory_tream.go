//主要用于tcp网络协议
package common

import (
	"encoding/binary"
	"bytes"
	"fmt"
)

const MaxLen  = ^uint16(0)

//数值类型偏移量
const (
	Byte_OFFSET 	= 1 //  byte 1 字节
	INI_8_OFFSET 	= 1 //  8 位 1 字节
	INI_16_OFFSET 	= 2	//  16 位 2 字节
	INI_32_OFFSET 	= 4 //  32 位 4 字节
	INI_64_OFFSET 	= 8 //  64 位 8 字节
	FLOAT_32_OFFSET = 4 //  32 位 4 字节
	FLOAT_64_OFFSET = 8 //  64 位 8 字节
)

type MemmoryStream struct {
	stream  	[]byte
	index       int
	byteOrder   binary.ByteOrder
}

func NewEmptyMemmoryStream() MemmoryStream {
	return MemmoryStream{
		stream:make([]byte,0),
		index:0,
		byteOrder:binary.LittleEndian,
	}
}

//方便读取外部数据
func NewMemmoryStream(data []byte) MemmoryStream {
	return MemmoryStream{
		stream:data,
		index:0,
		byteOrder:binary.LittleEndian,
	}
}

func (ms *MemmoryStream) SetByteOrder(bOrder binary.ByteOrder) {
	ms.byteOrder = bOrder
}

//获取字节数组
func (ms *MemmoryStream) GetBytes() []byte {
	return ms.stream
}

//清空数据流
func (ms *MemmoryStream) Clear()  {
	ms.stream = make([]byte,0)
	ms.index = 0
}


func (ms *MemmoryStream) ReadByte() byte {
	arr := ms.stream[ms.index:ms.index + Byte_OFFSET]
	ms.index+=Byte_OFFSET
	var ret byte
	binary.Read(bytes.NewBuffer(arr),ms.byteOrder,&ret)
	return ret
}

func (ms *MemmoryStream) WriteByte(value byte)  {
	byteBuffer := bytes.NewBuffer([]byte{})
	binary.Write(byteBuffer,ms.byteOrder,value)
	ms.stream = append(ms.stream,byteBuffer.Bytes()...)
}

func (ms *MemmoryStream) ReadInt8() int8 {
	arr := ms.stream[ms.index:ms.index + INI_8_OFFSET]
	ms.index+=INI_8_OFFSET
	var ret int8
	binary.Read(bytes.NewBuffer(arr),ms.byteOrder,&ret)
	return ret
}

func (ms *MemmoryStream) WriteInt8(value int8)  {
	byteBuffer := bytes.NewBuffer([]byte{})
	binary.Write(byteBuffer,ms.byteOrder,value)
	ms.stream = append(ms.stream,byteBuffer.Bytes()...)
}

func (ms *MemmoryStream) ReadUInt8() uint8 {
	arr := ms.stream[ms.index:ms.index+INI_8_OFFSET]
	ms.index+=INI_8_OFFSET
	var ret uint8
	binary.Read(bytes.NewBuffer(arr),ms.byteOrder,&ret)
	return ret
}

func (ms *MemmoryStream) WriteUInt8(value uint8)  {
	byteBuffer := bytes.NewBuffer([]byte{})
	binary.Write(byteBuffer,ms.byteOrder,value)
	ms.stream = append(ms.stream,byteBuffer.Bytes()...)
}

func (ms *MemmoryStream) ReadInt16() int16 {
	arr := ms.stream[ms.index:ms.index+INI_16_OFFSET]
	ms.index+=INI_16_OFFSET
	var ret int16
	binary.Read(bytes.NewBuffer(arr),ms.byteOrder,&ret)
	return ret
}

func (ms *MemmoryStream) WriteInt16(value int16)  {
	byteBuffer := bytes.NewBuffer([]byte{})
	binary.Write(byteBuffer,ms.byteOrder,value)
	ms.stream = append(ms.stream,byteBuffer.Bytes()...)
}

func (ms *MemmoryStream) ReadUInt16() uint16 {
	arr := ms.stream[ms.index:ms.index+INI_16_OFFSET]
	ms.index+=INI_16_OFFSET
	var ret uint16
	binary.Read(bytes.NewBuffer(arr),ms.byteOrder,&ret)
	return ret
}

func (ms *MemmoryStream) WriteUInt16(value uint16)  {
	byteBuffer := bytes.NewBuffer([]byte{})
	binary.Write(byteBuffer,ms.byteOrder,value)
	ms.stream = append(ms.stream,byteBuffer.Bytes()...)
}

func (ms *MemmoryStream) ReadInt32() int32 {
	arr := ms.stream[ms.index:ms.index+INI_32_OFFSET]
	ms.index+=INI_32_OFFSET
	var ret int32
	binary.Read(bytes.NewBuffer(arr),ms.byteOrder,&ret)
	return ret
}

func (ms *MemmoryStream) WriteInt32(value int32)  {
	byteBuffer := bytes.NewBuffer([]byte{})
	binary.Write(byteBuffer,ms.byteOrder,value)
	ms.stream = append(ms.stream,byteBuffer.Bytes()...)
}

func (ms *MemmoryStream) ReadUInt32() uint32 {
	arr := ms.stream[ms.index:ms.index+INI_32_OFFSET]
	ms.index+=INI_32_OFFSET
	var ret uint32
	binary.Read(bytes.NewBuffer(arr),ms.byteOrder,&ret)
	return ret
}

func (ms *MemmoryStream) WriteUInt32(value uint32)  {
	byteBuffer := bytes.NewBuffer([]byte{})
	binary.Write(byteBuffer,ms.byteOrder,value)
	ms.stream = append(ms.stream,byteBuffer.Bytes()...)
}

func (ms *MemmoryStream) ReadInt64() int64 {
	arr := ms.stream[ms.index:ms.index+INI_64_OFFSET]
	ms.index+=INI_64_OFFSET
	var ret int64
	binary.Read(bytes.NewBuffer(arr),ms.byteOrder,&ret)
	return ret
}

func (ms *MemmoryStream) WriteInt64(value int64)  {
	byteBuffer := bytes.NewBuffer([]byte{})
	binary.Write(byteBuffer,ms.byteOrder,value)
	ms.stream = append(ms.stream,byteBuffer.Bytes()...)
}

func (ms *MemmoryStream) ReadUInt64() uint64 {
	arr := ms.stream[ms.index:ms.index+INI_64_OFFSET]
	ms.index+=INI_64_OFFSET
	var ret uint64
	binary.Read(bytes.NewBuffer(arr),ms.byteOrder,&ret)
	return ret
}

func (ms *MemmoryStream) WriteUInt64(value uint64)  {
	byteBuffer := bytes.NewBuffer([]byte{})
	binary.Write(byteBuffer,ms.byteOrder,value)
	ms.stream = append(ms.stream,byteBuffer.Bytes()...)
}

func (ms *MemmoryStream) ReadFloat32() float32  {
	arr := ms.stream[ms.index:ms.index+FLOAT_32_OFFSET]
	ms.index+=FLOAT_32_OFFSET
	var ret float32
	binary.Read(bytes.NewBuffer(arr),ms.byteOrder,&ret)
	return ret
}

func (ms *MemmoryStream) WriteFloat32(value float32)  {
	byteBuffer := bytes.NewBuffer([]byte{})
	binary.Write(byteBuffer,ms.byteOrder,value)
	ms.stream = append(ms.stream,byteBuffer.Bytes()...)
}

func (ms *MemmoryStream) ReadFloat64() float64 {
	arr := ms.stream[ms.index:ms.index+FLOAT_64_OFFSET]
	ms.index+=FLOAT_64_OFFSET
	var ret float64
	binary.Read(bytes.NewBuffer(arr),ms.byteOrder,&ret)
	return ret
}

func (ms *MemmoryStream) WriteFloat64(value float64)  {
	byteBuffer := bytes.NewBuffer([]byte{})
	binary.Write(byteBuffer,ms.byteOrder,value)
	ms.stream = append(ms.stream,byteBuffer.Bytes()...)
}

func (ms *MemmoryStream) ReadBool() bool {
	v := ms.ReadUInt8()
	if v ==1 {
		return true
	}else {
		return false
	}
}

func (ms *MemmoryStream) WriteBool(value bool)  {
	v := uint8(0)
	if value {
		v = uint8(1)
	}
	ms.WriteUInt8(v)
}

func (ms *MemmoryStream) ReadString() string {
	strlen := ms.ReadUInt16()
	arr := ms.stream[ms.index:ms.index+int(strlen)]
	ms.index += int(strlen)
	return string(arr)
}

func (ms *MemmoryStream) WriteString(value string)  {
	arr := []byte(value)
	length := len(arr)
	if uint16(length) > MaxLen {
		fmt.Errorf("MemmoryStream.WriteString %s","massage too long")
		return
	}
	ms.WriteUInt16(uint16(length))
	ms.stream = append(ms.stream,arr...)
}
