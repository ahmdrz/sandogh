package fileserver

import (
	"net/http"
	"os"
	"path"

	"github.com/gin-gonic/gin"
)

func (s *Server) directoryCheck() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		p := path.Join(s.config.BaseDirectory, ctx.Param("directory"))
		if !directoryExists(p) {
			if ctx.Request.Method == "POST" {
				os.Mkdir(p, 0755)
			} else {
				ctx.String(http.StatusNotFound, "No such file or directory")
				ctx.Abort()
				return
			}
		}
		ctx.Next()
	}
}

func (s *Server) authMiddleware() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		method := ctx.Request.Method
		if method == "POST" || method == "DELETE" {
			token := ctx.GetHeader("Authentication")
			if token != s.config.SecretKey {
				ctx.AbortWithStatusJSON(http.StatusForbidden, gin.H{
					"status":  "forbidden",
					"message": "Forbidden access to this method",
					"method":  method,
				})
				return
			}
		}
		ctx.Next()
	}
}
