package gomirai

import (
	"github.com/Logiase/gomirai/api"
	"time"
)

// NewPlainTextMessage 构造一条PlainText消息
func (b *Bot) NewPlainTextMessage(text string) (m api.MessageCall) {
	m = api.MessageCall{
		MessageChain: []api.Message{
			{
				Type: "Plain",
				Text: text,
			},
		},
	}
	m.SessionKey = b.session
	return
}

// NewImageMessageFromURL 构造一条图片消息
func (b *Bot) NewImageMessageFromURL(url string) (m api.MessageCall) {
	m = api.MessageCall{
		Urls: []string{
			url,
		},
	}
	m.SessionKey = b.session
	return
}

// NewRecallMessageEventCall 构造一个撤回消息事件
func (b *Bot) NewRecallMessageEventCall(messageID int64) (m api.MessageCall) {
	m = api.MessageCall{}
	m.SessionKey = b.session
	m.Target = messageID
	return m
}

// NewMuteGroupMemberEventCall 构造一个禁言或解禁群员事件
func (b *Bot) NewMuteGroupMemberEventCall(groupID, memberID int64, time time.Duration) (m api.ManageCall) {
	m = api.ManageCall{
		MemberID: memberID,
	}
	m.Target = groupID
	m.SessionKey = b.session
	if int(time.Seconds()) != 0 {
		m.Time = int(time.Seconds())
	}
	return
}

// NewKickGroupMemberEventCall 构造一个踢出群员事件
func (b *Bot) NewKickGroupMemberEventCall(groupID, memberID int64, msg string) (m api.ManageCall) {
	m = api.ManageCall{
		MemberID: memberID,
		Msg:      msg,
	}
	m.SessionKey = b.session
	m.Target = groupID
	return
}
