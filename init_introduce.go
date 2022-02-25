
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

	设置函数
		QTable.OnInit(subtable SubTable, opts ...room.Option) error

	房间生命周期管理
		
		超时时间
			当桌子超过指定的时间未与客户端（指所有客户端）建立通信（上行或下行通信），则认为桌子已超时，
			会自动回收
				1.如timeout = 0则自动检查功能关闭（需要你手动控制桌子的生命周期）
				2.默认timeout = 60
					room.TimeOut(60)

		房间回收回调
			当房间被回收是我们希望能够知道，方便我们做一些资源回收工作，例如房间信息存档，将房间room移除
			等等
				room.DestroyCallbacks(func(table room.BaseTable) error {
            		log.Info("回收了房间: %v", table.TableId())
            		_ = self.room.DestroyTable(table.TableId())
            		return nil
				}),

*/