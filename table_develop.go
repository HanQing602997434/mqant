
// 桌子开发
/*
	概述
		具体的桌子开发
	
	第一步：桌子结构体
		桌子的定义
		mqant将桌子的核心功能封装到了QTable中，开发者主要继承QTable完成业务功能的开发
		type MyTable struct {
			room.QTable
			module module.RPCModule
			players map[string]room.BasePlayer
		}
*/