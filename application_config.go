
// 应用级别配置
/*
	应用级别配置
		应用级别配置可以设置应用全局所要用到的配置，例如数据库连接地址等等

		{
			"Settings":{
				"MongodbURL":"mongodb://xx:xx@xx:8015",
				"MongodbDB":"xx-server",
			}
		}

	在应用中获取

		_ = app.OnConfigurationLoaded(func(app module.App) {
			tools.MongodbUrl = app.GetSettings().Settings["MongodbURL"].(string)
			tools.MongoDB = app.GetSettings().Settings["MongodbDB"].(string)
		})

	在模块中获取
		
		func (self *admin_web) OnInit(app module.App, settings *conf.ModuleSettings) {
			self.BaseModule.Onit(self, app, settings)
			tools.MongodbUrl = app.GetSettings().Settings["MongodbURL"].(string)
			tools.MongodbDB = app.GetSettings().Settings["mongodbDB"].(string)
		}
*/