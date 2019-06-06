package config

import "github.com/kataras/iris"

const (
	// key定义
	CODE string = "code"
	MSG  string = "msg"
	DATA string = "data"

	// msg 消息定义
	Success         string = "操作成功！"
	Failure         string = "操作有误！"
	UserNameExisted string = "用户名已存在！"
	RegisteredSuc   string = "注册成功！"
	RegisteredErr   string = "注册失败！"
	UserLoginErr    string = "用户名或密码错误！"
	LoginSuc        string = "登陆成功！"
	LoginErr        string = "登陆失败！"
	CaptChaErr      string = "验证码有误！"
	CreatedSuc      string = "创建成功！"
	CreatedErr      string = "创建失败！"
	UpdatedSuc      string = "编辑成功！"
	UpdatedErr      string = "编辑失败！"
	DeletedSuc      string = "删除成功！"
	DeletedErr      string = "删除失败！"
)

// 200 define
func Ok_(ctx iris.Context, msg string) {
	Ok(ctx, msg, nil)
}

func Ok(ctx iris.Context, msg string, data interface{}) {
	ctx.StatusCode(iris.StatusOK)
	_, _ = ctx.JSON(iris.Map{
		CODE: iris.StatusOK,
		MSG:  msg,
		DATA: data,
	})
}

// 401 error define
func Unauthorized(ctx iris.Context, msg string, data interface{}) {
	unauthorized := iris.StatusUnauthorized

	ctx.StatusCode(unauthorized)
	_, _ = ctx.JSON(iris.Map{
		CODE: unauthorized,
		MSG:  msg,
		DATA: data,
	})
}

// common error define
func Error(ctx iris.Context, status int, msg string, data interface{}) {
	ctx.StatusCode(status)
	_, _ = ctx.JSON(iris.Map{
		CODE: status,
		MSG:  msg,
		DATA: data,
	})
}
