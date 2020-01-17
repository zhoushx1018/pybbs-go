package main

import (
  _ "github.com/zhoushx1018/pybbs-go/routers"
  "github.com/astaxie/beego"
  "github.com/astaxie/beego/orm"
  "github.com/zhoushx1018/pybbs-go/models"
  _ "github.com/go-sql-driver/mysql"
  _ "github.com/zhoushx1018/pybbs-go/utils"
  _ "github.com/zhoushx1018/pybbs-go/templates"
)

func init() {
  url := beego.AppConfig.String("jdbc.url")
  port := beego.AppConfig.String("jdbc.port")
  username := beego.AppConfig.String("jdbc.username")
  password := beego.AppConfig.String("jdbc.password")
  orm.RegisterDataBase("default", "mysql", username+":"+password+"@tcp("+url+":"+port+")/pybbs-go?charset=utf8&parseTime=true&charset=utf8&loc=Asia%2FShanghai", 30)
  orm.RegisterModel(
    new(models.User),
    new(models.Topic),
    new(models.Section),
    new(models.Reply),
    new(models.ReplyUpLog),
    new(models.Role),
    new(models.Permission))
  orm.RunSyncdb("default", false, true)
}

func main() {
  //orm.Debug = true
  //ok, err := regexp.MatchString("/topic/edit/[0-9]+", "/topic/edit/123")
  //beego.Debug(ok, err)
  beego.Run()
}
