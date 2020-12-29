package lib

import (
	"testing"
	"time"
)

func TestPut(t *testing.T) {

	loc, _ := time.LoadLocation("Local")
	uid := int64(011)
	action := "look"
	t1, _ := time.ParseInLocation("2006-01-02 15:04:05", "2020-12-29 17:30:12", loc)
	t2, _ := time.ParseInLocation("2006-01-02 15:04:05", "2020-12-29 17:33:05", loc)
	t3, _ := time.ParseInLocation("2006-01-02 15:04:05", "2020-12-29 17:31:17", loc)
	t4, _ := time.ParseInLocation("2006-01-02 15:04:05", "2020-12-29 17:48:52", loc)
	t5, _ := time.ParseInLocation("2006-01-02 15:04:05", "2020-12-29 17:19:23", loc)
	t6, _ := time.ParseInLocation("2006-01-02 15:04:05", "2020-12-29 17:46:40", loc)
	t7, _ := time.ParseInLocation("2006-01-02 15:04:05", "2020-12-29 17:32:31", loc)
	t8, _ := time.ParseInLocation("2006-01-02 15:04:05", "2020-12-29 17:31:17", loc)
	t9, _ := time.ParseInLocation("2006-01-02 15:04:05", "2020-12-29 17:33:19", loc)
	t10, _ := time.ParseInLocation("2006-01-02 15:04:05", "2020-12-29 17:18:27", loc)

	Put(t1, uid, action)
	Put(t2, uid, action)
	Put(t3, uid, action)
	Put(t4, uid, action)
	Put(t5, uid, action)
	Put(t6, uid, action)
	Put(t7, uid, action)
	Put(t8, uid, action)
	Put(t9, uid, action)
	Put(t10, uid, action)

	v, ok :=users.Load(uid)
	if !ok {
		t.Error("失败")
	}

	t.Log(v.(*NodeList).TopNode.Value)

}
