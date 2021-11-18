
// 第一个handler
/*
	我们实现了一个helloworld模块，但模块并没有真正的功能，仅仅进行声明周期的日志输出，
	实际上真正的模块应该有自身功能的核心实现
		
		作为一个网关模块

		作为一个后端模块提供核心功能的handler并且能被其他模块调用的

	代码组织结构
		首先我们重新组织了一下代码目录结构，新增了一个web目录用来存放http网关模块代码
			工程目录
				|-bin
					|-conf
						|-server.conf
				|-helloworld
					|-module.go
				|-web
					|-module.go
				|-main.go

	依赖

	服务发现
		我们需要服务发现，所以让我们启动Consul(默认)，或者通过go-plugins替换。
			不用意外mqant的服务发现模块是从go-mirco移植而来的，因此基本可以完全复用go-mirco服务发现相关插件和功能

	启动consul
		本地执行命令
		consul agent --dev

	RPC调用
		我们需要nats作为我们RPC的消息投递通道，mqant默认只内置了nats一种通道

	启动nats
		本地执行命令
		gnatsd

		bash-3.2$ gnatsd
		[21101] 2020/05/03 19:18:06.041187 [INF] Starting nats-server version 2.0.0-RC2
		[21101] 2020/05/03 19:18:06.041356 [INF] Git commit [not set]
		[21101] 2020/05/03 19:18:06.042136 [INF] Listening for client connections on 0.0.0.0:4222
		[21101] 2020/05/03 19:18:06.042148 [INF] Server id is NBJOWGQUQF44WTUDZJTUM6BPQWTPIFXXTBGYH3AZBP7U53TQP5PHCFHM
		[21101] 2020/05/03 19:18:06.042153 [INF] Server is ready

	mqant加入consul和nats
		consul和nats本身配置项较多，因此不能通过mqant的配置设置

		import (
			"github.com/liangdas/mqant"
			"github.com/liangdas/mqant/log"
			"github.com/liangdas/mqant/module"
			"github.com/liangdas/mqant/registry"
			"github.com/liangdas/mqant/registry/consul"
			"github.com/nats-io/nats.go"
			"mqant-helloworld/helloword"
		)

		func main() {
			rs := consul.NewRegistry(func(options *registry.Options) {
				options.Addrs = []string{"127.0.0.1:8500"}
			})

			nc, err := nats.Connect("nats://127.0.0.1:4222", nats.MaxReconnects(10000))
			if err != nil {
				log.Error("nats error %v", err)
				return
			}

			app := mqant.CreateApp(
				module.Debug(true), // 只有是在调式模式下才会在控制台打印日志，非调试模式下只在日志文件中输出日志
				module.Nats(nc),	// 指定nats rpc
				module.Registry(rs),// 指定服务发现
			)

			err = app.Run( // 模块都需要加到入口列表中传入框架
				helloworld.Module(),
			)

			if err != nil {
				log.Error(err.Error())
			}
		}

	编写第一个handler
		package helloworld

		import (
			"fmt"
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
			self.GetServer().RegisterGO("/say/hi", self.say) // handler
			log.Info("%v模块初始化完成...", self.GetType())
		}

		func (self *HelloWorld) OnDestory() {
			// 一定别忘了继承
			// self.BaseModule.OnDestory()
			log.Info("%v模块已回收...", self.GetType())
		}

		func (self *HelloWorld) say(name string) (r string, err error) {
			return fmt.Sprintf("hi %v", name), nil
		}

		新增handler函数
			name 为传入值
			r 为函数正常处理err情况下的返回内容
			err 为函数异常处理情况下的返回内容
			func say(name string) (r string, err error)

		将handler注册到模块中
			/say/hi 为访问地址
			self.GetServer().RegistryGO("/say/hi", self.say)

	创建一个web模块
		在helloworld模块中实现了第一个handler，但是没有地方在使用它，因此我们编写一个web模块尝试通过http
		能调用这个handler
			package web

			import (
				"context"
				"github.com/liangdas/mqant/conf"
				"github.com/liangdas/mqant/log"
				"github.com/liangdas/mqant/module"
				"github.com/liangdas/mqant/module/base"
				"github.com/liangdas/mqant/rpc"
				"io"
				"net/http"
			)

			var Module = func() module.Module {
				this := new(Web)
				return this
			}

			type Web struct {
				basemodule.BaseModule
			}

			func (self *Web) GetType() string {
				// 很关键，需要与配置文件中的Module配置对应
				return "Web"
			}

			func (self *Web) Version() string {
				// 可以在监控时了解代码版本
				return "1.0.0"
			}

			func (self *Web) OnInit(app module.App, settings *conf.ModuleSettings) {
				self.BaseModule.OnInit(self, app, settings)
			}

			func (self *Web)startHttpServer() *http.Server {
				srv := &http.Server{Addr: ":8080"}
				http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
					_ = r.ParseFrom()
					rstr, err := mqrpc.String(
						self.Call(
							context.Background(),
							"helloworld",
							"/say/hi",
							mqrpc.Param(r.Form.Get("name")),
						),
					)
					log.Info("RpcCall %v, err %v", rstr, err)
					_, _ = io.WriteString(w, rstr)
				})

				go func() {
					if err := srv.ListenAndServe(); err != nil {
						// cannot panic, because this probably is an itentional close
						log.Info("Httpserver: ListenAndServe() error: %s", err)
					}
				}()
				// returning reference so caller can call Shutdown()
				return srv
			}

			func (self *Web) Run(closeSig chan bool) {
				log.Info("web: starting HTTP server : 8080")
				srv := self.startHttpServer()
				<-closeSig
				log.Info("web: stopping HTTP server")
				// now close the server gracefully ("shutdown")
				// timeout could be given instead of nil as a https://golang.org/pkg/context/
				if err := srv.Shutdown(nil); err != nil {
					panic(err) // failure/timeout shutting down the server gracefully
				}
				log.Info("web: done. exiting")
			}

			func (self *Web) onDestory() {
				// 一定别忘了继承
				self.BaseModule.OnDestory()
			}

	特性
		1.web模块对外监听8080端口
			http://127.0.0.1:8080/say?name=mqant
		
		2.web模块收到请求后，通过RPC调用helloworld模块提供的handler，并将结果返回客户端
			http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
				_ = r.ParseForm()
				rstr, err = mqrpc.String(
					self.Call(
						conext.Background(),
						"helloworld",
						"/say/hi",
						mqrpc.Param(r.Form.Get("name")),
					),
				)
				log.Info("RpcCall %v, err %v", rstr, err)
				_, _ = io.WriteString(w, rstr)
			})

	注意事项
		1.web模块需要加入main.go函数入口中
			err = app.Run( //模块都需要加到入口列表中传入框架
				helloworld.Module(),
				web.Module(),
			)

		2.web模块需要在mqant配置文件中添加配置
			"Module":{
				"helloworld":[
					{
						"Id":"helloworld",
						"ProcessID":development
					}
				],
				"Web":[
					"Id":"Web001",
					"ProcessID":"development"
				]
			},

	运行
		Server configuration path : //work/go/mqant-helloworld/bin/conf/server.json
		2020-05-03T19:59:19.643351+08:00 [-] [-] [development] [I] [app.go:203] mqant  starting up
		2020-05-03T19:59:19.643603+08:00 [-] [-] [development] [I] [ModuleManager.go:50] This service ModuleGroup(ProcessID) is [development]
		2020-05-03T19:59:19.643762+08:00 [-] [-] [development] [I] [rpc_server.go:142] Registering node: HelloWorld@50d3be3d45da93b6
		2020-05-03T19:59:19.64951+08:00 [-] [-] [development] [I] [module.go:34] HelloWorld模块初始化完成...
		2020-05-03T19:59:19.649629+08:00 [-] [-] [development] [I] [rpc_server.go:142] Registering node: Web@6b3350636a9387de
		2020-05-03T19:59:19.64981+08:00 [-] [-] [development] [I] [module.go:38] HelloWorld模块运行中...
		2020-05-03T19:59:19.6595+08:00 [-] [-] [development] [I] [module.go:64] web: starting HTTP server :8080

	通过浏览器访问
		http://127.0.0.1:8080/say?name=mqant

	访问结果
		hi mqant
*/