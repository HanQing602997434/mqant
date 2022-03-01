
// 初始化设置说明
/*
	概述
		room.QTable有许多设置初始项，可以用来控制以下的几大功能
			1.房间生命周期管理
				1）超时时间
				2）房间回收回调
				3）每帧更新函数Update

			2.运行周期（帧）
				多少间隔处理一次消息，一般根据业务场景而定

			3.消息队列容量管理
				1）消息接收队列容量管理
				2）消息发送队列容量管理

			4.异常处理函数
				1）handler未找到
				2）handler返回错误信息
				3）handler执行崩溃

	设置函数
		QTable.OnInit(subtable SubTable, opts ...room.Option) error

	房间生命周期管理
		
		超时时间
			当桌子超过指定的时间未与客户端（指所有客户端）建立通信（上行或下行通信），则认为桌子已超时，
			会自动回收
				1.如timeout = 0则自动检查功能关闭（需要你手动控制桌子的生命周期）
				2.默认timeout = 60
					room.TimeOut(60)

		房间回收回调
			当房间被回收是我们希望能够知道，方便我们做一些资源回收工作，例如房间信息存档，将房间room移除
			等等
				room.DestroyCallbacks(func(table room.BaseTable) error {
            		log.Info("回收了房间: %v", table.TableId())
            		_ = self.room.DestroyTable(table.TableId())
            		return nil
				}),

		Update
			桌子每周期（帧）都会执行，如果想每周期都做一些工作，可以设置update函数，如果不设置则只执行handler
				room.Update(this.Update)
				每帧都会调用
				func (this *MyTable) Update(ds time.Duration) {
    				//ds 上一帧跟当前帧的间隔时间
				}

		设置运行周期（帧）时间
			帧时间
				room.QTable使用mqant内置的时间轮，每一个周期执行完成后就会注册下一个执行周期的时间
				1.如果执行时间出现崩溃则不会再注册桌子，桌子会被强行Finish()回收。
				2.默认值为100ms
					room.RunInterval(50*time.Millisecond)

		消息队列容量管理
			如果每一个周期超过设置大小的消息数将插入失败

		设置接收消息的队列容量
			1.容量是设置大小的双倍，因为会创建两个消息队列
			2.默认值为256
				room.Capacity(256)

	异常处理函数
		
		当桌子handler未找到时
			基于它你可以实现一些特殊的业务逻辑
			room.NoFound(func(msg *room.QueueMsg) (value reflect.Value, e error) {
        		//return reflect.ValueOf(this.doSay), nil
        		return reflect.Zero(reflect.ValueOf("").Type()), errors.New("no found handler")
			})

		handler返回错误消息
			room.SetErrorHandle(func(msg *room.QueueMsg, err error) {
        		log.Error("Error %v Error: %v", msg.Func, err.Error())
			})

			如下情况就会调用ErrorHandler
			func (this *MyTable) doJoin(session gate.Session, msg map[string]interface{}) (err error) {
    			return errors.New("逻辑异常了")
			}

		handler执行崩溃
			当handler出现未处理的异常时，可以通过RecoverHandle监控到
				room.SetRecoverHandle(func(msg *room.QueueMsg, err error) {
        			log.Error("Recover %v Error: %v", msg.Func, err.Error())
    			})
*/