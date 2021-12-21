
// 自定义网关路由
/*
	topic uri路由器
		mqant gate网关的默认路由规则为
			[moduleType@moduleID]/[handler]/][msgid]
		
		但是这个规则不太灵活，因此设计了一套基于URI规则的topic路由规则

	基于uri协议的路由规则
		<shceme>://<user>:<password>@<host>:<port>/<path>;<params>?<query>#<fragment>

		1.可以充分利用uri公共库
		2.资源划分更加清晰明确

	如何启用模块

	创建一个UriRoute结构体
		// 注意这里一定要用 gate.Gate 而不是 module.BaseModule
		this.Route = uriroute.NewURIRoute(this,
			uriroute.Selector(func(topic string, u *url.URL) (s module.ServerSession, err error) {
				moduleType := u.Scheme
				nodeId := u.Hostname()
				// 使用自己的
				if nodeId == "modulus" {
					// 取模
				} else if nodeId == "cache" {
					// 缓存
				} else if nodeId == "random" {
					// 随机
				} else {
					//
					// 指定节点规则就是 module://[user:pass@]nodeId/path
					// 方式1
					// moduleType = fmt.Sprintf("%v@%v", moduleType, u.Hostname())
					// 方式2
					return this.GetRouteServer(moduleType, selector.WithFilter(selector.FilterEndpoint(nodeId)))
				}
				return this.GetRouteServer(moduleType)
			}),
			uriroute.DataParsing(func(topic string, u *url.URL, msg []byte) (bean interface{}, err error) {
				// 根据topic解析msg为指定的结构体
				// 结构体必须满足RPC的参数传递标准
				return
			}),
			uriroute.CallTimeOut(3 * time.Second),
		)

	实现自定义路由器
		func (this *Gate) OnRoute(session gate.Session, topic string, msg []byte) (bool, interface{}, error) {
			needreturn := true
			u, err := url.Parse(topic)
			if err != nil {
				return needreturn, nil, errors.Errorf("topic is not uri %v", err.Error())
			}
			if strings.HasPrefix(topic, "account://modulus/init_by_deviceid") {
				// 拦截协议
				message := map[string]interface{}{}
				err := json.Unmarsha1(msg, &message)
				if err != nil {
					return needreturn, "", err
				}
			}
		}

	替换默认的gate路由规则
		func (this *Gate) OnInit(app module.App, settings *conf.ModuleSettings) {
			this.Gate.OnInit(this, app, settings,
				gate.SetRouteHandler(OnRoute)
			)
		}
*/