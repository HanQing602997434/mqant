
// 连接建立与断连
/*
	监听器定义
		type SessionLearner interface {
			Connect(a Session) // 当连接建立 并且MQTT协议握手成功
			DisConnect(a Session) // 当连接关闭 或者客户端主动发送MQTT DisConnect命令
		}

	实现示例
		当连接建立 并且MQTT协议握手成功
		func (this *Gate) Connect(session gate.Session) {
			agent, err := this.GetGateHandler().GetAgent(session.GetSessionId())
			if err != nil {

			}
			agent.ConnTime()
		}
		
		当连接关闭 或者客户端主动发送MQTT DisConnect命令，这个函数中Session无法再继续
		后续的设置操作，只能读取部分配置内容了
		func (this *Gate) DisConnect(session gate.Session) {

		}

	设置监听器
		func (this *Gate) OnInit(app module.App, settings *conf.ModuleSettings) {
			this.Gate.OnInit(this, app, settings,
				gate.SetSessionLearner(this),
			)
		}

	获取连接代理对象(agent)
		agent, err := this.GetGateHandler().GetAgent(session.GetSessionId())
*/