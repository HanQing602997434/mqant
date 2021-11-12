
// docker
/*
	docker安装（https://www.cnblogs.com/keyou1/p/11511067.html）

		什么是docker

			装应用的容器

			开发、测试、运维都偏爱的容器化技术

			轻量级

			扩展性

			一次构建、多次分享、随处运行

	概述
		1.镜像中包含了Golang编译环境和mqant所必须的中间件
			consul
			nats

		2.前提条件是已安装docker环境

	快速部署docker版本服务端
		1.下载镜像
			docker pull 1587790525/mqant-example:latest

		2.启动镜像
			docker run -p 3563:3563 -p 3653:3653 -p 8080:8080 1587790525/mqant-example

		3.访问服务接口
			http://127.0.0.1:8080/say?name=mqant

		4.验证golang客户端 运行 robot/robot_task.go

	使用dockerfile编译镜像
		1.克隆git工程
			git clone https://github.com/liangdas/mqant-example

		2.制作docker镜像
			首先进入工程目录下
			docker build -t <镜像名称>
*/