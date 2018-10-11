//===================================================
//作    者：宋宋 QQ邮箱：2845960938@qq.com
//创建时间：2018-09-11 13:09:19
//备    注：此代码为工具生成 请勿手工修改
//===================================================
using System.Collections;
using System.Collections.Generic;
using System;

/// <summary>
/// 服务器返回进入关卡消息
/// </summary>
public struct GameLevel_EnterReturnProto //: IProto
{
    public ushort ProtoCode { get { return 12002; } }

    public bool IsSuccess; //是否成功
    public short MessageId; //错误编号

    public byte[] ToArray()
    {
		using (MMoMemorySteam ms = new MMoMemorySteam())
        {
            ms.WriteUShort(ProtoCode);
            ms.WriteBool(IsSuccess);
            if(!IsSuccess)
            {
                ms.WriteShort(MessageId);
            }
            return ms.ToArray();
        }
    }

    public static GameLevel_EnterReturnProto GetProto(byte[] buffer)
    {
        GameLevel_EnterReturnProto proto = new GameLevel_EnterReturnProto();
		using (MMoMemorySteam ms = new MMoMemorySteam(buffer))
        {
            proto.IsSuccess = ms.ReadBool();
            if(!proto.IsSuccess)
            {
                proto.MessageId = ms.ReadShort();
            }
        }
        return proto;
    }
}
