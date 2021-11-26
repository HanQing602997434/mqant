
// 自定义数据结构
/*
	介绍
		跟protocolbuffer类似，mqant也识别mqrpc.Marshaler接口实现的数据结构，开发者值需要自己实现序列化
		与反序列化即可

	Marshaler接口定义
		序列化
		func (this *mystruct)Marshal() ([]byte, error)

		反序列化
		func (this *mystruct)Unmarshal(data []byte) error

		数据结构名称
		func (this *mystruct)String() string

	注意事项
		1.mqrpc.Marshaler是请求方和服务方约定的数据结构，因此需要双方都能够明确数据结构的类型（可以直接断言的）

		2.服务函数返回结构一定要用指针（例如*rsp）否则mqant无法识别

	编写支持Marshaler传参的handler
		首先我们重新组织了一下代码目录结构，新增一个marshaler.go用来存放自定义数据结构代码
			工程目录
				|-bin
					|-conf
				|-helloworld
					|-module.go
				|-web
					|-module.go
				|-rpctest
					|-module.go
					|-marshaler.go
				|-main.go

	定义数据结构
		package rpctest

		// 请求数据结构
		type Rep struct {
			Id string
		}

		func (this *Req) Marshal() ([]byte, error) {
			return []byte(this.Id), nil
		}

		func (this *Req) Unmarshal(data []byte) error {
			this.Id = string(data)
			return nil
		}

		func (this *Req) String() string {
			return "req"
		}

		// 响应数据结构
		type Rsp struct {
			Msg string
		}

		func (this *Rsp) Marshal() ([]byte, error) {
			return []byte(this.Msg), nil
		}

		func (this *Rsp) Unmarshal(data []byte) error {
			this.Msg = string(data)
			return nil
		}

		func (this *Rsp) String() string {
			return "rsp"
		}

	新增handler
		self.GetServer().RegisterGO("/test/marshal", self.testMarshal)

		...

		func (self *rpctest) testMarshal(req Req) (*Rsp, error) {
			r := &Rsp{Msg: fmt.Sprintf("你的ID: %v", req.Id)}
			return r, nil
		}

	调用marshaler的handler
		我们依然在web模块中新加一个api来测试
		http.HandleFunc("/test/marshal", func(w http.ResponseWriter, r *http.Request) {
			_ = r.ParseForm()
			ctx, _ := context.WithTimeout(context.TODO, time.Second*3)
			rspbean := new(rpctest.Rsp)
			err := mqrpc.Marshal(rspbean, func() (reply interface{}, errstr interface{}) {
				return self.Call(
					ctx,
					"rpctest", // 要访问的moduleType
					"/test/marshal", // 访问模块中的handler路径
					mqrpc.Param(&rpctest.Req{Id: r.Form.Get("mid")}),
				)
			})
			log.Info("RpcCall %v, err %v", rspbean, err)
			if err != nil {
				_, _ = io.WriteString(w, err.Error())
			}
			_, _ = io.WriteString(w, rspbean.Msg)
		})

	测试
		http://127.0.0.1:8080/test/marshal?mid=1314

	结果
		你的ID：1314
*/