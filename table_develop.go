
// 桌子开发
/*
	概述
		具体的桌子开发
	
	第一步：桌子结构体
		桌子的定义
		mqant将桌子的核心功能封装到了QTable中，开发者主要继承QTable完成业务功能的开发
		type MyTable struct {
			room.QTable
			module module.RPCModule
			players map[string]room.BasePlayer
		}

	第二部：必须要继承的函数
		继承函数
		QTable有两个结构体需要开发者提供，以实现其内部功能
			1.GetModule()
				返回module.RPCModule

			2.GetSeats()
				返回一个桌子内座位(用户)的map
				GetSeats()  map[string]room.BasePlayer

		可选继承函数
		当桌子初始化后会调用OnCreate()，可以在这做一些业务初始化(加载数据)等操作，但一定要调用
		QTable.OnCreate()
			func (this *MyTable) OnCreate() {
				// 可以加载数据
				log.Info("MyTable OnCreate")
				// 一定要调用QTable.OnCreate()
				this.QTable.OnCreate()
			}

	第三步：初始化QTable
		初始化参数
		QTable有许多设置参数都可以在初始化时设置
			OnInit(subtable SubTable, opts ...Option) error

			this := &MyTable{
        		module:  module,
        		players: map[string]room.BasePlayer{},
    		}
			this.OnInit(this, opts...)

	第四步：实现handler
		桌子handler
		房间模块跟mqant的模块工作方式类似，也支持路由+handler的方式处理消息

		handler实现
			加入房间的handler
				1.首先将session跟room.BasePlayer绑定
				2.将room.BasePlayer注册桌子座位管理map中
				3.最后广播了一条消息给所有已加入房间的客户端
					func (this *MyTable) doJoin(session gate.Session, msg map[string]interface{}) (err error) {
    					player := &room.BasePlayerImp{}
    					player.Bind(session)
    					player.OnRequest(session)
    					this.players[session.GetSessionId()] = player
    					_ = this.NotifyCallBackMsg("/room/join", []byte(fmt.Sprintf("welcome to %v", msg["name"])))
    					return nil
					}
		
		handler注册
			路由是/room/join 处理函数为doJoin
				func NewTable(module module.RPCModule, opts ...room.Option) *MyTable {
    				this := &MyTable{
        				module:  module,
        				players: map[string]room.BasePlayer{},
    				}
    				this.OnInit(this, opts...)
    				this.Register("/room/say", this.doSay)
    				this.Register("/room/join", this.doJoin)
    				return this
				}
		
		跟客户端约定的数据结构
			{
				"table_id":"{table_id}",
				"action":"/room/say",
				"name":"{name}"
			}

		
*/