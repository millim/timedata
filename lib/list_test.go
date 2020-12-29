package lib

import "testing"

func TestGetUserAction(t *testing.T) {
	b := GetUserAction(1, "test")
	t.Log("--->", b)
}
