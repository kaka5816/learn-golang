package main

import (
	"github.com/gin-gonic/gin"
	"go-learn/webook/internal/web"
)

func main() {
	hdl := web.NewUserHandler()
	server := gin.Default()
	//server.Use(cors.New(cors.Config{
	//	//AllowOrigins:     []string{"http://localhost:3000"},
	//	//AllowHeaders:     []string{"Content-Type"},
	//	AllowCredentials: true,
	//	AllowOriginFunc: func(origin string) bool {
	//		return true
	//	},
	//	MaxAge: 12 * time.Hour,
	//}))
	hdl.RegisterRoutes(server)
	server.Run(":8080")

}
