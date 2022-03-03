
// 用户管理
/*
	概述
		给玩家发送消息都是通过session.Send实现，但在room.QTable做一下封装，能够更方便使用

		注意事项
			通过room.QTable函数发消息跟session.Send区别在于，room.QTable函数会先将要发送
			的消息都存放到消息队列中，待本周期执行完成以后再统一发送给客户端，发送之前会合并
			网关，如多个用户在同一个网关且发送的消息相同只会进行一次RPC操作

	给多个玩家发送消息
		SendCallBackMsg(player []string, topic string, body []byte) error

		其中players是session().GetSessionId()列表

	给桌子内所有玩家广播消息
		NotifyCallBackMsg(topic string, body []byte) error

	不关注结构的发送消息
		SendCallBackMsgNR(players []string, topic string, body []byte) error
		NotifyCallBackMsgNR(topic string, body []byte) error

	
*/