package admin

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
	"beego-blog/models"
)


func init() {
	// set default database
	orm.RegisterDataBase("default", "mysql", "root:@tcp(127.0.0.1:3306)/beego-blog?charset=utf8", 30)

	orm.RegisterModel(new(models.User))

	orm.RunSyncdb("default", false, true)

}

type BaseController struct {
	beego.Controller
}

func (c *BaseController) Islogin(){
	admin_login := c.GetSession("admin_login")
	if admin_login == nil{
		c.Redirect("/admin/login", 302)
		return
	}
	return
}
