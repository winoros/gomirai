/*
 * Author: Logiase
 * Time: 2020/4/16 10:08
 */

package message

import "encoding/json"

type Source struct {
	Type string
	Id   int64
	Time int64
}

func (m *Source) GetType() string {
	return m.Type
}

func (m *Source) toJsonString() string {
	b, _ := json.Marshal(m)
	return string(b)
}
