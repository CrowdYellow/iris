package Controllers

import (
	"github.com/kataras/iris"
	"iris/app/Models"
	"iris/config"
)

func ShowUserById(ctx iris.Context) {
	var user = new(Models.User)
	id, _ := ctx.Params().GetUint("id")

	user.Id = int64(id)

	// 查找用户
	has, _ := Models.GetUserByName(user)

	// 是否存在该用户
	if !has {
		config.NotFound(ctx, config.TheUserNotFound)
		return
	}

	config.Ok(ctx, config.LoginSuc, user)
}

func UpdateUserById(ctx iris.Context) {
	//
}
