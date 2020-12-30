package lib

import (
	"sync"
)

var usersMap sync.Map

//GetUserAction 获取用户指定用户的数据集
func getUserAction(userID int64, action string) *NodeList {
	c, ok := usersMap.Load(userID)
	if !ok {
		var s sync.Map
		c = &s
		usersMap.Store(userID, c)
	}
	list, ok := c.(*sync.Map).Load(action)
	if !ok {
		list = &NodeList{}
		c.(*sync.Map).Store(action, list)
	}
	return list.(*NodeList)
}
