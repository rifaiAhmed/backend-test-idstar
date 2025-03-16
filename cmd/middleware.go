package cmd

import (
	"backend-test/helpers"
	"log"
	"net/http"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func (d *Dependency) MiddlewareValidateAuth(ctx *gin.Context) {
	auth := ctx.Request.Header.Get("Authorization")
	if auth == "" {
		log.Println("authorization empty")
		helpers.SendResponseHTTP(ctx, http.StatusUnauthorized, "unauthorized", nil)
		ctx.Abort()
		return
	}

	_, err := d.UserRepository.GetUserSessionByToken(ctx.Request.Context(), auth)
	if err != nil {
		log.Println("failed to get user session on DB: ", err)
		helpers.SendResponseHTTP(ctx, http.StatusUnauthorized, "unauthorized", nil)
		ctx.Abort()
		return
	}

	claim, err := helpers.ValidateToken(ctx.Request.Context(), auth)
	if err != nil {
		log.Println(err)
		helpers.SendResponseHTTP(ctx, http.StatusUnauthorized, "unauthorized", nil)
		ctx.Abort()
		return
	}

	if time.Now().Unix() > claim.ExpiresAt.Unix() {
		log.Println("jwt token is expired: ", claim.ExpiresAt)
		helpers.SendResponseHTTP(ctx, http.StatusUnauthorized, "unauthorized", nil)
		ctx.Abort()
		return
	}

	ctx.Set("token", claim)

	ctx.Next()
}

func MiddlewareCORS() gin.HandlerFunc {
	log.Println("Middleware CORS Aktif") // Logging untuk memastikan middleware berjalan
	return cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost:3039"},
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Authorization", "Content-Type"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	})
}
