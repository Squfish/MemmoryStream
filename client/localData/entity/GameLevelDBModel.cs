
//===================================================
//作    者：宋宋 QQ邮箱：2845960938@qq.com
//创建时间：2018-09-11 10:09:16
//备    注：此代码为工具生成 请勿手工修改
//===================================================
using System.Collections;
using System.Collections.Generic;
using System;

/// <summary>
/// GameLevel数据管理
/// </summary>
public partial class GameLevelDBModel : AbstractDBModel<GameLevelDBModel, GameLevelEntity>
{
    /// <summary>
    /// 文件名称
    /// </summary>
    protected override string FileName { get { return "GameLevel.data"; } }

    /// <summary>
    /// 创建实体
    /// </summary>
    /// <param name="parse"></param>
    /// <returns></returns>
    protected override GameLevelEntity CreateEntity(ReadExcelAndParss parse)
    {
        GameLevelEntity entity = new GameLevelEntity();
        entity.Id = parse.GetFieldValue("Id").ToInt();
        entity.ChapterID = parse.GetFieldValue("ChapterID").ToInt();
        entity.Name = parse.GetFieldValue("Name");
        entity.SceneName = parse.GetFieldValue("SceneName");
        entity.SmallMapImg = parse.GetFieldValue("SmallMapImg");
        entity.isBoss = parse.GetFieldValue("isBoss").ToInt();
        entity.Ico = parse.GetFieldValue("Ico");
        entity.PosInMap = parse.GetFieldValue("PosInMap");
        entity.DlgPic = parse.GetFieldValue("DlgPic");
        entity.CameraRotation = parse.GetFieldValue("CameraRotation");
        entity.Audio_BG = parse.GetFieldValue("Audio_BG");
        return entity;
    }
}
