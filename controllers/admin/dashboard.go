package admin

type DashboardController struct {
	BaseController
}


func (c *DashboardController) Dashboard() {
	c.Layout  = "admin/layout.tpl"
	c.LayoutSections  = make(map[string]string)
	c.LayoutSections["LayoutSidebar"] = "admin/sidebar.tpl"
	c.LayoutSections["LayoutHeader"] = "admin/header.tpl"
	c.TplName = "admin/dashboard.tpl"
}
