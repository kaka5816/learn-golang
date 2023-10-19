package web

import (
	regexp "github.com/dlclark/regexp2"
	"github.com/gin-gonic/gin"
	"net/http"
)

// 正则表达式，邮箱，密码,必须包含字母，数字，特殊字符并不少于八位
const (
	emailRegexPattern    = "^\\w+([-+.]\\w+)*@\\w+([-.]\\w+)*\\.\\w+([-.]\\w+)*$"
	passwordRegexPattern = `^(?=.*[A-Za-z])(?=.*\d)(?=.*[$@$!%*#?&])[A-Za-z\d$@$!%*#?&]{8,}$`
)

type UserHandler struct {
	emailRexExp    *regexp.Regexp
	passwordRexExp *regexp.Regexp
}

// 预加载正则表达式
func NewUserHandler() *UserHandler {
	return &UserHandler{
		emailRexExp:    regexp.MustCompile(emailRegexPattern, regexp.None),
		passwordRexExp: regexp.MustCompile(passwordRegexPattern, regexp.None),
	}
}

// 注册路由,由gin.Engine来实现
func (h *UserHandler) RegisterRoutes(server *gin.Engine) {
	server.POST("/users/signup", h.SignUp)  //注册  h.SignUp表示调用UserHandler结构体实例h上的SignUp方法
	server.POST("/users/login", h.Login)    //登录
	server.POST("/users/edit", h.Edit)      //编辑
	server.GET("/users/profile", h.Profile) //查看信息
	//分组路由，等价
	//ug := server.Group("/users")
	//ug.POST("/signup", h.SignUp)
	//ug.POST("/login", h.Login)
	//ug.POST("/edit", h.Edit)
	//ug.GET("/profile", h.Profile)
}

// 注册邮箱密码
// ctx是一个 gin.Context类型的参数，它代表了当前请求上下文,包含了请求的信息以及用于处理响应的方法
func (h *UserHandler) SignUp(ctx *gin.Context) {
	type SignUpReq struct {
		Email           string `json:"email"`
		Password        string `json:"password"`
		ConfirmPassword string `json:"confirmPassword"`
	}
	//ctx.Bind(&req)是将请求的数据解析并填充到req结构体中
	var req SignUpReq
	if err := ctx.Bind(&req); err != nil {
		return
	}
	isEmail, err := h.emailRexExp.MatchString(req.Email)
	if err != nil {
		ctx.String(http.StatusOK, "系统错误")
		return
	}
	if !isEmail {
		ctx.String(http.StatusOK, "非法邮箱格式")
		return
	}
	//
	if req.Password != req.ConfirmPassword {
		ctx.String(http.StatusOK, "两次密码不一致")
		return
	}
	//
	isPassword, err := h.passwordRexExp.MatchString(req.Password)
	if err != nil {
		ctx.String(http.StatusOK, "系统错误")
		return
	}
	if !isPassword {
		ctx.String(http.StatusOK, "密码格式错误，必须包含字母，数字，特殊字符并不少于八位")
		return
	}
	ctx.String(http.StatusOK, "hello,你在注册") //第一个参数为响应状态码，第二个参数为响应主体body
}

// 登录
func (h *UserHandler) Login(ctx *gin.Context) {

}

// 修改
func (h *UserHandler) Edit(ctx *gin.Context) {

}

// 查看
func (h *UserHandler) Profile(ctx *gin.Context) {

}
