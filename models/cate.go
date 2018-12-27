package models

import (
	"github.com/astaxie/beego/orm"
	"time"
)

type Cate struct {
	Id      int64
	Pid     int
	Name    string
	Sort    int
	Status  int
	Insuid  int64
	Instime time.Time `orm:"column(instime);auto_now_add;type(datetime)"`
}

func (this *Cate) Query() orm.QuerySeter {
	return orm.NewOrm().QueryTable(new(Cate)).Filter("status", 1)
}

func NewCate() *Cate {
	return &Cate{}
}

func (this *Cate) Add() (int64, error) {
	this.Status = 1
	return orm.NewOrm().Insert(this)
}

// 获取全部分类
func (this *Cate) GetAll() ([]*Cate, error) {
	cate := []*Cate{}
	_, err := this.Query().OrderBy("-sort").All(&cate)
	return cate, err
}
