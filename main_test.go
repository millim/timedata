package main

import "testing"

func TestRun(t *testing.T){
	a := "3"
	b := "2"
	t.Log(a > b)
	t.Log(a == b)
}
