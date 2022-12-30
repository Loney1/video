package models

import "github.com/astaxie/beego/orm"

type Advert struct {
	Id       int
	Title    string
	SubTitle string
	AddTime  int64
	Img      string
	Url      string
}

func init() {
	orm.RegisterModel(new(Advert))
}
