package lib

import "testing"

import "time"

func TestGetUserAction(t *testing.T) {
	loc, _ := time.LoadLocation("Local")
	uid := int64(011)
	action := "look"
	t1, _ := time.ParseInLocation("2006-01-02 15:04:05", "2020-12-29 17:30:12", loc)
	t2, _ := time.ParseInLocation("2006-01-02 15:04:05", "2020-12-29 17:33:05", loc)
	t3, _ := time.ParseInLocation("2006-01-02 15:04:05", "2020-12-29 17:31:17", loc) // 出现两次，下标 5
	t4, _ := time.ParseInLocation("2006-01-02 15:04:05", "2020-12-29 17:48:52", loc)
	t5, _ := time.ParseInLocation("2006-01-02 15:04:05", "2020-12-29 17:19:23", loc)
	t6, _ := time.ParseInLocation("2006-01-02 15:04:05", "2020-12-29 17:46:40", loc)
	t7, _ := time.ParseInLocation("2006-01-02 15:04:05", "2020-12-29 17:32:31", loc)
	t8, _ := time.ParseInLocation("2006-01-02 15:04:05", "2020-12-29 17:31:17", loc)
	t9, _ := time.ParseInLocation("2006-01-02 15:04:05", "2020-12-29 17:33:19", loc)
	t10, _ := time.ParseInLocation("2006-01-02 15:04:05", "2020-12-29 17:18:27", loc)

	nodeList := getUserAction(uid, action)
	nodeList.add(t1)
	nodeList.add(t2)
	nodeList.add(t3)
	nodeList.add(t4)
	nodeList.add(t5)
	nodeList.add(t6)
	nodeList.add(t7)
	nodeList.add(t8)
	nodeList.add(t9)
	nodeList.add(t10)

	a := nodeList.TopNode

	bigValue := int64(0)
	sizeArray := make([]uint8, 0)
	for {
		if a == nil {
			break
		}

		if bigValue != 0 && bigValue < a.Value {
			t.Fatal("节点前面的值必须大于后面的值")
		}

		bigValue = a.Value

		sizeArray = append(sizeArray, a.Size)
		a = a.LeftTree
	}
	if sizeArray[5] != 2 {
		t.Fatal("序列出错")
	}
	t.Log("finish")

	v := nodeList.getCountInTime(t3.Unix())
	if v != 7 {
		t.Fatal("统计数量出错")
	}
}
