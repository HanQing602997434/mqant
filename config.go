
// 配置文件结构
/*
	概述：
		配置文件可分为五大块
			1.应用级别配置

			2.模块（服务）配置

			3.日志配置

	结构
		配置文件是json格式

		{
			"Settings":{
			},
			"Module":{
				"moduletype":[
					{
						"Id":"moduletype",
						"ProcessID":"development",
						"Settings":{
						}
					}
				],
			},
			"Log":{
			}
		}
*/