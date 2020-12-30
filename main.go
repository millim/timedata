package main

import "github.com/millim/timedata/server"

//ADD_COUNT 为了方便使用get方法，标准应该使用post
//localhost:8080/api/add_count?user_id=&action=
//
//GET_COUNT
//localhost:8080/api/get_count?user_id=&action=&minute=
func main() {
	s := server.Get()
	s.Run(":8080")
}
