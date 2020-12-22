/******
** @创建时间 : 2020/12/22 16:18
** @作者 : SongZhiBin
******/
package biz

import (
	"context"
)

// IClient: 客户端接口
type IClient interface {
	// SendMes 发送消息 非阻塞 失败返回
	SendMes(msg IMessage) bool

	// ReturnMsgChan 返回一个可收取的channel 用于接收消息
	MsgChan() <-chan IMessage

	// GetCtx: 返回context.Context
	GetCtx() context.Context

	// Shutdown 回收资源
	// 多次调用无副作用!
	Shutdown()
}
