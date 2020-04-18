/*
 * Author: Logiase
 * Time: 2020/4/16 9:14
 */

package message

type TypeMessage interface {
	GetType() string
	toJsonString() string
}
