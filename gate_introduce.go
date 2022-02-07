
// 长连接网关介绍
/*
	概述
		mqant中的Gate网关模块相对来说非常重要，它支撑了服务器与客户端的长连接通信

	特性
		1.支撑tcp、websocket通信方式
		2.默认支撑MQTT协议
		3.可自定义通信协议

	使用Gate网关模块
		gate网关模块包含的功能虽然多，但在实际开始时并不需要做过多的二次开发，开发
		者只需要继承basegate.Gate这个基础模块即可，示例如下：
			type Gate struct {
				basegate.Gate // 继承
			}

			func (this *Gate) GetType() string {
				// 很关键，需要与配置文件中的Module配置对应
				return "Gate"
			}

			func (this *Gate) Version() string {
				// 可以再监控时了解代码版本
				return "1.0.0"
			}

			func (this *Gate) OnInit(app module.App, settings *conf.ModuleSettings) {
				// 注意这里一定要用gate.Gate而不是module.BaseModule
				this.Gate.OnInit(this, app, settings)
			}
*/