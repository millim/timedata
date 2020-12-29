package lib

import (
	"fmt"
	"sync"
)

var users sync.Map

func GetUserAction(userID int64, action string)*NodeList{
		c, ok := users.Load(userID)
		if !ok {
			var s sync.Map
			c = &s
			users.Store(userID, s)
		}
		fmt.Println(c)
		list, ok := c.(*sync.Map).Load(action)
		if !ok {
			list = &NodeList{}
			c.(*sync.Map).Store(action, list)
		}
		return list.(*NodeList)
}
