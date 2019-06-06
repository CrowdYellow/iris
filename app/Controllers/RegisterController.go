package Controllers

import (
	"fmt"
	"github.com/kataras/iris"
	"iris/app/Models"
	"iris/config"
	"time"
)

func Register(ctx iris.Context) {
	var user = new(Models.User)
	//if err := ctx.ReadJSON(&user); err != nil {
	//	config.Error(ctx, iris.StatusInternalServerError, config.RegisteredErr, nil)
	//}

	user.Name      = ctx.PostValue("name")
	user.Phone     = ctx.PostValue("phone")
	user.Password  = config.AESEncrypt([]byte(ctx.PostValue("password")))
	user.RoleId    = 1
	user.Enable    = true
	user.CreatedAt = time.Now()
	user.UpdatedAt = time.Now()

	fmt.Println(ctx.PostValue("name"))
	fmt.Println(ctx.PostValue("phone"))

	effect, err := Models.CreateUser(user)

	if effect <= 0 || err != nil {
		ctx.Application().Logger().Errorf("用户[%s]注册失败。%s", user.Name, err.Error())
		return
	}

	config.Ok(ctx, config.RegisteredSuc, user)
}
