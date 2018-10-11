package msg

import (
	"github.com/name5566/leaf/network/bytes"
	"github.com/name5566/leaf/log"
	"github.com/Squfish/Gameserver/common"
)

var Processor = bytes.NewProcessor()

const TestBytesId  = 1001

func init() {

	Processor.Register(TestBytesId,&TestBytes{})

}

type TestBytes struct {

	Id uint16

	Name string

	Status bool

	money float32

}

func (t *TestBytes)GetMsgId () uint16 {
	return t.Id
}

func (t *TestBytes)UnMarshal(data []byte) (interface{}, error) {

	stream := common.NewMemmoryStream(data)
	t.Id = stream.ReadUInt16()
	t.Name = stream.ReadString()
	t.money = stream.ReadFloat32()
	log.Debug(" received masg: name %s  money %f ",t.Name,t.money)
	return t,nil
}

func (t *TestBytes)Marshal(msg interface{}) ([][]byte, error) {

	data,ok := msg.(*TestBytes)

	if ok {
		stream := common.NewEmptyMemmoryStream()
		stream.WriteUInt16(data.Id)
		stream.WriteString(data.Name)
		stream.WriteFloat32(data.money)
		return [][]byte{stream.GetBytes()},nil
	}

	return nil,nil
}