package routers

import (
	"github.com/hao3304/blog/controllers"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/plugins/cors"
	"github.com/astaxie/beego/context"
	"strings"
	"github.com/hao3304/blog/util"
)

func init() {
	beego.InsertFilter("*", beego.BeforeRouter, cors.Allow(&cors.Options{
		AllowAllOrigins:  true,
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Authorization", "Access-Control-Allow-Origin", "Access-Control-Allow-Headers", "Content-Type"},
		ExposeHeaders:    []string{"Content-Length", "Access-Control-Allow-Origin", "Access-Control-Allow-Headers", "Content-Type"},
		AllowCredentials: true,
	}))

	beego.Router("/", &controllers.MainController{})
	beego.Router("/login", &controllers.LoginController{})

	ns := beego.NewNamespace("/admin",
		beego.NSBefore(func(ctx *context.Context) {

			authToken := ctx.Input.Header("Authorization")
			beego.Debug("auth token:", authToken)
			kv := strings.Split(authToken, " ")
			if len(kv) != 2 || kv[0] != "Bearer" {
				beego.Error("AuthString invalid:", authToken)
				ctx.Abort(503,"认证信息格式错误。")
			}
			token := kv[1]
			claim, err :=util.ParseJwt(token)
			if err == nil {
				ctx.Input.SetData("username",claim.Audience)
			}
			ctx.Abort(503,"认证失败，请重新登陆。")
		}),
		beego.NSRouter("/user",&controllers.UserController{}),
	)
	beego.AddNamespace(ns)
}
