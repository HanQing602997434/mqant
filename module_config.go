
// 模块级别配置
/*
	模块配置
		模块配置分两大部分
			1.模块的启动分组ProcessID

			2.模块级别的自定义配置

	分组ID（ProcessID）
		分组ID在分布式部署种非常重要，mqatn的默认分组为development

	模块自定义配置

		"Module":{
			"moduletype":[
				{
					"Settings":{
						"StaticPath" : "static",
						"Port" : 6010
					}
				}
			],
		}

	使用
		func (self *admin_web) OnInit(app module.App, settings *conf.ModuleSettings) {
			self.BaseModule.OnInit(self, app, settings)
			self.StaticPath = self.GetModuleSettings().Settings["StaticPath"].(string)
			self.Port = int(self.GetModuleSettings().Settings["Port"].(float64))
		}
*/