package web

import (
	regexp "github.com/dlclark/regexp2"
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"go-learn/webook/internal/domain"
	"go-learn/webook/internal/service"
	"net/http"
)

// 正则表达式，邮箱，密码,必须包含字母，数字，特殊字符并不少于八位
const (
	emailRegexPattern    = "^\\w+([-+.]\\w+)*@\\w+([-.]\\w+)*\\.\\w+([-.]\\w+)*$"
	passwordRegexPattern = `^(?=.*[A-Za-z])(?=.*\d)(?=.*[$@$!%*#?&])[A-Za-z\d$@$!%*#?&]{8,}$`
)

// 账号登录信息
type UserHandler struct {
	emailRexExp    *regexp.Regexp
	passwordRexExp *regexp.Regexp
	svc            *service.UserService
}

// 预加载正则表达式
func NewUserHandler(svc *service.UserService) *UserHandler {
	return &UserHandler{
		emailRexExp:    regexp.MustCompile(emailRegexPattern, regexp.None),
		passwordRexExp: regexp.MustCompile(passwordRegexPattern, regexp.None),
		svc:            svc,
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
// ctx是一个 gin.Context类型的参数，它代表了当前请求上下文,包含了请求和响应的信息，以及提供一些方法，用于获取请求和响应的信息、设置响应头、设置响应状态码等操作
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
	//校验两次密码是否一致
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
	//将post请求参数email,password传给domain.User对应
	err = h.svc.SignUp(ctx, domain.User{
		Email:    req.Email,
		Password: req.Password,
	})
	switch err {
	case nil:
		ctx.String(http.StatusOK, "注册成功")
	case service.ErrDuplicateEmail:
		ctx.String(http.StatusOK, "邮箱冲突")
	default:
		ctx.String(http.StatusOK, "系统错误")
	}
	//判定邮箱冲突，需要拿到数据库的唯一索引冲突错误。这里用MYSQL GO驱动的error定义，找到错误

	//if err != nil {
	//	ctx.String(http.StatusOK, "系统错误")
	//	return
	//}
	//ctx.String(http.StatusOK, "hello,你在注册") //第一个参数为响应状态码，第二个参数为响应主体body
}

// 登录
func (h *UserHandler) Login(ctx *gin.Context) {
	type Req struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}
	var req Req
	if err := ctx.Bind(&req); err != nil {
		return
	}
	u, err := h.svc.Login(ctx, req.Email, req.Password)
	switch err {
	case nil:
		//登录后设置Sessions
		sess := sessions.Default(ctx)
		sess.Set("UserId", u.Id)
		sess.Options(sessions.Options{
			MaxAge: 900, //设置900s过期，也就是15分钟
		})
		err = sess.Save()
		if err != nil {
			ctx.String(http.StatusOK, "系统错误")
			return
		}
		ctx.String(http.StatusOK, "登陆成功")
	case service.ErrInvalidUserOrPassword:
		ctx.String(http.StatusOK, "用户名或密码不对")
	default:
		ctx.String(http.StatusOK, "系统错误")
	}
}

// 修改
func (h *UserHandler) Edit(ctx *gin.Context) {

}

// 查看
func (h *UserHandler) Profile(ctx *gin.Context) {

}
