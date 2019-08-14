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

type TerminalController struct {
	beego.Controller
}

func (self *TerminalController) Get() {
	self.Data["context"] = self.GetString("context")
	self.Data["namespace"] = self.GetString("namespace")
	self.Data["pod"] = self.GetString("pod")
	self.Data["container"] = self.GetString("container")
	self.TplName = "terminal.html"
}
