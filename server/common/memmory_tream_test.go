package common

import (
	"testing"
	"fmt"
)

//按顺序读写即可
func TestMemmoryStream(t *testing.T) {
	//写入数据流
	stream_1 := NewEmptyMemmoryStream()
	stream_1.WriteUInt8(45)
	stream_1.WriteInt16(-45)
	stream_1.WriteString("测试1")
	stream_1.WriteFloat32(123.36)
	stream_1.WriteString("测试2")
	stream_1.WriteBool(true)
	stream_1.WriteString("测试3")
	stream_1.WriteFloat64(4555.36)

	bytes := stream_1.GetBytes()
	fmt.Println(" bytes len:",len(bytes))
	fmt.Println(" bytes:",bytes)
	//读取数据流
	stream_2 := NewMemmoryStream(bytes)
	fmt.Println(" uint8:",stream_2.ReadUInt8())
	fmt.Println(" ReadInt16:",stream_2.ReadInt16())
	fmt.Println(" ReadString:",stream_2.ReadString())
	fmt.Println(" ReadFloat32:",stream_2.ReadFloat32())
	fmt.Println(" ReadString:",stream_2.ReadString())
	fmt.Println(" ReadBool:",stream_2.ReadBool())
	fmt.Println(" ReadString:",stream_2.ReadString())
	fmt.Println(" ReadFloat64:",stream_2.ReadFloat64())

}


