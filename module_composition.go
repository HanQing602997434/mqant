
// 模块组成
/*
	概述
		房间模块由以下几个功能块组成

	组织结构
		1.房间
			通常一个进程就一个房间
		
		2.桌子
			桌子在房间中，一个房间可以包含多个桌子，桌子是我们房间具体实现的主体，绝大部分功能
			都在桌子内实现

			room进程——
				|- table01
				|- table02
				|- table03
				|- ...

	依赖模块
		go.mod
		
		github.com/liangdas/mqant-modules v1.3.1

	引用
		import (
			"github.com/liangdas/mqant-modules/room"
		)

	开发
		mqant-example中的示例源码
			房间demo源码

	测试
		func main() {
			task := task.LoopTask{
				C : 1, // 并发数
			}
			manager := table_test.NewManager(task) // 房间模型的demo
			// manager := test_task.NewManager(task) // gate demo
			fmt.Println("开始压测请等待")
			task.Run(manager)
			c := make(chan os.Signal, 1)
			signal.Notify(c, os.Interrupt)
			<-c
			task.Stop()
			os.Exit(1)
		}

		Connect...
		OnConnectHandler
		tabletest/HD_room_say/2 {"Trace":"0475a842d84c95ec","Error":"","Result":"success"}
		me is 81 /room/join =》 welcome to 81
		me is 81 /room/say =》 say hi from 81
*/