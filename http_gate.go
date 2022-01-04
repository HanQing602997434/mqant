
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

*/