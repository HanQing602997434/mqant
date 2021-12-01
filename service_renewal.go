
// 服务续约
/*
	基本原理
		服务在启动时注册服务发现，并在关闭时取消注册。有时这些服务可能会意外死亡，被强行杀死或面临暂时
		的网络问题。在这些情况下，遗留的节点将存在服务发现中。

	解决办法
		服务发现支持注册的TTL选项和间隔。TTL指定在发现之后注册的信息存在多长时间，然后过期并被删除。
		时间间隔是服务应该重新注册的时间，以保留其在服务发现中的注册信息。

	设置
		mqant默认的ttl=20，重新注册间隔为10秒

		func (self *rpctest) OnInit(app module.App, settings *conf.ModuleSettings) {
			self.BaseModule.OnInit(self, app, settings,
				server.RegisterInterval(15*time.Second),
				server.RegisterTTL(30*time.Second),
			)
		}
		以上这个例子，我们设置了一个30秒的ttl，重新注册间隔为15秒。如果应用进程被强杀，服务未取消注册，则
		30秒内其他服务节点无法感知节点已失效。
*/