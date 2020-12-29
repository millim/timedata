package service

import (
	//"github.com/millim/timedata/lib/pool"
	"github.com/millim/timedata/lib"
	"time"
)

//AddCount add user's action
func AddCount(userID int64, action string) error {

	now := time.Now()

	//请求效率优先，将数据存入队列便可以返回请求，但是无法及时错误信息
	//p := getPool()
	//p.AddWork(func(){lib.Put(now, userID, action)})

	return lib.Put(now, userID, action)
}
