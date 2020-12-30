package server

import (
	"github.com/gin-gonic/gin"
	"github.com/millim/timedata/server/api"
)

//Get 获取Server
func Get() *gin.Engine {
	g := gin.Default()

	//为了方便使用get方法，标准应该使用post
	g.GET("/api/add_count", api.GetAddCount)
	g.GET("/api/get_count", api.GetActionCount)
	return g
}
