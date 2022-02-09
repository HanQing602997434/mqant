
// 长连接编写第一个网关
/*
	代码组织结构
		重新组织一下代码目录结构，新增了一个gate目录用来存放网关代码，robot目录用来存放
		访问网关的mqtt客户端代码
			工程目录
				|-bin
					|-conf
						|-server.conf
				|-helloworld
					|-module.go
					|-xxx.go
				|-gate
					|-module.go
				|-robot
					|-test
						|-manager.go
						|-work.go
					|-robot_task.go
				|-main
*/