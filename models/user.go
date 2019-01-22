package models


type User struct {
	Id   int
	Username string `orm:"size(100)"`
	Password string `orm:"size(200)"`
	Img string `orm:"size(200)"`
	IsAdmin int `orm:"size(1)"`
	CreateTime int64 `orm:"size(20)"`
}

