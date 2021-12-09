
// 动态监听handler
/*
	概述
		有些场景下，我们无法在编译阶段提前实现或注册好所有的handler，但在执行时可以通过
		一些动态规则动态分配handler。mqant也支持这样的功能场景

	handler监听器
		type RPCListener interface {
			// NoFoundFunction当未找到请求的handler时会触发该方法
			// FunctionInfo选择可执行的handler
			// return error
			
			NoFoundFunction(fn string) (*FunctionInfo, error)
			
			// BeforeHandle会对请求做一些前置处理，如：检查当前玩家是否已登录，打印统计日志等。
			// @session 可能为nil
			// return error 当error不为nil时将直接返回错误信息而不会再执行后续调用

			BeforeHandle(fn string, callInfo *CallInfo) error
			OnTimeOut(fn string, Expired int64)
			OnError(fn string, callInfo *CallInfo, err error)

			// fn 方法名
			// params 参数
			// result 执行结果
			// exec_time 方法执行时间 单位为Nano 纳秒 1000000纳秒等于1毫秒

			OnComplete(fn string, callInfo *CallInfo, result *rpcpb.ResultInfo, exec_time int64)
		}

	设置监听器
		func (self *HttpGateWay) OnInit(app module.App, settings *conf.ModuleSettings) {
			self.SetListener(self)
		}

	示例
		我们实现一个http网关路由示例，将handler转换为http请求的path路由

	监听器实现
		func (self *HttpGateWay) NoFoundFunction(fn string)(*mqrpc.FunctionInfo, error) {
			return &mqrpc.FunctionInfo{
				Function:reflect.ValueOf(self.CloudFunction),
				Goroutine:true,
			}, nil
		}

		func (self *HttpGateWay) BeforeHandle(fn string, callInfo *mqrpc.CallInfo) error {
			return nil
		}

		func (self *HttpGateWay) OnTimeOut(fn string, Expired int64) {

		}

		func (self *HttpGateWay) OnError(fn string, callInfo *mqrpc.CallInfo, err error){}
		
		// fn 方法名
		// params 参数
		// result 执行结果
		// exec_time 方法执行时间 单位为Nano纳秒 1000000纳秒等于1毫秒

		func (self *HttpGateWay) OnComplete(fn string, callInfo *mqrpc.CallInfo, result *rpcpb.ResultInfo, exec_time int64){}

	请求转发器实现
		以下是一段伪代码

		1.监听http网关的请求
		2.解析http的path(url)
		3.填充http请求参数
		4.通过httptest模拟http请求
		5.将结果返回http网关

			func (self *HttpGateWay) CloudFunction(trace log.TraceSpan, request *go_api.Request) (*go_api.Response, error) {
				e := echo.New()
				ectest := httgatewaycontrollers.SetupRouter(self, e)
				req, err := http.NewRequest(request.Method, request.Url, strings.NewReader(request.Body))
				if err != nil {
					return nil, err
				}

				for _, v := range request.Header {
					req.Header.Set(v.Key, strings.Join(v.Values, ","))
				}

				rr := httptest.NewRecorder()
				ectest.ServerHTTP(rr, req)
				resp := &go_api.Response{
					StatusCode:int32(rr.Code),
					Body:rr.Body.String(),
					Header:make(map[string]*go_api.Pair),
				}
				for key, vals := range rr.Header() {
					header, ok := resp.Header[key]
					if !ok {
						header = &go_api.Pair{
							Key: key,
						}
						resp.Header[key] = header
					}
					header.Values = vals
				}
				return resp, nil
			}
*/