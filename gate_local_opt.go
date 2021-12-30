
// 网关本地操作
/*
	概述
		如果业务需要再网关模块操作本网关的功能，可以使用GateHandler，这样可以避免RPC操作

	获取GateHandler
		this.Gate.GetGateHandler()

	获取网关当前连接数
		this.Gate.GetGateHandler().GetAgentNum()

	获取连接agent
		this.Gate.GetGateHandler().GetAgent(Sessionid string) (Agent, error)

	给指定连接发送消息
		this.Gate.GetGateHandler().Send(span log.TraceSpan, Sessionid string, topic string, body []byte) (result interface{}, err string)

	群发消息
		this.Gate.GetGateHandler().BroadCast(span log.TraceSpan, topic string, body []byte) (int64, string)
*/