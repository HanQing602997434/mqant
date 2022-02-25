
// 初始化设置说明
/*
	概述
		room.QTable有许多设置初始项，可以用来控制以下的几大功能
			1.房间生命周期管理
				1）超时时间
				2）房间回收回调
				3）每帧更新函数Update

			2.运行周期（帧）
				多少间隔处理一次消息，一般根据业务场景而定

			3.消息队列容量管理
				1）消息接收队列容量管理
				2）消息发送队列容量管理

			4.异常处理函数
				1）handler未找到
				2）handler返回错误信息
				3）handler执行崩溃
*/