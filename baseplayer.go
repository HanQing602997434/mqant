
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

	绑定用户
		当用户请求到桌子时，在通过安全验证后可以将用户的session绑定的对应的BasePlayer中
		func (this *MyTable) doJoin(session gate.Session, msg map[string]interface{}) (err error) {
    		player := &room.BasePlayerImp{}
    		player.Bind(session)
    		。。。
		}

	更新用户状态
		用户每次给桌子发送请求，都应该更新BasePlayer状态
			1.可能session状态改变
				1）断连重连上来的网关已换
				2）有更加新的session.settings数据

			2.更新用户和桌子最后通信时间
				桌子会根据最后通信时间判断用户是否已断开跟桌子的连接（心跳）

		player.OnRequest函数会帮你做上面的两件事，但你必须在每一个用户给桌子发送的handler消息中主动调用
			func (this *MyTable) doSay(session gate.Session, msg map[string]interface{}) (err error) {
    			player := this.FindPlayer(session)
    			if player == nil {
        			return errors.New("no join")
    			}
    			player.OnRequest(session)
    			...
			}
*/