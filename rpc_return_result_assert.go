
// RPC返回结果断言
/*
	概述：
		由于RpcCall是一个通用函数，我们无法对其返回值指定类型，为了简化代码，mqant参考redis封装了几个RPC返回类型断言
		函数，方便开发者使用

	断言函数介绍
		protocolbuffer断言
			protobean := new(rpcpb.ResultInfo)
			err := mqrpc.Proto(protobean, func() (reply interface{}, errstr interface{}) {
				return self.Call(
					ctx,
					"rpctest", // 要访问moduleType
					"/test/proto", // 访问模块中的handler路径
					mqrpc.Param(&rpcpb.ResultInfo{Error: *proto.String(r.Form.Get("message"))}),
				)
			})
			log.Info("RpcCall %v, err %v", protobean, err)

		自定义结构断言
			rspbean := new(rpctest.Rsp)
			err := mqrpc.Marshal(rspbean, func() (reply interface{}, errstr interface{}) {
				return self.Call(
					ctx,
					"rpctest", // 要访问得moduleType
					"/test/marshal", // 访问模块中的handler路径
					mqrpc.Param(&rpctest.Req{Id: "hello 我是RpcInvoke"}),
				)
			})
			log.Info("RpcCall %v, err %v", protobean, err)

		字符串断言
			rstr, err := mqrpc.String(
				self.Call(
					context.Background(),
					"helloworld",
					"/say/hi",
					mqrpc.Param(r.Form.Get("name")),
				),
			)
			log.Info("RpcCall %v, err %v", rstr, err)

		其他类型断言
			int

			bool

			map\[string\]string

			...
*/