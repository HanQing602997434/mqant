
// RPC
/*
	mqant RPC本身是一个相对独立的功能，RPC有以下的几个特点：
		1.目前支持nats作为服务发现通道，理论上可以扩展其他通信方式
		2.支持服务注册发现，是一个相对完善的微服务框架

	在模块中使用RPC
		module.BaseModule中已经集成了RPC，使用方式如下

	服务提供者
		注册handler

		注册服务函数
		module.GetServer().RegisterGO(_func string, fn interface{})

		注册服务函数
		module.GetServer().Register(_func string, fn interface{})

		RegisterGO和Register的区别是前者为每一条消息创建一个单独的协程来处理，后者注册的函数共用一个
		协程来处理所有消息，具体使用哪一种方式可以根据实际情况来定，但Register方式的函数一定注意不要
		执行耗时功能，以免引起消息阻塞

	服务调用者
		在开发过程中，模块A可能需要用到模块B的服务，这时模块A就成为了服务调用方。mqant提供了多种RPC
		调用方法，也支持高级扩展（服务发现）

	RPC路由规则
		mqant每一类模块可以部署到多台服务器中，因此需要一个nodeId对同一类模块进行区分。在框架假如服务
		注册和发现功能后，nodeId通过服务发现模块在服务启动时自动生成，无法提前编程指定。

	RPC调用方法介绍
		1.通过Call函数调度（推荐）
			通用RPC调度函数
			ctx			context.Context		上下文，可以设置这次请求的超时时间
			moduleType	string				服务名称 serverId 或 serverId@nodeId
			_func		string				需要调度的服务方法
			param		mqrpc.ParamOption	方法传参
			opts ...selector.SelectOption	服务发现模块过滤，可以用来选择调用哪个服务节点
			Call(ctx context.Context, moduleType, _func string, param mqrpc.ParamOption, opts ...selector.SelectOption) (interface{}, string)

			特点
				支持设置调用超时时间
				支持自定义的服务节点选择过滤器

			超时时间设置
				ctx, _ := context.WithTimeout(context.TODO(), time.Second * 3) // 3s超时
				rstr, err := mqrpc.String(
					self.Call(
						ctx,
						"helloworld", // 要访问的moduleType
						"/say/hi", // 访问模块中handler路径
						mqrpc.Param(r.Form.Get("name")),
					),
				)
				超时时间仅是调用方有效，超时后无法取消被调用方正在执行的任务。

			服务节点选择过滤器
				伪代码
				ctx, _ := context.WithTimeout(context.TODO(), time.Second * 3)
				rstr, err := mqrpc.String(
					self.Call(
						ctx,
						"helloworld", // 要访问的moduleType
						"/say/hi", // 访问模块中handler路径
						mqrpc.Param(r.Form.Get("name")),
						selector.WithStrategy(func(services []*registry.Service) selector.Next {
							var nodes []*registry.Node

							// Filter the nodes for datacenter
							for _, service := range services {
								for _, node := range service.Nodes {
									if node.Metadata["version"] == "1.0.0" {
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
								return node[index], nil
							}
						}),
					),
				)

			moduleType的格式
				指定到模块级别
					当moduleType为模块名时func GetType()值一样，rpc将查找模块已启动的所有节点，然后根据【节点
					选择过滤器】选择一个节点发起调用

				指定到节点级别
					格式为moduleType@moduleID 例如 helloworld@1b0073cbbab33247，rpc将直接选择节点
					1b0073cbbab33247发起调用
		
		2.通过Rpclnvoke函数调度
			module.Invoke(moduleType string, _func string, params ...interface{})

			特点
				不支持设置调用超时时间（只能通过配置文件设置全局RPC超时时间）
				不支持自定义的服务节点选择过滤器
				支持moduleType过滤

		3.通过InvokeNR函数调度
			module.InvokeNR(moduleType string, _func string, params ...interface{})

			特点
				包含Invoke所有特点
				本函数无需等待返回结果（不会阻塞），仅投递RPC消息

		4.指定节点调用
			查找到节点（module.ServerSession），通过节点结构体提供的方法调用
			moduleType 模块名称（类型）
			opts	   服务节点选择过滤器
			func GetRouteServer(moduleType string, opts ...selector.SelectOption) (s module.ServerSession, err error)

			SvrSession, err := self.GetRouteServer("helloworld",
			selector.WithStrategy(func(services []*registry.Service) selector.Next {
				var nodes []*registry.Node

				// Filter the nodes for datacenter
				for _, service := range services {
					for _, node := range service.Nodes {
						if node.Metadata["version"] == "1.0.0" {
							nodes = append(nodes, node)
						}
					}
				}

				var mtx sync.Mutex
				// log.Info("services[0] &v", servoces[0].Nodes[0])
				return func() (*registry.Node, error) {
					mtx.Lock()
					defer mtx.Unlock()
					if len(nodes) == 0 {
						return nil, fmt.Errorf("no node")
					}
					index := rand.Intn(int(len(nodes)))
					return nodes[index], nil
				}
			}), )
			if err != nil {
				log.Warning("HelloWorld error:%v", err.Error())
				return
			}

			rstr, err := mqrpc.String(SvrSession.Call(ctx, "/say/hi", r.Form.Get("name")))
			if err != nil {
				log.Warning("HelloWorld error:%v", err)
				return
			}

	以上的调用方法在module级别和app级别都有对应实现，可灵活选择


	RPC传参数据结构
		RPC可传参数据类型
			1-9为基础数据类型，可直接使用。10、11为自定义结构体，需要单独定义（章节后续会单独讲解）
			
			1.bool
			2.int32
			3.int64
			4.long64
			5.float32
			6.float64
			7.[]byte
			8.string
			9.map[string]interface{}
			10.protocol buffer结构体
			11.自定义结构体
			
			注意调用参数不能为nil 如： result, err = module.Invoke("user", "login", "mqant", nil)会出现异常无法调用

		返回值可使用的参数类型
			handler的返回值固定为两个，其中result表示正常业务返回值，err表示异常业务返回值

			result:
				1.bool
				2.int32
				3.int64
				4.long64
				5.float32
				6.float64
				7.[]byte
				8.string
				9.map[string]interface{}
				10.protocol buffer结构体
				11.自定义结构体
			
			err:
				1.string
				2.error

			示例
				func (self *HelloWorld)say(name string) (result string, err error) {
					return fmt.Sprintf("hi %v", name), nil
				}

				result, err := mqrpc.String(
					self.Call(
						ctx,
						"helloworld", // 要访问的moduleType
						"/say/hi", // 访问模块中的handler路径
						mqrpc.Param(r.Form.Get("name")),
					)
				)
*/