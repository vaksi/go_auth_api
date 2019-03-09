/*
=====================Golang======================
	Author 	: Audy Vaksi Pranata
	Link	: https://github.com/vaksi
	Email	: vaksipranata@gmail.com
==================================================
*/
package main

import (
	"go_auth_api/controllers"
	"os"
	"time"

	jwt "gopkg.in/appleboy/gin-jwt.v2"

	"github.com/fvbock/endless"
	"github.com/gin-gonic/gin"
)

func main() {

	port := os.Getenv("PORT")
	if port == "" {
		port = "3100"
	}
	router := gin.New()
	router.Use(gin.Logger())
	router.Use(gin.Recovery())

	// the jwt middleware
	authMiddleware := &jwt.GinJWTMiddleware{
		Realm:      "test zone",
		Key:        []byte("secret key"),
		Timeout:    time.Hour,
		MaxRefresh: time.Hour * 24,
		Authenticator: func(userId string, password string, c *gin.Context) (string, bool) {
			if controllers.UserAuth(userId, password) {
				return userId, true
			}

			return userId, false
		},

		Authorizator: func(userId string, c *gin.Context) bool {
			if userId == "admin" {
				return true
			}

			return false
		},

		Unauthorized: func(c *gin.Context, code int, message string) {
			c.JSON(code, gin.H{
				"code":    code,
				"message": message,
			})
		},
		// TokenLookup is a string in the form of "<source>:<name>" that is used
		// to extract token from the request.
		// Optional. Default value "header:Authorization".
		// Possible values:
		// - "header:<name>"
		// - "query:<name>"
		// - "cookie:<name>"
		TokenLookup: "header:Authorization",
		// TokenLookup: "query:token",
		// TokenLookup: "cookie:token",
	}

	/*
		======================Declare your route================
	*/
	router.POST("/login", authMiddleware.LoginHandler)
	router.POST("/register", controllers.CreateUser)
	auth := router.Group("/auth")
	auth.Use(authMiddleware.MiddlewareFunc())
	{
		auth.GET("/users", controllers.GetUsers)
		auth.GET("/refresh_token", authMiddleware.RefreshHandler)
		auth.GET("/check", controllers.CheckHandler)
		auth.GET("/check_role", controllers.CheckRole)
	}
	auth.GET("/hello", controllers.HelloHandler)

	endless.ListenAndServe(":"+port, router)

}
