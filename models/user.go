package models

import (
	"github.com/astaxie/beego/orm"
)

type User struct {
	Id int `orm:"pk;auto"`
	UserName string
	Password string
	Roles []*Role `orm:"rel(m2m)"`
	Base
}

func init()  {
	orm.RegisterModel(new(User))
}

