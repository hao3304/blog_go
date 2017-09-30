package main

import (
	_ "github.com/hao3304/blog/routers"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
	_ "github.com/hao3304/blog/models"
	"github.com/hao3304/blog/models"
)

func init()  {
	orm.RegisterDriver("mysql", orm.DRMySQL)
	orm.RegisterDataBase("default","mysql","root:hao3304@/blog?charset=utf8")
	orm.RunCommand()


	o := orm.NewOrm()
	u := new(models.User)
	u.UserName = "jack"
	u.Password = "hao3304"
	o.Insert(u)
}

func main() {
	beego.Run()
}

