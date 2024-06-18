package common

import (
	"fmt"

	"github.com/calendarproject/model/resp"
	"github.com/gin-gonic/gin"

	"net/http"
)

// 跨域中间件
func CORSMiddleware() gin.HandlerFunc {

	return func(context *gin.Context) {

		context.Header("Access-Control-Allow-Origin", "*")

		context.Header("Access-Control-Allow-Headers", "Content-Type,AccessToken,X-CSRF-Token, Authorization, Token, x-token")

		context.Header("Access-Control-Allow-Methods", "POST, GET, OPTIONS, DELETE, PATCH, PUT")

		context.Header("Access-Control-Expose-Headers", "Content-Length, Access-Control-Allow-Origin, Access-Control-Allow-Headers, Content-Type")

		context.Header("Access-Control-Allow-Credentials", "true")

		if context.Request.Method == "OPTIONS" {

			context.AbortWithStatus(http.StatusNoContent)

		} else {

			context.Next()

		}

	}

}

// 异常中间件
func RecoveryMiddleware() gin.HandlerFunc {
	return func(context *gin.Context) {
		defer func() {
			if err := recover(); err != nil {
				resp.Fail(context, nil, fmt.Sprint(err))
			}
		}()

		context.Next()
	}

}
