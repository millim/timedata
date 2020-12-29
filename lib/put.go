package lib

import "time"

//Put 添加数据
func Put(t time.Time, userID int64, action string) error {
	nodeList := getUserAction(userID, action)
	nodeList.add(t)
	return nil
}
