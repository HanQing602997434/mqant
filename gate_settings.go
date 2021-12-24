
// 网关可选设置
/*
	最大同时任务数
		为了防止客户端上次消息太过频繁，影响网关性能，可以设置一个连接的最大同时任务书
			默认值
				func (this *Gate) OnInit(app module.App, settings *conf.ModuleSettings) {
					this.Gate.OnInit(this, app, settings,
					gate.ConcurrentTasks(s),
					)
				}
			如果连接最大任务数超过限制，会报the work queue is full!错误，此时你需要关注客户端消息
			频率以及后端任务响应耗时是否合理

	网络读写缓存大小
		缓存分为读缓存和写缓存，缓存大小设置需要根据具体业务场景而定，如缓存设置过大对服务器内存消耗
		会比较多，如设置过小则可能导致读写卡顿，写数据时如缓存已满，则数据包会丢弃。
			
			默认值2048
			func (this *Gate) OnInit(app module.App, settings *conf.ModuleSettings) {
				this.Gate.OnInit(this, app, settings,
					gate.BufSize(1024),
				)
			}

	数据包最大长度
		为防止解包错误或恶意攻击，导致服务器内存溢出，需要限制每一个数据包的最大长度，超过最大长度限制
		的连接将被断开
			
			默认值65535
			func (this *Gate) OnInit(app module.App, settings *conf.ModuleSettings) {
				this.Gate.OnInit(this, app, settings,
					gate.Heartbeat(1*time.Minute)
				)
			}

	建连超时时间
		为防止连接到网关又为建立mqtt协议的异常连接，我们设置一个超时机制，客户端连接到网关以后需要在设置
		时间内完成mqtt协议建立，否则连接将被断开
			
			默认值 time.Second * 10
			func (this *Gate) OnInit(app module.App, settings *conf.ModuleSettings) {
				this.Gate.OnInit(this, app, settings,
					gate.OverTime(20 * time.Second),
				)
			}

	建立TLS加密通信
		tcp和ws都可以建立安全的加密通信(tls)
		
			默认 不安全通信
			func (this *Gate) OnInit(app module.App, settings *conf.ModuleSettings) {
				this.Gate.OnInit(this, app, settings,
					gate.Tls(true),
					gate.CertFile("xxx.cert"),
					gate.KeyFile("xxx.key"),
				)
			}
*/