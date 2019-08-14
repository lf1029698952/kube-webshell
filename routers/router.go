/*

 ____             __               __
/\  _`\          /\ \             /\ \
\ \,\L\_\     __ \ \ \___      __ \ \ \____     __
 \/_\__ \   /'__`\\ \  _ `\  /'__`\\ \ '__`\  /'__`\
   /\ \L\ \/\ \L\.\\ \ \ \ \/\ \L\.\\ \ \L\ \/\ \L\.\_
   \ `\____\ \__/.\_\ \_\ \_\ \__/.\_\ \_,__/\ \__/.\_\
    \/_____/\/__/\/_/\/_/\/_/\/__/\/_/\/___/  \/__/\/_/

*/

package routers

import (
	"github.com/astaxie/beego"
	"github.com/du2016/web-terminal-in-go/k8s-webshell/controllers"
)

func init() {
	beego.Router("/", &controllers.HomeController{})
	beego.Router("/terminal", &controllers.TerminalController{}, "get:Get")
	beego.Handler("/terminal/ws", &controllers.TerminalSockjs{}, true)
}
