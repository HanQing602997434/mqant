
// 基础玩家
/*
	概述
		room.BasePlayer表示在桌子中的座位它有以下几个功能和特性
			1.管理了座位跟实际用户的绑定关系
				这个座位可能有用户，也可能没有用户

			2.记录了桌子与用户的最后通信时间
				对已绑定用户的情况，如果用户断线或异常退出，需要有一种超时机制来检查

	定义
		type BasePlayer interface {
    		IsBind() bool
    		Bind(session gate.Session) BasePlayer
			
			玩家主动发请求时触发
    		OnRequest(session gate.Session)
			
			服务器主动发送消息给玩家时触发
    		OnResponse(session gate.Session)
		
			服务器跟玩家最后一次成功通信时间
    		GetLastReqResDate() int64
    		Body() interface{}
    		SetBody(body interface{})
			Session() gate.Session
    		Type() string
		}

*/