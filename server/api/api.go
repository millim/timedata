package api

import (
	"github.com/gin-gonic/gin"
	"github.com/millim/timedata/server/service"
	"net/http"
	"time"
)

//GetAddCountBody body
type GetAddCountBody struct {
	UID    int64  `form:"user_id"`
	Action string `form:"action"`
}

//GetAddCount 为了方便使用get方法，标准应该使用post
func GetAddCount(c *gin.Context) {
	var body GetAddCountBody
	c.Bind(&body)

	err := service.AddCount(body.UID, body.Action)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"err": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{})
}

//GetCountBody body
type GetCountBody struct {
	UID    int64  `form:"user_id"`
	Action string `form:"action"`
	Minute int64  `form:"minute"`
}

//GetActionCount get
func GetActionCount(c *gin.Context) {
	var body GetCountBody
	c.Bind(&body)
	count, err := service.GetCount(body.UID, body.Action, time.Duration(body.Minute))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"err": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"count": count})
}
