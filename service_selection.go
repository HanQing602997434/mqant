
// 服务选择
/*
	概述
		紧跟上一章内容，微服务中每一个服务都会部署多个节点，并且根据实际情况我们还可能临时新增或摘除
		节点，通常节点选择是结合业务而定的，因此灵活的节点选择器是框架必备的功能

	用法
		mqant的节点选择器（selector）是从go-micro移植而来的，其使用规则可参考go-micro实现

	默认选择器
		mqant默认的选择器是一个随机负载均衡选择器

	RPC级别
		如果需要针对某一个RPC调用定制选择器可以这样做
		RpcCall函数可选参数中支持设置选择器

		rstr, err := mqrpc.String(
			self.RpcCall(
				ctx,
				"helloworld", // 要访问的moduleType
				"/say/hi", // 访问模块中handler路径
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

	应用级别
		大部分情况下，我们只需要定制一个全局统一的通用选择器，那么可以在应用（app）级别设置

		app := mqant.CreateApp(
			module.Debug(false),
		)

		_ = app.Options().Selector.Init(selector.SetStrategy(func(services []*registry.Service) selector.Next {
			var nodes []WeightNode

			// Filter the nodes for datacenter
			for _, service := range services {
				for _, node := range service.Nodes {
					weight := 100
					if w, ok := node.Metadata["weight"]; ok {
						wint, err := strconv.Atoi(w)
						if err == nil {
							weight = wint
						}
					}
					if state, ok := node.Metadata["state"]; ok {
						if state != "forbidden" {
							nodes = append(nodes, WeightNode {
								Node: node
								Weight: weight
							})
						}
					} else {
						nodes = append(nodes, WeightNode{
							Node: node,
							Weight: weight,
						})
					}
				}
			}
			//log.Info("services[0] $v", services[0].Nodes[0])
			return func() (*registry.Node, error) {
				if len(nodes) == 0 {
					return nil, fmt.Errorf("no node")
				}
				rand.Seed(time.Now().UnixNano())
				// 按权重选
				total := 0
				for _, n := range nodes {
					total += n.Weight
				}
				if total > 0 {
					weight := rand.Intn(total)
					togo := 0
					for _, a := range nodes {
						if (togo <= weight) && (weight < (togo + a.Weight)) {
							return a.Node, nil
						} else {
							togo += a.Weight
						}
					}
				}
				// 降级为随机
				index := rand.Intn(int(len(nodes)))
				return nodes[index].Node, nil
			}
		}))

		以上的选择器利用节点元数据（Metadata）定制了一个节点选择规则
			按权重（weight）
			按节点当前状态（forbidden）
			最后降级为随机
*/