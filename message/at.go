/*
 * Author: Logiase
 * Time: 2020/4/16 10:15
 */

package message

import "encoding/json"

type At struct {
	Type    string
	Target  int64
	Display string
}

func (m *At) GetType() string {
	return m.Type
}

func (m *At) toJsonString() string {
	b, _ := json.Marshal(m)
	return string(b)
}

type AtAll struct {
	Type string
}

func (m *AtAll) GetType() string {
	return m.Type
}

func (m *AtAll) toJsonString() string {
	b, _ := json.Marshal(m)
	return string(b)
}
