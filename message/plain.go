/*
 * Author: Logiase
 * Time: 2020/4/16 10:16
 */

package message

import "encoding/json"

type Plain struct {
	Type string
	Text string
}

func (m *Plain) GetType() string {
	return m.Type
}

func (m *Plain) toJsonString() string {
	b, _ := json.Marshal(m)
	return string(b)
}
