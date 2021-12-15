
// 网络路由协议
/*
	概述
		上一章我们实现了第一个网关，并且还实现了客户端跟后端模块之间的通信，但并没有详细阐述其中
		的路由原理，以下章节内容将详细讲解网关的路由协议和数据结构

	通信模式
		mqant是支持与客户端双向通信的长连接框架，与客户端通信有以下三种模式
		
			1.Request-Response模式
				类似http的Request-Response模式，一问一答

			2.Request-NoResponse模式
				客户端发出消息后不需要服务端回答，一问

			3.ServerPush模式
				服务器主动给客户端发送消息与app的推送功能相似

		Request-NoResponse模式通常与ServerPush模式配合使用，当后端异步响应客户端消息时非常有用

	默认路由协议
		mqant网关是这样进行约定的

		topic格式约定

			[moduleType]/[handler]/[msgid]
			moduleType	模块名称
			handler		模块中的方法
			msgid		本次消息唯一ID[可选]

		mqant网关将收到消息topic解析为以上三部分，moduleType、handler、msgid。moduleType和handler其中
		用来定位到后端模块的对应处理方法，然后进行远程RPC调用。msgid作为客户端是否需要消息应答的标记，即
		类似http的Request-Response模式。如果不设置msgid就是Request-NoResponse模式
*/