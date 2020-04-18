/*
 * Author: Logiase
 * Time: 2020/4/16 10:19
 */

package message

import "encoding/json"

type Xml struct {
	Type string
	XML  string
}

func (m *Xml) GetType() string {
	return m.Type
}

func (m *Xml) toJsonString() string {
	b, _ := json.Marshal(m)
	return string(b)
}

type Json struct {
	Type string
	Json string
}

func (m *Json) GetType() string {
	return m.Type
}

func (m *Json) toJsonString() string {
	b, _ := json.Marshal(m)
	return string(b)
}

type App struct {
	Type    string
	Content string
}

func (m *App) GetType() string {
	return m.Type
}

func (m *App) toJsonString() string {
	b, _ := json.Marshal(m)
	return string(b)
}

type Poke struct {
	Type string
	Name string
}

func (m *Poke) GetType() string {
	return m.Type
}

func (m *Poke) toJsonString() string {
	b, _ := json.Marshal(m)
	return string(b)
}
