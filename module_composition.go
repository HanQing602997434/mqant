
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
*/