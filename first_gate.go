
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
		func (this *Gate) OnInit(app module.App, settings *conf.ModuleSettings) {
			// 注意这里一定要用gate.Gate 而不是 module.BaseModule
			this.Gate.OnInit(this, app, settings,
				gate.WsAddr(":3653"),
				gate.TcpAddr(":3563"),
			)
		}

	运行
		2020-05-05T20:12:32.5603+08:00 [-] [-] [development] [I] [rpc_server.go:142] Registering node: Gate@ba5ccc4ce9feb31c
	2020-05-05T20:12:32.568484+08:00 [-] [-] [development] [I] [module.go:45] rpctest模块运行中...
	2020-05-05T20:12:32.573808+08:00 [-] [-] [development] [I] [ws_server_x.go:131] WS Listen ::3653
	2020-05-05T20:12:32.574043+08:00 [-] [-] [development] [I] [tcp_server.go:39] TCP Listen ::3563

	编写第一个客户端
		核心逻辑在robot/test/work.go

		package test_task

		import (
			"encoding/json"
			"fmt"
			MQTT "github.com/eclipse/paho.mqtt.golang"
			"github.com/liangdas/armyant/task"
			"github.com/liangdas/armyant/work"
		)

		func NewWork(manager *Manager) *Work {
			this := new(Work)
			this.manager = manager
			// opts := this.GetDefaultOptions("tcp://127.0.0.1:3563")
			opts := this.GetDefaultOptions("ws://127.0.0.1:3653")
			opts.SetConnectionLostHandler(func(client MQTT.Client, err error) {
				fmt.Println("ConnectionLost", err.Error())
			})
			opts.SetConnectionHandler(func(client MQTT.Client) {
				fmt.Println("OnConnectHandler")
			})
			err := this.Connect(opts)
			if err != nil {
				fmt.Println(err.Error())
			}

			this.On("/gate/send/test", func(client MQTT.Client, msg MQTT.Message) {
				fmt.Println(msg.Topic(), string(msg.Payload()))
			})
			return this
		}

		type Work struct {
			work.MqttWork
			manager *Manager
		}

		func (this *Work) UnmarshalResult(payload []byte) map[string]interface{} {
			rmsg := map[string]interface{}{}
			json.Unmarshal(payload, &rmsg)
			return rmsg["Result"].(map[string]interface{})
		}

		每一次请求都会调用该函数，在该函数内实现具体请求操作

		task := task.Task{
			N:1000, // 一共请求次数，会被平均分配给每一个并发协程
			C:100, // 并发数
			//QPS:10, // 每一个并发平均每秒请求次数（限流）不填代表不限流
		}

		N/C可计算每一个Work(协程) RunWorker将要调用的次数
		
		func (this *Work) RunWorker(t task.Task) {
			msg, err := this.Request("helloworld/HD_say", []byte(`{"name":"mqant"}`))
			if err != nil {
				return
			}

			fmt.Println(msg.Topic(), string(msg.Payload()))
		}
		func (this *Work) Init(t task.Task) {

		}
		func (this *Work) Init(t task.Task) {
			this.GetClient().Disconnect(0)
		}

	主动发起请求
		给后端helloworld发起一个handler(HD_say)调用
		msg, err := this.Request("helloworld/HD_say", []byte(`"name":"mqant"`))
		if err != nil {
			return
		}
		fmt.Println(msg.Topic(), string(msg.Payload()))

	监听服务器主动下发消息
		this.On("/gate/send/test", func(client MQTT.Client, msg MQTT.Message) {
			fmt.Println(msg.Topic(), string(msg.Payload()))
		})

	编写后端handler
	 
	handler实现
		func (self *HelloWorld) gatesay(session gate.Session, msg map[string]interface{}) (r string err error) {
			session.Send("/gate/send/test", []byte(fmt.Sprintf("send hi to %v", msg["name"])))
			return fmt.Sprintf("hi %v 你在网关 %v", msg["name"], session.GetServerId()), nil
		}

	注册handler
		self.GetServer().RegistryGO("HD_say", self.gatesay)

	主动发消息给客户端
		session.Send("gate/send/test", []byte(fmt.Sprintf("send hi to %v", msg["name"])))

	运行客户端（robot）
		go run robot_task.go
			开始压测请等待
			Connect...
			OnConnectHandler
			/gate/send/test send hi to mqant
			helloworld/HD_say/1 {"Trace":"5f7f87ee73f79b7a","Error":"","Result":"hi mqant 你在网关 Gate@1a8a1b29c7496c04"}
*/