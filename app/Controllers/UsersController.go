package Controllers

import (
	"github.com/dgrijalva/jwt-go"
	"github.com/kataras/iris"
	"iris/app/Models"
	"iris/config"
)

// 获取token中的用户ID
func GetThisTokensUserId(ctx iris.Context) int64 {
	userMsg := ctx.Values().Get("jwt").(*jwt.Token).Claims.(jwt.MapClaims)

	return int64(userMsg["Id"].(float64))
}

// 获取当前用户
func Me(ctx iris.Context) {
	var user = new(Models.User)

	user.Id = int64(GetThisTokensUserId(ctx))

	// 查找用户
	has, _ := Models.GetUserByModelsUser(user)

	// 是否存在该用户
	if !has {
		config.NotFound(ctx, config.TheUserNotFound)
		return
	}

	config.Ok(ctx, config.Success, user)
}

// 修改昵称
func UpdateUsersNameById(ctx iris.Context) {
	var user = new(Models.User)

	user.Id = int64(GetThisTokensUserId(ctx))

	user.Name = ctx.PostValue("name")

	effect, err := Models.UpdateUserById(user)

	if effect <= 0 || err != nil {
		ctx.Application().Logger().Errorf("用户[%s]修改失败。%s", user.Name, err.Error())
		config.Unauthorized(ctx, config.UpdatedErr, nil)
		return
	}

	config.Ok(ctx, config.UpdatedSuc, user)
}

// 修改密码
func UpdateUserPasswordById(ctx iris.Context) {
	var (
		user  = new(Models.User)
		mUser = new(Models.User)
	)

	user.Id = int64(GetThisTokensUserId(ctx))

	user.Password = ctx.PostValue("oldPassword")

	mUser.Id = user.Id
	// 查找用户
	has, err := Models.GetUserByModelsUser(mUser)

	if err != nil {
		ctx.Application().Logger().Errorf("用户[%s]不存在。%s", user.Id, err.Error())
		return
	}

	if !has {
		config.Unauthorized(ctx, config.TheUserNotFound, nil)
		return
	}

	// 校验密码
	checkPassword := config.CheckPWD(user.Password, mUser.Password)

	if !checkPassword {
		config.Unauthorized(ctx, config.UserLoginErr, nil)
		return
	}

	mUser.Password = config.AESEncrypt([]byte(ctx.PostValue("password")))

	effect, err := Models.UpdateUserById(mUser)

	if effect <= 0 || err != nil {
		ctx.Application().Logger().Errorf("用户[%s]修改失败。%s", mUser.Name, err.Error())
		config.Unauthorized(ctx, config.UpdatedErr, nil)
		return
	}

	config.Ok(ctx, config.UpdatedSuc, mUser)
}
