/*
 * Author: Logiase
 * Time: 2020/4/16 10:17
 */

package message

import "encoding/json"

type Image struct {
	Type    string
	ImageID string
	URL     string
	Path    string
}

func (m *Image) GetType() string {
	return m.Type
}

func (m *Image) toJsonString() string {
	b, _ := json.Marshal(m)
	return string(b)
}

type FlashImage struct {
	Type    string
	ImageID string
	URL     string
	Path    string
}

func (m *FlashImage) GetType() string {
	return m.Type
}

func (m *FlashImage) toJsonString() string {
	b, _ := json.Marshal(m)
	return string(b)
}
