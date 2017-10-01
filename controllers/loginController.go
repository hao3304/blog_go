package controllers

import (
	"github.com/hao3304/blog/models"
	"encoding/json"
	"github.com/hao3304/blog/util"
	"github.com/astaxie/beego/orm"
)

type LoginController struct {
	BaseController
}

type Auth struct {
	Token string
	Expires int64
}

func (this *LoginController) Post() {
	o := orm.NewOrm()
	inputs := new(models.User)

	json.Unmarshal(this.Ctx.Input.RequestBody,&inputs)
	user := models.User{
		UserName:inputs.UserName,
	}

	err := o.Read(&user,"UserName")

	if err==nil {
		if user.Password == inputs.Password {
			t, expires := util.GenToken(&user)
			token := Auth{
				Token:t,
				Expires:expires,
			}
			this.Success(token)
		}else{
			this.Fail("用户名或者密码错误!",10001)
		}

	}else {
		this.Fail("用户名或者密码错误!",10001)
	}
}