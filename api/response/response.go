package response

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func SrsOkJson(c *gin.Context) {
	// 开始时间
	c.JSON(http.StatusOK, 0)
	c.Abort()
}

func SrsErrJson(c *gin.Context) {
	// 开始时间
	c.JSON(http.StatusOK, 1)
	c.Abort()
}

func SrsErrForwardJson(c *gin.Context) {
	// 开始时间
	c.JSON(http.StatusOK, map[string]interface{}{
		"code": 0,
		"data": map[string]interface{}{
			"urls": []string{},
		},
	})
	c.Abort()
}

func SrsOkForwardJson(c *gin.Context, urls []string) {
	// 开始时间
	c.JSON(http.StatusOK, map[string]interface{}{
		"code": 0,
		"data": map[string]interface{}{
			"urls": urls,
		},
	})
	c.Abort()
}
