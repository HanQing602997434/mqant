
// 房间管理
/*
	概述
		房间管理比较简单，通常我们希望一个进程中可以创建多个房间，这样才能最大化利用
		服务器资源，因此我们将房间模块划分为room、table两个级别，room用来管理table

	创建房间的结构体
		func (self *tabletest) OnInit(app module.App, settings *conf.ModuleSettings) {
			self.BaseModule.OnInit(self, app, settings,
				server.RegisterInterval(15*time.Second),
				server.RegisterTTL(30*time.Second),
			)
			self.room = room.NewRoom(self)
		}
*/