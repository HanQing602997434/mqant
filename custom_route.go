
// 网关自定义路由
/*
	概述
		网关默认路由规则可能不满足业务场景的路由需求，可以自定义

	网关默认路径规则
		网关默认路由规则是从URL.Path的第一个段取出moduleType
		/[moduleType]/path

	举例
		http://127.0.0.1:8090/httpgate/topic
			moduleType httpgate
			handler /httpgate/topic

	编写自定义路由规则器
		srv := &http.Server{
			Addr : ":8090",
			Handler:httpgateway.NewHandler(self.App,
				httpgateway.SetRoute(func(app module.App, r *http.Request) (service *httpgateway.Service, e error) {
					return nil, nil
				})
			),
		}

	Service
		type Service struct {
			// hander
			Hander string
			// mode
			SrvSession module.ServerSession
		}

	ServerSession
		可以通过app.GetRouteServer函数获取
		session, err := app.GetRouteServer
*/