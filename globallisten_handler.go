
// 全局监听handler
/*
	概述
		我们通常希望能监控handler的具体执行情况，例如做监控报警等等

	应用级别handler监控
		app := mqant.CreateApp(
			module.SetClientRPChandler(func(app module.App, server registry.Node, rpcinfo rpcpb.RPCInfo, result interface{}, err string, exec_time int64) {			
			}),
			module.SetServerRPChandler(func(app module.App, server module.Module, callInfo mqrpc.CallInfo) {			
			}),
		)

	调用方监控
		module.SetClientRPChandler(func(app module.App, server registry.Node, rpcinfo rpcpb.RPCInfo, result interface{}, err string, exec_time int64) {
		})

	服务方监控
		module.SetServerRPChandler(func(app module.App, server module.Module, callInfo mqrpc.CallInfo) {
		})
*/