
// 消息拦截
/*
	概述
		有些场景我们需要拦截或监听部分上行或者下行消息

	上行消息拦截
		利用自定义路由器实现(RouteHandler)

	实现自定义路由器
		func (this *Gate) OnRoute(session gate.Session, topic string, msg []byte) (bool, interface{}, error) {
			needreturn := true
			u, err != nil {
				return needreturn, nil, errors.Errorf("topic is not uri %v", err.Error())
			}
			if strings.HasPrefix(topic, "account://mudulus/init_by_deviceid") {
				// 拦截协议
				message := map[string]interface{}{}
				err := json.Unmarshal(msg, &message)
				if err != nil {
					return needreturn, "", err
				}
				r, errstr := this.init_by_deviceid(session, message)
				if errstr != "" {
					return needreturn, "", errors.Errorf(errstr)
				}
				return needreturn, r, nil
			} else {
				// 利用UriRoute将消息转发至后端模块，也可以实现自己的Route规则
				return this.Route.OnRoute(session, topic, msg)
			}
			return needreturn, "success", nil
		}

	替换默认的gate路由规则
		func (this *Gate) OnInit(app module.App, settings *conf.ModuleSettings) {
			this.Gate.OnInit(this, app, settings,
				gate.SetRouteHandler(OnRoute),
			)
		}

	下行消息拦截
		可以通过SendMessageHook替换下行消息体(Body)，也可以直接拦截不允许发送
		func (this *Gate) OnInit(app module.App, settings *conf.ModuleSettings) {
			this.Gate.OnInit(this, app, settings,
				gate.SetSendMessageHook(func(session gate.Session, topic string, msg []byte) (bytes []byte, e error) {
					return nil, errors.New("本协议不允许发送给客户端")
				}),
			)
		}
*/