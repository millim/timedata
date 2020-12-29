package service

import (
	"errors"
	"time"
)

const (
	//MaxTime 24 * 60 min
	MaxTime = 1440
	MinTime = 1
)

//GetCount get user action count
func GetCount(userID int64, ation string, duration time.Duration)(int64,error){
	if duration > MaxTime || duration < MinTime{
		return 0, errors.New("时间长度在1 ~ 1440分")
	}
	return 0, nil
}
