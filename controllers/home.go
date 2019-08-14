/*

 ____             __               __
/\  _`\          /\ \             /\ \
\ \,\L\_\     __ \ \ \___      __ \ \ \____     __
 \/_\__ \   /'__`\\ \  _ `\  /'__`\\ \ '__`\  /'__`\
   /\ \L\ \/\ \L\.\\ \ \ \ \/\ \L\.\\ \ \L\ \/\ \L\.\_
   \ `\____\ \__/.\_\ \_\ \_\ \__/.\_\ \_,__/\ \__/.\_\
    \/_____/\/__/\/_/\/_/\/_/\/__/\/_/\/___/  \/__/\/_/

*/

package controllers

import (
	"github.com/astaxie/beego"
)

type HomeController struct {
	beego.Controller
}

func (c *HomeController) Get() {
	c.TplName = "index.html"
}
