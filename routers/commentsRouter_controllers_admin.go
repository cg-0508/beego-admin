package routers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context/param"
)

func init() {

    beego.GlobalControllerRouter["beego-blog/controllers/admin:DashboardController"] = append(beego.GlobalControllerRouter["beego-blog/controllers/admin:DashboardController"],
        beego.ControllerComments{
            Method: "index",
            Router: `/`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

}
