package lib

import "time"

//Get 获取数量
func Get(d time.Duration, userID int64, action string) (int64, error) {
	n := getUserAction(userID, action)
	endTime := time.Now().Unix() - int64(d)
	//cache 可以处理片段缓存
	return n.getCountInTime(endTime), nil
}
