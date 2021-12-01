
// 服务元数据
/*
	介绍
		服务还支持设置元数据，即服务得自身属性，通过这些属性我们可以定制服务发现策略

	节点ID
		一般情况下节点ID在模块初始化时由系统自动生成一个不重复的随机数，但我们也可以指定节点ID

		func (self *rpctest) OnInit(app module.App, settings *conf.ModuleSettings) {
			self.BaseModule.OnInit(self, app, settings,
				server.RegisterInterval(15*time.Second),
				server.RegisterTTL(30*time.Second),
				server.Id("mynode001"),
			)
		}

	启动
		2020-05-05T11:06:35.224305+08:00 [-] [-] [development] [I] [rpc_server.go:142] Registering node: Web@a6a49427222f84c4
		2020-05-05T11:06:35.224656+08:00 [-] [-] [development] [I] [module.go:39] HelloWorld模块运行中...
		2020-05-05T11:06:35.232826+08:00 [-] [-] [development] [I] [rpc_server.go:142] Registering node: rpctest@mynode001
		2020-05-05T11:06:35.233042+08:00 [-] [-] [development] [I] [module.go:179] web: starting HTTP server :8080
		2020-05-05T11:06:35.241667+08:00 [-] [-] [development] [I] [module.go:46] rpctest模块运行中...

	调用
		如果明确知道节点ID，那你可以这样直接找到它，虽然通常不这样用

		err := mqrpc.Marshal(rspbean, func() (reply interface{}, errstr interface{}) {
			return self.Call(
				ctx,
				"rpctest@mynode001", // 要访问的moduleType
				"/test/marshal", // 访问模块中handler路径
				mqrpc.Param(&rpctest.Req{Id: r.Form.Get("mid")}),
			)
		})

	服务版本（Version）
		模块（服务）启动时，会自动注册模块func Version() string的返回值作为服务的版本
		如果你愿意可以利用服务版本过滤节点

		rstr, err := mqrpc.String(
			self.Call(
				ctx,
				"helloworld", // 要访问的moduleType
				"/say/hi", // 访问模块中的handler路径
				mqrpc.Param(r.Form.Get("name")),
				selector.WithStrategy(func(services []*registry.Service) selector.Next {
					var nodes []*registry.Node

					// Filter the nodes for datacenter
					for _, service := range services {
						if service.Version != "1.0.0" {
							continue
						}
						for _, node := range service.Nodes {
							nodes = append(nodes, node)
						}
					}

					var mtx sync.Mutex
					// log.Info("services[0] $v", services[0].Nodes[0])
					return func() (*registry.Node, error) {
						mtx.Lock()
						defer mtx.Unlock()
						if len(nodes) == 0 {
							return nil, fmt.Error("no node")
						}
						index := rand.Intn(int(len(nodes)))
						return nodes[index], nil
					}
				}),
			),
		)

	元数据（Metadata）
		你还可以为服务节点指定设置元数据，元数据是节点级别的，且可以随时修改，利用好它可以灵活的实现定制
		化的服务发现，比如实现灰度发布，熔断策略等等

	设置元数据
		self.GetServer().Options().Metadata["state"] = "alive"

	立即刷新
		设置好的元数据会等到下一次重新注册是更新到配置中心并同步至其他节点，如果我们想立即生效的话可以
		这样做
			_ := self.GetServer().ServiceRegister()

	示例
		rstr, err := mqrpc.String(
			self.Call(
				ctx,
				"helloworld", // 要访问的moduleType
				"/say/hi", // 访问模块中的handler路径
				mqrpc.Param(r.Form.Get("name")),
				selector.WithStrategy(func(services []*registry.Service) selector.Next {
					var nodes []*registry.Node

					// Filter the nodes for datacenter
					for _, service := range services {
						if service.Version != "1.0.0" {
							continue
						}
						for _, node := range service.Nodes {
							nodes = append(nodes, node)
							if node.Metadata["state"] == "alive" || node.Metadata["state"] == "" {
								nodes = append(nodes, node)
							}
						}
					}

					var mtx sync.Mutex
					// log.Info("services[0] $v", services[0].Nodes[0])
					return func() (*registry.Node, error) {
						mtx.Lock()
						defer mtx.Unlock()
						if len(nodes) == 0 {
							return nil, fmt.Errorf("no node")
						}
						index := rand.Intn(int(len(nodes)))
						return nodes[index], nil
					}
				}),
			),
		)
*/