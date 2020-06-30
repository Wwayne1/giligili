package api

import (
	"giligili/serializer"

	"github.com/gin-gonic/gin"
)

//CreateVideo 视频投稿接口
func CreateVideo(c *gin.Context) {
	c.JSON(200, serializer.Response{
		Code: 0,
		Msg:  "投稿成功",
	})
}
