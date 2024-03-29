
// mqtt协议
/*
	概述
		MQTT协议是为大量计算能力有限，且工作在低带宽、不可靠的网络的远程传感器和控制设备通讯而设计的
		协议，它具有以下主要的几项特性：
			
			1.使用发布/订阅消息模式，提供一对多的消息发布，解除应用程序耦合；
			2.对负载内容屏蔽的消息传输；
			3.使用TCP/IP提供网络连接；
			4.有三种消息发布服务质量；
				1."至多一次"，消息发布完全依赖底层TCP/IP网络。会发生消息丢失或重复。这一级别可用于如下情
				况，环境传感器数据，丢失一次读记录无所谓，因为不久后还会有第二次发送。
				
				2."至少一次"，确保消息到达，但消息重复可能会发生。

				3."只有一次"，确保消息到达一次。这一级别可用于如下情况，在计费系统中，消息重复或丢失会导
				致不正确的结果。
			5.小型传输，开销很小（固定长度的头部是2字节），协议交换最小化，以降低网络流量；
			6.使用Last Will和Testament特性通知有关各方客户端异常中断的机制；

		总的来说MQTT协议是非常精简的通信协议，同时也有完善的【心跳包检测】和【重连机制】，很适合移动游戏环
		境使用

	消息体概述
		除去MQTT协议的实现，在实际游戏过程中我们可以只需要关注以下内容：MQTT协议消息体由两部分组成
		【topic】和【body】

	主题（topic）
		NQTT是通过主题对消息进行分类的，本质上就是一个UTF-8的字符串，不过可以通过反斜杠表示多个层级关系。
		主题并不需要创建，直接使用就是了。

		主题可以通过通配符进行过滤。其中，+可以过滤一个层级，而*只能出现在主题最后表示过滤任意级别的层级。
		举个例子：
			baidu/chatroom：代表百度公司的聊天室
			+/chatroom：代表任何公司的聊天室
			baidu/*：代表百度公司所有的子频道

	消息体（body）
		消息体是二进制数据流

	如何使用MQTT协议实现游戏路由
		由于mqant目前主要用于游戏开发，因此mqant只使用了mqtt协议的一小部分功能。

		mqant网关将接收到消息topic解析出moduleType和handler用来定位到后端模块的对应处理方法，然后进行远程
		RPC调用。msgid作为客户端是否需要消息应答的标记

		如下图

				Topic:Chat/HD_Join/20
				Body:{"roomName":"mqant","nick","liangdas"}

							——————————————————————》
			客户端												mqant网关
							《——————————————————————				 |
																  |		moduleType : Chat
				Topic:Chat/HD_Join/20							  |		fn : HD_Join
				Body:{"Error":"","Result":"join success"}		  |		agrs : 
																  |		Session,
																  |		{"roomName":"mqant","nick","liangdas"}
																  |
																后端模块
*/	