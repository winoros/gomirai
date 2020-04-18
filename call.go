package gomirai

import "github.com/Logiase/gomirai/message"

// CommonCall 通用请求
type CommonCall struct {
	SessionKey string `json:"sessionKey"`
}

// MessageCall 消息用
type MessageCall struct {
	CommonCall
	Target       int64             `json:"target"`
	MessageChain []message.Message `json:"messageChain"`
}
