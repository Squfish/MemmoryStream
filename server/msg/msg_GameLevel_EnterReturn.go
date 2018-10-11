//------------------------------------------------------
//	作者：宋宋 QQ邮箱:2845960938@qq.com
//	时间:2018-09-05 18:10:58
//	备注：此代码由工具生成
//------------------------------------------------------
package msg
import(
    "github.com/Squfish/Gameserver/common" 
)
const MsgCodeGameLevel_EnterReturn  = 12002
func init() { 
    Processor.Register(MsgCodeGameLevel_EnterReturn, &GameLevel_EnterReturn{ })
}
type GameLevel_EnterReturn struct {

    ProtoCode uint16
    IsSuccess bool //是否成功 
    MsgCode int32 //消息码 
}

func(self* GameLevel_EnterReturn)GetMsgId() uint16 {

    return self.ProtoCode
}

func(self * GameLevel_EnterReturn)UnMarshal(data[]byte)(interface{ }, error) {
    stream := common.NewMemmoryStream(data)
    self.ProtoCode = stream.ReadUInt16()
    self.IsSuccess = stream.ReadBool();

        if !self.IsSuccess { 
        self.MsgCode = stream.ReadInt32();
            }
    return self,nil
}

func(self* GameLevel_EnterReturn)Marshal(msg interface{ }) ([][]byte, error) {
    data,ok := msg.(* GameLevel_EnterReturn)
    if ok {
        stream := common.NewEmptyMemmoryStream()
        stream.WriteUInt16(data.ProtoCode)
        stream.WriteBool(data.IsSuccess)

        if !data.IsSuccess {
            stream.WriteInt32(data.MsgCode)
        }
        return [][]byte{ stream.GetBytes() },nil
    }
    return nil,nil
}
