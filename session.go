
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

	同步网关最新数据
		通常session.SetPush，session.SetBatch也会同步最新数据
			session.Push() // 如果想同步网关最新数据

	批量设置远程属性
		没调用一次session.SetPush就会进行一次RPC通信，如果有多个属性需要设置的话可以合并批量设置
			session.SetBatch(map[string]string)

	判断session是否游客
		session.IsGuest()

	自定义游客判断规则
		func main() {
			gate.JudgeGuest = func(session gate.Session) bool {
				if session.GetUserId() != "" {
					return false
				}
				return true
			}
		}

	特性
		session最新的信息都由网关维护，也就是说只要不是最新从网关同步来的session，信息都有可能是
		不完整或者错误的。
			1.时效性session跟客户端连接直接挂钩，如果连接中断session即失效
			2.不确定性 如果客户端连接中断重连，可能从网关1切换成网关2，那么以前持有的session则失效
			3.可序列化

		合理利用session可以提高

	序列化
		bytes, err := session.Serializable()
		
		app module.App
		session, err := basegate.NewSession(app, bytes)

	持久化
		如果我们不希望客户端网络中断以后导致Session自定义数据丢失，以保证客户端在指定时间内重连以后
		能继续使用这些数据，我们需要对用户的Session进行持久化

		Gate网关模块目前提供一个接口用来控制Session的持久化，但具体的持久化方式需要开发者自己来实现
	
	session监听器
		type StorageHandler interface {
			存储用户的Session信息，触发条件
			1.session.Bind(Userid)
			2.session.SetPush(key, value)
			3.session.SetBatch(map[string]string)

			Storage(session Session) (err error)

			强制删除Session信息，触发条件
			1.暂无
			Delete(session Session) (err error)
			
			获取用户Session信息，触发条件
			1.session.Bind(Userid)

			Query(Userid string) (data []byte, err error)

			用户心跳，触发条件
			1.每一次客户端心跳

			可以用来延长Session信息过期时间
			Heartbest(session Session)
		}

	设置监听器
		func (this *Gate) OnInit(app module.App, settings *conf.ModuleSettings) {
			// 注意这里一定要用gate.Gate 而不是 module.BaseModule
			this.Gate.OnInit(this, app, settings,
				gate.SetStorageHandler(this),
			)
		}
*/