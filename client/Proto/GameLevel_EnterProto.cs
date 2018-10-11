//===================================================
//作    者：宋宋 QQ邮箱：2845960938@qq.com
//创建时间：2018-09-11 13:09:18
//备    注：此代码为工具生成 请勿手工修改
//===================================================
using System.Collections;
using System.Collections.Generic;
using System;

/// <summary>
/// 客户端发送进入游戏关卡消息
/// </summary>
public struct GameLevel_EnterProto// : IProto
{
    public ushort ProtoCode { get { return 12001; } }

    public int GameLevelId; //游戏关卡Id
    public byte Grade; //难度等级

    public byte[] ToArray()
    {
		using (MMoMemorySteam ms = new MMoMemorySteam())
        {
            ms.WriteUShort(ProtoCode);
            ms.WriteInt(GameLevelId);
            ms.WriteByte(Grade);
            return ms.ToArray();
        }
    }

    public static GameLevel_EnterProto GetProto(byte[] buffer)
    {
        GameLevel_EnterProto proto = new GameLevel_EnterProto();
		using (MMoMemorySteam ms = new MMoMemorySteam(buffer))
        {
			ms.ReadUShort();
            proto.GameLevelId = ms.ReadInt();
            proto.Grade = (byte)ms.ReadByte();
        }
        return proto;
    }
}
