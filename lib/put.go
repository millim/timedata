package lib

import "time"

//Put
func Put(t time.Time, userID int64, action string) error{
	nodeList := GetUserAction(userID, action)
	nodeList.Add(t)
	return nil
}
