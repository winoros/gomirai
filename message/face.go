/*
 * Author: Logiase
 * Time: 2020/4/16 10:16
 */

package message

import "encoding/json"

type Face struct {
	Type   string
	FaceID int64
	Name   string
}

func (m *Face) GetType() string {
	return m.Type
}

func (m *Face) toJsonString() string {
	b, _ := json.Marshal(m)
	return string(b)
}
