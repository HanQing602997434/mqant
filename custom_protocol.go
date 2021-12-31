
// 自定义通信协议
/*
	概述
		如果项目无法使用mqtt协议，需要自定义通信协议，可以使用下面的方法，如果
		替换默认的agent，以上长连接网关文章的功能都将不可用，可以自己实现

	agent定义
		type Agent interface {
			OnInit(gate Gate, conn network.Conn) error
			WriteMsg(topic string, body []byte) error
			Close()
			Run() (err error)
			OnClose() error
			Destory()
			ConnTime() time.Time
			RevNum() int64
			SendNum() int64
			IsClosed() bool
			GetSession() Session
		}

	实现自定义的agent
		CustomAgent.go
		package mgate

		import (
			"bufio"
			"github.com/liangdas/mqant/gate"
			"github.con/liangdas/mqant/log"
			"github.com/liangdas/mqant/module"
			"github.com/liangdas/mqant/network"
			"time"
		)

		func NewCustomAgent(module module.RPCModule) *CustomAgent {
			a := &CustomAgent{
				module: module,
			}
			return a
		}

		type CustomAgent struct {
			gate.Agent
			module								module.RPCModule
			session								gate.Session
			conn								network.Conn
			r									*bufio.Reader
			w									*bufio.Writer
			gate								gate.Gate
			rev_num								int64
			send_num							int64
			last_storage_heartbeat_data_time 	time.Duration // 上一次发送存储心跳时间
			isclose								bool
		}

		func (this *CustomAgent) OnInit(gate gate.Gate, conn network.Conn) error {
			log.Info("CustomAgent", "OnInit")
			this.conn = conn
			this.gate = gate
			this.r = bufio.NewReader(conn)
			this.w = bufio.NewWriter(conn)
			this.isclose = false
			this.rev_num = 0
			this.send_num = 0
			return nil
		}

		// 给客户端发送消息
		func (this *CustomAgent) WriteMsg(topic string, body []byte) error {
			this.send_num++
			// 粘包完成后调下面的语句发送数据
			// this.w.Write()
			return nil 
		}

		func (this *CustomAgent) Run() (err error) {
			log.Info("CustomAgent", "开始读数据了")

			this.session, err = this.gate.NewSessionByMap(map[string]interface{} {
				"Sessionid" : "生成一个随机数",
				"Network" : this.conn.RemoteAddr().Network(),
				"IP" : this.conn.RemoteAddr().String(),
				"Serverid" : this.module.GetServerId(),
				"Settings" : make(map[string]string),
			})

			// 这里可以循环读取客户端的数据

			// 这个函数返回后连接就会被关闭
			return nil
		}

		// 接收到一个数据包
		func (this *CustomAgent) OnRecover(topic string, msg []byte) {
			// 通过解析的数据得到
			moduleType := ""
			_func := ""

			// 如果要对这个请求进行分布式跟踪调式，就执行下面这行语句
			// a.session.CreateRootSpan("gate")

			// 然后请求后端模块，第一个参数为session
			result, e := this.module.RpcInvoke(moduleType, _func, this.session, msg)
			log.Info("result", result)
			log.Info("error", e)

			// 回复客户端
			this.WriteMsg(topic, []byte("请求成功了谢谢"))

			this.heartbeat()
		}

		func (this *CustomAgent) heartbeat() {
			// 自定义网关需要你自己设计心跳协议
			if this.GetSession().GetUserId() != "" {
				// 这个连接已经绑定userid
				interval := int64(this.last_storage_heartbeat_data_time) + int64(this.gate.Options().Heartbeat) // 单位纳秒
				if interval < time.Now().UnixNano() {
					// 如果用户信息存储心跳包的时长已经大于一秒
					if this.gate.GetStorageHandler() != nil {
						this.gate.GetStorageHandler().Heartbeat(this.GetSession())
						this.last_storage_heartbeat_data_time = time.Duration(time.Now().UnixNano())
					}
				}
			}
		}

		func (this *CustomAgent) Close() {
			log.Info("CustomAgent", "主动断开连接")
			this.conn.Close()
		}

		func (this *CustomAgent) Close() error {
			this.isclose = true
			log.Info("CustomAgent", "连接断开事件")
			// 这个一定要调用，不然gate可能注销不了，造成内存溢出
			this.gate.GetAgentLearner().DisConnect(this) // 发送连接断开的事件
			return nil
		}

		func (this *CustomAgent) Destory() {
			this.conn.Destory()
		}

		func (this *CustomAgent) RevNum() int64 {
			return this.rev_num
		}

		func (this *CustomAgent) SendNum() int64 {
			return this.send_num
		}

		func (this *CustomAgent) IsClosed() bool {
			return this.isclose
		}

		func (this *CustomAgent) GetSession().gate.Session {
			return this.session
		}

		替换默认agent生成器
		func (this *Gate) OnInit(app module.App, settings *conf.ModuleSettings) {
			// 注意这里一定要用gate.Gate 而不是module.BaseModule
			this.Gate.OnInit(this, app, settings, gate.Heartbeat(time.Second*10))
			this.Gate.SetCreateAgent(func() gate.Agent {
				agent := NewCustomAgent(this)
				return agent
			})
		}
*/
