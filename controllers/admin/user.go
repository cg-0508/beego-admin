package admin

import (
	"beego-blog/models"
	"fmt"
	"github.com/astaxie/beego/orm"
	"time"
)

type UserController struct {
	BaseController
}

func (c *UserController) UserList () {
	c.Islogin()
	o := orm.NewOrm()
	var maps []orm.Params
	num, _ := o.Raw("SELECT * FROM user").Values(&maps)
	fmt.Println(num)
	//page := util.PageUtil(1, 1, 2, maps)
	//fmt.Println(page)
	c.Data["userList"] = maps
	//c.Data["Page"] = page
	c.Layout = "admin/layout.tpl"
	c.LayoutSections  = make(map[string]string)
	c.LayoutSections["LayoutSidebar"] = "admin/sidebar.tpl"
	c.LayoutSections["LayoutHeader"] = "admin/header.tpl"
	c.TplName = "admin/user/list.tpl"
}

func (c *UserController) UserAdd() {
	c.Islogin()
	c.Layout = "admin/layout.tpl"
	c.LayoutSections  = make(map[string]string)
	c.LayoutSections["LayoutSidebar"] = "admin/sidebar.tpl"
	c.LayoutSections["LayoutHeader"] = "admin/header.tpl"
	c.TplName = "admin/user/add.tpl"
}

func (c *UserController) Add() {
	username := c.GetString("username")
	password := c.GetString("password")
	user_id, _ := c.GetInt("user_id")
	fmt.Println(username, password, user_id)
	if username == "" || password == "" {
		return
	}

	f, _, err := c.GetFile("image")

	if err != nil {
		fmt.Println(err)
		return
	}
	defer f.Close()
	userImg := "/static/upload/" + username + ".jpg"
	err = c.SaveToFile("image", userImg)
	if err != nil {
		fmt.Println(err)
	}
	o := orm.NewOrm()
	var user  = models.User{}
	if user_id != 0 {
		user.Id = user_id
		if o.Read(&user) == nil {
			user.Username = username
			user.Password = password
			user.IsAdmin = 0
			user.Img = userImg
			if num, err := o.Update(&user); err == nil {
				fmt.Println(num)
			}
		}
	}else{
		t := time.Now()
		fmt.Println(t)
		user.Username = username
		user.Password = password
		user.CreateTime = t.Unix()
		user.IsAdmin = 0
		user.Img = userImg
		id, err := o.Insert(&user)
		if err != nil {
			return
		}
		fmt.Println(id)
	}
	c.Redirect("/admin/user", 302)
}

func (c *UserController) EditUser() {
	userId, _ := c.GetInt(":id")
	fmt.Println(userId)
	o := orm.NewOrm()
	user := models.User{Id: userId}
	err := o.Read(&user)
	if err == orm.ErrNoRows {
		fmt.Println("查询不到")
	}else if err == orm.ErrMissPK {
		fmt.Println("找不到主键")
	}else {
		fmt.Println(user)
		c.Data["userInfo"] = user
		c.Layout = "admin/layout.tpl"
		c.LayoutSections  = make(map[string]string)
		c.LayoutSections["LayoutSidebar"] = "admin/sidebar.tpl"
		c.LayoutSections["LayoutHeader"] = "admin/header.tpl"
		c.TplName = "admin/user/edit.tpl"
	}
}
