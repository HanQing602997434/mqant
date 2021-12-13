
// 编写第一个网关
/*
	代码组织结构
		首先我们重新组织了以下代码目录结构，新增了一个gate目录用来存放网关代码，robot目录用来存放访问网关的mqtt客户端代码
			工程目录
				|-bin
					|-conf
						|-server.conf
				|-helloworld
					|-module.go
					|-xxx.go
				|-gate
					|-module.go
				|-robot
					|-test
						|-manager.go
						|-work.go
					|-robot_task.go
				|-main.go

	编写第一个网关
		package mgate

		import (
			"github.com/liangdas/mqant/conf"
			"github.com/liangdas/mqant/gate"
			"github.com/liangdas/mqant/gate/base"
			"github.com/liangdas/mqant/module"
		)

		var Module = func() module.Module {
			gate := new(Gate)
			return gate
		}

		type Gate struct {
			basegate.Gate // 继承
		}

		func (this *Gate) GetType() string {
			// 很关键，需要与配置文件中的Module配置对应
			return "Gate"
		}

		func (this *Gate) Version() string {
			// 可以在监控时了解代码版本
			return "1.0.0"
		}

		func (this *Gate) OnInit(app module.App, settings *conf.ModuleSettings) {
			// 注意这里一定要用gate.Gate而不是module.BaseModule
			this.Gate.OnInit(this, app, settings,
				gate.WsAddr(":3653"),
				gate.TcpAddr(":3563"),
			)
		}

	网关监听端口
*/