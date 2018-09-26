package controllers

import (
	"github.com/astaxie/beego"
)

// MainController definition.
type MainController struct {
	beego.Controller
}

// Get method.
func (c *MainController) Get() {
	c.Data["Website"] = "igetspot"
	c.Data["Email"] = "aditya.bayu@mncgroup.com"
	c.TplName = "index.tpl"
	c.Render()
}
