
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
*/