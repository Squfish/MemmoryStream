# MemmoryStream
自定义通讯协议  主要是为了减少网络传输数据的大小

协议 GameLevel_EnterProto 包含  ProtoCode，GameLevelId，Grade 三个字段

如果使用json 则为 {"ProtoCode":12001,"GameLevelId":21,"Grade":5}
如果使用MemmoryStream 则省去了字段名 {12001,21,5}

  
  
  
 
  
