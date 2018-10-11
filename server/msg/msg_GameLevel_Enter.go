//------------------------------------------------------
//	作者：宋宋 QQ邮箱:2845960938@qq.com
//	时间:2018-09-05 18:10:54
//	备注：此代码由工具生成
//------------------------------------------------------
package msg
import(
    "github.com/Squfish/Gameserver/common" 
)
const MsgCodeGameLevel_Enter  = 12001
func init() { 
    Processor.Register(MsgCodeGameLevel_Enter, &GameLevel_Enter{ })
}
type GameLevel_Enter struct {

    ProtoCode uint16
    GameLevelId int32 //游戏关卡Id 
    Grade byte //难度等级 
}

func(self *GameLevel_Enter)GetMsgId() uint16 {

    return self.ProtoCode
}

func(self *GameLevel_Enter)UnMarshal(data[]byte)(interface{ }, error) {
    stream := common.NewMemmoryStream(data)
    self.ProtoCode = stream.ReadUInt16()
    self.GameLevelId = stream.ReadInt32();
    self.Grade = stream.ReadByte();
    return self,nil
}

func(self *GameLevel_Enter)Marshal(msg interface{ }) ([][]byte, error) {
    data,ok := msg.(* GameLevel_Enter)
    if ok {
        stream := common.NewEmptyMemmoryStream()
        stream.WriteUInt16(data.ProtoCode)
        stream.WriteInt32(data.GameLevelId)
        stream.WriteByte(data.Grade)
        return [][]byte{ stream.GetBytes() },nil
    }
    return nil,nil
}
