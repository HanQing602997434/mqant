
// 编写第一个模块
/*
	代码组织结构
		首先我们重新组织了一下代码目录结构，新增了一个helloworld目录来存放模块代码

			工作目录
				|-bin
					|-conf
						|-server.conf
				|-helloworld
					|-module.go
					|-xxx.go
				|-main.go

	编写第一个模块
		package helloworld
		import (
			"github.com/liangdas/mqant/conf"
			"github.com/liangdas/mqant/log"
			"github.com/liangdas/mqant/module"
			"github.com/liangdas/mqant/module/base"
		)

		var Module = func() module.Module {
			this := new(HelloWorld)
			return this
		}

		type HelloWorld struct {
			basemodule.BaseModule
		}

		func (self *HelloWorld) GetType() string {
			// 很关键，需要与配置文件中的Module配置对应
			return "helloworld"
		}

		func (self *HelloWorld) Version() string {
			// 可以在监控时了解代码版本
			return "1.0.0"
		}

		func (self *HelloWorld) OnInit(app module.App, settings *conf.ModuleSettings) {
			self.BaseModule.OnInit(self, app, settings)
			log.Info("%v模块初始化完成...", self.GetType())
		}

		func (self *HelloWorld) Run(closeSig chan bool) {
			log.Info("%v模块运行中...", self.GetType())
			log.Info("%v say hello world...", self.GetType())
			<-closeSig
			log.Info("%v模块已停止...", self.GetType())
		}

		func (self *HelloWorld) onDestory() {
			// 一定别忘了继承
			// self.BaseModule.OnDestory()
			log.Info("%v模块已回收...", self.GetType())
		}

	尝试运行
		2020-05-03T18:27:19.224684+08:00 [-] [-] [] [E] nats agent: nats: no servers available for connection
		Server configuration path : //work/go/mqant-helloworld/bin/conf/server.json
		2020-05-03T18:27:19.225643+08:00 [-] [-] [development] [I] [app.go:203] mqant  starting up
		2020-05-03T18:27:19.225725+08:00 [-] [-] [development] [I] [ModuleManager.go:50] This service ModuleGroup(ProcessID) is [development]

		运行以后没有看到我们关注的模块日志。原因是还有两项工作我们没有完成
			1.将模块加入main.go入口函数
				func main() {
					go func() {
						http.ListenAndServe("0.0.0.0:6060", nil)
					}()
					app := mqant.CreateApp(
						module.Debug(true), // 只有是在调式模式下才会在控制台打印日志，非调试模式下只在日志文件中输出日志
					)
					err := app.Run(
						helloworld.Module(),
					)
					if err != nil {
						log.Error(err.Error())
					}
				}

			2.在配置文件中加入模块配置
				{
					"Module":{
						"helloworld":[
							{
								"Id":"helloworld",
								"ProcessID":"development"
							}
						]
					},
					...
				}

				配置说明
				Module
					|-moduleType 与func GetType() string 值保持一致
						|-ProcessID 模块分组，在今后分布式部署时非常有用，默认分组为development

	运行
		2020-05-03T18:42:09.465739+08:00 [-] [-] [] [E] nats agent: nats: no servers available for connection
		Server configuration path : //work/go/mqant-helloworld/bin/conf/server.json
		2020-05-03T18:42:09.466789+08:00 [-] [-] [development] [I] [app.go:203] mqant  starting up
		2020-05-03T18:42:09.466889+08:00 [-] [-] [development] [I] [ModuleManager.go:50] This service ModuleGroup(ProcessID) is [development]
		2020-05-03T18:42:09.467058+08:00 [-] [-] [development] [I] [rpc_server.go:142] Registering node: HelloWorld@1b0073cbbab33247
		2020-05-03T18:42:09.468696+08:00 [-] [-] [development] [I] [module.go:32] HelloWorld模块初始化完成...
		2020-05-03T18:42:09.468907+08:00 [-] [-] [development] [I] [module.go:36] HelloWorld模块运行中...
		2020-05-03T18:42:09.468984+08:00 [-] [-] [development] [I] [module.go:37] HelloWorld say hello world...
		2020-05-03T18:42:09.468725+08:00 [-] [-] [development] [I] [rpc_server.go:142] Registering node: HelloWorld@1b0073cbbab33247

	停止
		追加日志

		2020-05-03T18:42:31.347133+08:00 [-] [-] [development] [I] [module.go:39] HelloWorld模块已停止...
		2020-05-03T18:42:31.347248+08:00 [-] [-] [development] [I] [module.go:45] HelloWorld模块已回收...
		2020-05-03T18:42:31.34735+08:00 [-] [-] [development] [I] [app.go:238] mqant closing down (signal: interrupt)

*/