package main

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"go-learn/webook/internal/repository"
	"go-learn/webook/internal/repository/dao"
	"go-learn/webook/internal/service"
	"go-learn/webook/internal/web"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"time"
)

func main() {
	db := initDB()
	server := initWebServer()
	initUserhdl(db, server)
	server.Run(":8080")
}

func initUserhdl(db *gorm.DB, server *gin.Engine) {
	ud := dao.NewUserDAO(db)
	ur := repository.NewUserRepository(ud)
	us := service.NewUserService(ur)
	hdl := web.NewUserHandler(us)
	hdl.RegisterRoutes(server)
}

func initDB() *gorm.DB {
	db, err := gorm.Open(mysql.Open("root:root@tcp(localhost:13316)/webook"))
	if err != nil {
		panic(err)
	}
	err = dao.InitTables(db)
	if err != nil {
		panic(err)
	}
	return db
}

func initWebServer() *gin.Engine {
	server := gin.Default()
	server.Use(cors.New(cors.Config{
		//	//AllowOrigins:     []string{"http://localhost:3000"},
		//	//AllowHeaders:     []string{"Content-Type"},
		AllowCredentials: true,
		AllowOriginFunc: func(origin string) bool {
			return true
		},
		MaxAge: 12 * time.Hour,
	}))
	return server
}
