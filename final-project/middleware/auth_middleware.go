package middleware

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/cperdiansyah/golang-h8-dts/final-project/helper"
	"github.com/cperdiansyah/golang-h8-dts/final-project/model/http/response"
)

func AuthMiddleware() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		verifyToken, err := helper.VerifyToken(ctx)

		if err != nil {
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, response.ErrorResponse{
				Code:   http.StatusUnauthorized,
				Status: "Unauthorized",
				Errors: err.Error(),
			})

			return
		}

		ctx.Set("userData", verifyToken)
		ctx.Next()
	}
}
