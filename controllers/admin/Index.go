package admin

type IndexController struct {
	BaseController
}

func (c *IndexController) Index() {
	c.Islogin()
	c.Redirect("/admin/dashboard", 302)
}


func (c *IndexController) Login(){
	c.SetSession("admin_login", int(1))
	c.Redirect("/admin/dashboard", 302)
}