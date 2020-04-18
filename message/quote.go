/*
 * Author: Logiase
 * Time: 2020/4/16 10:12
 */

package message

import "encoding/json"

type Quote struct {
	Type     string
	Id       int64
	GroupID  int64
	SenderID int64
	TargetID int64
	Origin   []Message
}

func (m *Quote) GetType() string {
	return m.Type
}

func (m *Quote) toJsonString() string {
	b, _ := json.Marshal(m)
	return string(b)
}
