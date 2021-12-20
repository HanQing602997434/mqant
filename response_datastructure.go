
// Response数据结构
/*
	概述
		在Request-Response模式下，网关需要将后端模块返回的数据在返回给客户端即Response，
		此时需要网关与客户端约定一套数据结构，方便客户端解析

		客户端与网关是长连接，因此不同的mqtt包会交叉传递

	示例时序图
		略

	主题（topic）
		Response时，网关发送给客户端的topic跟Request时完全一致，因此客户端可通过唯一topic
		确定返回数据是哪个请求的。

	消息体（body）

	mqant默认封装规则
		mqant默认封装为json结构体，具体结构体如下
		{
			Error string
			Result interface{}
		}

	Result
		代表hander函数执行正确时返回的结果

	类型
		1.bool
		2.int32
		3.int64
		4.long64
		5.float32
		6.float64
		7.[]byte
		8.string
		9.map[string]interface{}

	Error
		代表hander函数执行错误时返回的错误描述

	类型
		1.string
		2.error

	自定义返回格式
		mqant默认规则使用json来封装，但实际情况下不同的应用场景可能需要的封装数据格式有所不同，
		例如有些场景倾向于用protobuf封装。且默认的封装规则不支持自定义数据类型。
			app.SetProtocolMarshal(func(Result interface{}, Error string)(module.ProtocolMarsha1, string){
				// 下面可以实现你自己的封装规则（数据结构）
				r := &resultInfo{
					Error : Error,
					Result : Result,
				}
				b, err := json.Marshal(r)
				if err == nil {
					// 解析得到[]byte后用NewProtocolMarshal封装为module.ProtocolMarsha1
					return app.NewProtocolMarsha(b),
				} else {
					return nil, err.Error()
				}
			})

	ProtocolMarsha1
		如下mqant默认返回值是这样的 result map[string][string], err string
			func (m *Login) login(session)

*/