package Controllers

import (
	"github.com/kataras/iris"
	"iris/app/Middleware"
	"iris/app/Models"
	"iris/config"
	"time"
)

func Register(ctx iris.Context) {
	var (
		err    error
		user   = new(Models.User)
		effect int64
	)
	err = ctx.ReadForm(user)
	if err != nil {
		ctx.Application().Logger().Errorf("用户[%s]注册失败。%s", user.Name, err.Error())
		config.Error(ctx, iris.StatusInternalServerError, config.RegisteredErr, nil)
		return
	}

	user.NickName  = user.Name
	user.Avatar    = "/images/user/default.jpg"

	user.Password  = config.AESEncrypt([]byte(user.Password))
	user.RoleId    = 1
	user.Enable    = true
	user.CreatedAt = time.Now()
	user.UpdatedAt = time.Now()

	effect, err = Models.CreateUser(user)

	if effect <= 0 || err != nil {
		ctx.Application().Logger().Errorf("用户[%s]注册失败。%s", user.Name, err.Error())
		config.Error(ctx, iris.StatusInternalServerError, config.RegisteredErr, nil)
		return
	}

	token, err := Middleware.GenerateToken(user)

	config.Ok(ctx, config.RegisteredSuc, ResponseUserWithUser(token, user))
}
