package middleware

import "github.com/gin-gonic/gin"

func AccessControlAllow(c *gin.Context) {
	header := c.Writer.Header()
	if val := header["Access-Control-Allow-Origin"]; len(val) == 0 {
		header["Access-Control-Allow-Origin"] = []string{"*"}
	}
	if val := header["Access-Control-Allow-Methods"]; len(val) == 0 {
		header["Access-Control-Allow-Methods"] = []string{"*"}
	}
	if val := header["Access-Control-Allow-Headers"]; len(val) == 0 {
		header["Access-Control-Allow-Headers"] = []string{"*"}
	}
}
