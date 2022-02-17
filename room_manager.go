
// 房间管理
/*
	概述
		房间管理比较简单，通常我们希望一个进程中可以创建多个房间，这样才能最大化利用
		服务器资源，因此我们将房间模块划分为room、table两个级别，room用来管理table

	创建房间的结构体
		func (self *tabletest) OnInit(app module.App, settings *conf.ModuleSettings) {
			self.BaseModule.OnInit(self, app, settings,
				server.RegisterInterval(15*time.Second),
				server.RegisterTTL(30*time.Second),
			)
			self.room = room.NewRoom(self)
		}

	创建桌子
		CreateById(module module.RPCModule, tableId string, newTablefunc NewTableFunc) (BaseTable, error) 

		//调用代码
		table, err = self.room.CreateById(self, table_id, self.NewTable)
		
		CreateById
		table_id 桌子唯一ID，作为这个房间内桌子的唯一标识
		NewTableFunc 创建桌子的具体方法，room负责创建桌子的具体逻辑，它仅仅维护桌子在room下的对应

	NewTableFunc
		type NewTableFunc func(module module.RPCModule, tableId string) (BaseTable, error)

		由开发者自己实现桌子的具体创建逻辑，如下
			func (self *tabletest) NewTable(module module.RPCModule, tableId string) (room.BaseTable, error) {
    			table := NewTable(
        			module,
        			room.TableId(tableId),
        			room.DestroyCallbacks(func(table room.BaseTable) error {
            			log.Info("回收了房间: %v", table.TableId())
            			_ = self.room.DestroyTable(table.TableId())
            			return nil
       				}),
    			)
    			return table, nil
			}

	获取桌子
		GetTable(tableId string) BaseTable

		// 代码
		table := self.room.GetTable(table_id)

		如果没有创建过table_id的桌子将返回nil

	示例代码
		简单聊天室
		一个简单的聊天室功能，table_id由客户端指定，当桌子不存在则创建一个新的，然后将客户端消息写入桌子的消息队列中
		func (self *tabletest) gatesay(session gate.Session, msg map[string]interface{}) (r string, err error) {
    		table_id := msg["table_id"].(string)
    		action := msg["action"].(string)
    		table := self.room.GetTable(table_id)
    		if table == nil {
        		table, err = self.room.CreateById(self, table_id, self.NewTable)
        		if err != nil {
            		return "", err
        		}
    		}
    		erro := table.PutQueue(action, session, msg)
    		if erro != nil {
        		return "", erro
    		}
    		return "success", nil
		}

	handler注册
		func (self *tabletest) OnInit(app module.App, settings *conf.ModuleSettings) {
    		self.BaseModule.OnInit(self, app, settings,
        		server.RegisterInterval(15*time.Second),
        		server.RegisterTTL(30*time.Second),
    			)
    		self.room = room.NewRoom(self)
    		self.GetServer().RegisterGO("HD_room_say", self.gatesay)
		}

	跟客户端约定数据结构
		{
			"table_id":"{table_id}",
			"action":"/room/say",
			"name":"{name}"
		}
*/