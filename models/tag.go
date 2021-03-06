package models

import (
	"github.com/astaxie/beego/orm"
	"time"
)

type Tag struct {
	Id      int64     `json:"id"`
	Name    string    `json:"name"`
	Status  int       `json:"status"`
	Insuid  int64     `json:"insuid"`
	Instime time.Time `orm:"column(instime);auto_now_add;type(datetime)" json:"instime"`
}

func NewTag() *Tag {
	return &Tag{}
}

func (this *Tag) Query() orm.QuerySeter {
	return orm.NewOrm().QueryTable(new(Tag)).Filter("status", 1)
}

func (this *Tag) Add() (int64, error) {
	this.Status = 1
	return orm.NewOrm().Insert(this)
}

// 获取全部标签
func (this *Tag) GetAll() ([]*Tag, error) {
	tag := []*Tag{}
	_, err := this.Query().All(&tag)
	return tag, err
}
