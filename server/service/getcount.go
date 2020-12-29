package service

import (
	"errors"
	"time"

	"github.com/millim/timedata/lib"
)

const (
	//MaxTime 24 * 60 min
	MaxTime = 1440
	//MinTime 1 min
	MinTime = 1
)

//GetCount get user action count
func GetCount(userID int64, action string, duration time.Duration) (int64, error) {
	if duration > MaxTime || duration < MinTime {
		return 0, errors.New("时间长度在1 ~ 1440分")
	}
	return lib.Get(duration, userID, action)
}
