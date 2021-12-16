
// session
/*
	概述
		gate handler传参第一个参数是session gate.Session

	网关和后端通信
		在客户端通过网关向后端业务模块发送消息时，网关会将session放在第一个参数，
		因此再设计handler是需要将第一个参数定义为session
		
		func (self *HellWorld) gatesay(session gate.Session, msg map[string]interface{}) (r string err error) {
			session.Send("/gate/send/test", []byte(fmt.Sprintf("send hi to %v", msg["name"])))
			return fmt.Sprintf("hi %v 你在网关 %v", msg["name"], session.GetServerId()), nil
		}

	session定义
		Session是由Gate网关模块维护的，代表客户端跟网关建立的一条连接，session封装了网关和客户端连接的信息

	组成
		网关信息
		连接自定义信息

	它的大致字段如下：
		{
			Userid		string
			IP			string
			Network		string
			Sessionid	string
			Serverid	string
			Settings	<key-value map>
		}

		1.Userid
			需要调用Bind绑定来设置，默认为""当客户端登陆以后可以设置该参数，其他业务模块通过判断Userid来判断该
			连接是否合法
		2.IP
			客户端IP地址
		3.Network
			网络类型TCP websocket
		4.Sessionid
			Gate网关模块生成的该连接唯一ID
		5.Serverid
			Gate网关模块唯一ID，后端模块可以通过它来与Gate网关进行RPC调用
		6.Settings
			可以给这个连接设置一些参数，例如当用户假如对战方健以后可以设置一个参数roomName="mqant"

	功能
		session最大的功能是它封装了跟网关通信的方法，利用session我们可以给客户端发送消息

	给客户端发送消息
		遵守mqtt的主题(topic)、消息体(body)
			session.Send("/gate/send/test", []byte(fmt.Sprintf("send hi to %v", msg["name"])))
		不阻塞给客户端发消息
			session.SendNR("/gate/send/test", []byte(fmt.Sprintf("send hi to %v", msg["name"])))

	绑定uid
		连接建立后并没有任何用户属性，当用户通过该连接登陆后，可以将uid绑定到session上，这样session也就跟
		用户关联到一起了
			session.Bind("userId")

	设置一个远程属性
		session.SetPush(key, value)

	获取一个属性
		session.Get(key)
*/