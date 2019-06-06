package Controllers

import (
	"github.com/kataras/iris"
	"iris/app/Models"
	"iris/config"
)

func Login(ctx iris.Context) {
	var (
		user = new(Models.User)
		mUser = new(Models.User)
		checkPassword bool
	)
	name := ctx.PostValue("name")
	password := ctx.PostValue("password")

	user.Name = name
	user.Password = password

	mUser.Name = user.Name

	// 查找用户
	has, err := Models.GetUserByName(mUser)

	if err != nil {
		ctx.Application().Logger().Errorf("用户[%s]登录失败。%s", user.Name, err.Error())
		return
	}

	if !has {
		config.Unauthorized(ctx, config.UserLoginErr, nil)
		return
	}

	// 校验密码
	checkPassword = config.CheckPWD(user.Password, mUser.Password)

	if !checkPassword {
		config.Unauthorized(ctx, config.UserLoginErr, nil)
		return
	}

	config.Ok(ctx, config.LoginSuc, mUser)
}
