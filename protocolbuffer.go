
// protocolbuffer
/*
	介绍
		protocolbuffer结构体（推荐）

	注意事项
		1.proto.Message是protocol buffer约定的数据结构，因此需要双方都能够明确数据结构的类型（可以直接断言）

		2.服务函数返回结构一定要用指针（例如*rpcpb.ResultInfo）否则mqant无法识别

	代码组织结构
		首先我们重新组织了一下代码目录结构，新增了一个rpctest目录用来存放rpctest模块代码

		工程目录
			|-bin
				|-conf
					|-server.conf
			|-helloworld
				|-module.go
			|-web
				|-module.go
			|-rpctest
				|-module.go
			|-main.go
		
	编写支持pb传参的handler
		为了简化操作，我们直接使用mqant内部的protocolbuffer结构体rpcpb.ResultInfo

		var Module = func() module.Module {
			this := new(rpctest)
			return this
		}

		type rpctest struct {
			basemodule.BaseModule
		}

		func (self *rpctest) GetType() string {
			// 很关键，需要配置文件中的Module配置对应
			return "rpctest"
		}

		func (self *rpctest) Version() string {
			// 可以在监控时了解代码版本
			return "1.0.0"
		}

		func (self *rpctest) OnInit(app module.App, settings *conf.ModuleSettings) {
			self.BaseModule.OnInit(self, app, settings)
			self.GetServer().RegisterGO("/test/proto", self.testProto)
		}

		func (self *rpctest) Run(closeSig chan bool) {
			log.Info("%v模块运行中...", self.GetType())
			<-closeSig
		}

		func (self *rpctest) OnDestory() {
			// 一定别忘了继承
			self.BaseModule.OnDestory()
		}

		func (self *rpctest) testProto(req *rpcpb.ResultInfo) (*rpcpb.ResultInfo, error) {
			r := &rpcpb.ResultInfo{Error: *proto.String(fmt.Sprintf("你说：%v", req.Error))}
			return r, nil
		}

	调用pb的handler

		我们依然在web模块中新加一个api来测试

		http.HandleFunc("/test/proto", func(w http.ResponseWriter, r *http.Request) {
			_ = r.ParseForm()
			ctx, _ := context.WithTimeout(context.TODO(), time.Second*3)
			protobean := new(rpcpb.ResultInfo)
			err := mqrpc.Proto(protobean, func() (reply interface{}, errstr interface{}) {
				return self.RpcCall(
					ctx,
					"rpctest", // 要访问的moduleType
					"/test/proto", // 访问模块中的handler路径
					mqrpc.Param(&rpcpb.ResultInfo{Error: *proto.String(r.Form.Get("message"))}),
				)
			})
			log.Info("RpcCall %v, err %v", protobean, err)
			if err != nil {
				_, _ = io.WriteString(w, err.Error())
			}
			_, _ = io.WriteString(w, protobean.Error)
		})

	测试
		http://127.0.0.1:8080/test/proto?message=this%20is%20the%20protocolbuffer%20test

	结果
		你说：this is the protocolbuffer test
*/