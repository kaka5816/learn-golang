package main

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"
	"go-learn/webook/internal/repository"
	"go-learn/webook/internal/repository/dao"
	"go-learn/webook/internal/service"
	"go-learn/webook/internal/web"
	"go-learn/webook/internal/web/middleware"
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
	login := &middleware.LoginMiddlewareBuilder{}
	//存储数据的，也就是userid
	//通过设置的 secret 字符串，来计算 hash 值并放在 cookie 中
	store := cookie.NewStore([]byte("secret"))
	//cookie的名字叫做ssid, login.CheckLogin()登陆校验
	server.Use(sessions.Sessions("ssid", store), login.CheckLogin())
	return server
}
