package models

import "github.com/astaxie/beego/orm"

// init database model
func init() {
	orm.RegisterModel(new(Hook), new(User))
}
