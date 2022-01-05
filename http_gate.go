
// http网关介绍
/*
	概述
		mqant实现了一个比较简单得http网关，可以用来代理前端发起的http api请求

	使用http网关
		mqant提供了一个创建代理网关的http.Handler

	代码组织结构
		工程目录
        |-bin
            |-conf
                |-server.conf
        |-httpgateway
            |-module.go
		|-main.go
		
	初始化网关监听
		var Module = func() module.Module {
    		this := new(httpgate)
    		return this
		}

		type httpgate struct {
    		basemodule.BaseModule
		}

		func (self *httpgate) GetType() string {
    		//很关键,需要与配置文件中的Module配置对应
    		return "httpgate"
		}
		func (self *httpgate) Version() string {
    		//可以在监控时了解代码版本
    		return "1.0.0"
		}
		func (self *httpgate) OnInit(app module.App, settings *conf.ModuleSettings) {
    		self.BaseModule.OnInit(self, app, settings)
    		self.SetListener(self)
		}

		func (self *httpgate) startHttpServer() *http.Server {
    		srv := &http.Server{
        		Addr: ":8090",
        		Handler:httpgateway.NewHandler(self.App),
    		}
    		//http.Handle("/", httpgateway.NewHandler(self.App))

    		go func() {
        		if err := srv.ListenAndServe(); err != nil {
            		// cannot panic, because this probably is an intentional close
            		log.Info("Httpserver: ListenAndServe() error: %s", err)
        		}
    		}()
    		// returning reference so caller can call Shutdown()
    		return srv
		}

		func (self *httpgate) Run(closeSig chan bool) {
    		log.Info("httpgate: starting HTTP server :8090")
    		srv := self.startHttpServer()
    		<-closeSig
    		log.Info("httpgate: stopping HTTP server")
    		// now close the server gracefully ("shutdown")
    		// timeout could be given instead of nil as a https://golang.org/pkg/context/
    		if err := srv.Shutdown(nil); err != nil {
        		panic(err) // failure/timeout shutting down the server gracefully
    		}
    		log.Info("httpgate: done. exiting")
		}

		func (self *httpgate) OnDestroy() {
    		//一定别忘了继承
    		self.BaseModule.OnDestroy()
		}

	创建网关http.Handler
		httpgateway.NewHandler(self.App)

	网关默认路径规则
		网关默认路由规则是从URL.Path的第一个段取出moduleType
		/[moduleType]/path

	举例
		http://127.0.0.1:8090/httpgate/topic
			moduleType httpgate
			handler /httpgate/topic

	实现后端http协议
		网关转发RPC的handler定义为
			func (self *httpgate) httpgateway(request *go_api.Request) (*go_api.Response, error) {}

	编写handler
		func (self *httpgate) httpgateway(request *go_api.Request) (*go_api.Response,error) {
    		mux := http.NewServeMux()
    		mux.HandleFunc("/httpgate/topic", func(writer http.ResponseWriter, request *http.Request) {
        		writer.Write([]byte(`hello world`))
    		})

    		req, err := http.NewRequest(request.Method, request.Url, strings.NewReader(request.Body))
    		if err != nil {
        		return nil,err
    		}
    		for _,v:=range request.Header{
        		req.Header.Set(v.Key, strings.Join(v.Values,","))
    		}
    		rr := httptest.NewRecorder()
    		mux.ServeHTTP(rr, req)
    		resp := &go_api.Response{
        		StatusCode:  int32(rr.Code),
        		Body: rr.Body.String(),
        		Header: make(map[string]*go_api.Pair),
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

	注册handler(方案一)
		self.GetServer().RegisterGO("/httpgate/topic", self.httpgateway)

	注册handler(方案二)
		方案一比较固化，跟mqant框架绑定，我们希望在编写http函数时能利用golang已存在的web库，又能跟
		mqant网关结合

		利用mqrpc, RPCListener监听未实现的handler，然后将请求通过httptest路由web框架中

		当RPC未找到已注册的handler时会调用func NoFoundFunction(fn string)(*mqrpc.FunctionInfo, error)
		func (self *httpgate) NoFoundFunction(fn string)(*mqrpc.FunctionInfo,error){
    		return &mqrpc.FunctionInfo{
        		Function:reflect.ValueOf(self.httpgateway),
        		Goroutine:true,
   			},nil
		}
		func (self *httpgate) BeforeHandle(fn string, callInfo *mqrpc.CallInfo) error{
    		return nil
		}
		func (self *httpgate) OnTimeOut(fn string, Expired int64){

		}
		func (self *httpgate) OnError(fn string, callInfo *mqrpc.CallInfo, err error){}
		func (self *httpgate) OnComplete(fn string, callInfo *mqrpc.CallInfo, result *rpcpb.ResultInfo, exec_time int64){}

	设置handler监听器
		func (self *httpgate) OnInit(app module.App, settings *conf.ModuleSettings) {
			self.BaseModule.OnInit(self, app, settings)
			self.SetListener(self)
		}

	运行
		2020-05-06T16:09:19.98671+08:00 [-] [-] [development] [I] [rpc_server.go:142] Registering node: httpgate@4bfaf93bd31c512a
		2020-05-06T16:09:19.988701+08:00 [-] [-] [development] [I] [ws_server_x.go:131] WS Listen ::3653
		2020-05-06T16:09:19.989593+08:00 [-] [-] [development] [I] [tcp_server.go:39] TCP Listen ::3563
		2020-05-06T16:09:19.995531+08:00 [-] [-] [development] [I] [module.go:62] httpgate: starting HTTP server :8090

	访问
		http://127.0.0.1:8090/httpgate/topic

	结果
		hello world

	用gin替代默认http框架
		如上的httpgateway函数中，网络框架可以用gin, echo, beego等其他web框架替代
*/