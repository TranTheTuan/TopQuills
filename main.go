package main

import (
	"github.com/TranTheTuan/TopQuills/middlewares"
	_ "github.com/TranTheTuan/TopQuills/routers"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	"github.com/astaxie/beego/plugins/cors"
	_ "github.com/go-sql-driver/mysql"
)

func init() {
	orm.RegisterDriver("mysql", orm.DRMySQL)
	dataString := beego.AppConfig.String("mysqluser") + ":" + beego.AppConfig.String("mysqlpass") + "@" + beego.AppConfig.String("mysqlurl") + "/" + beego.AppConfig.String("mysqldb") + "?charset=utf8"
	orm.RegisterDataBase("default", "mysql", dataString)
}

func main() {
	if beego.BConfig.RunMode == "dev" {
		beego.BConfig.WebConfig.DirectoryIndex = true
		beego.BConfig.WebConfig.StaticDir["/swagger"] = "swagger"
	}
	beego.InsertFilter("*", beego.BeforeRouter, cors.Allow(&cors.Options{
		AllowOrigins:     []string{"*"},
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Authorization", "authorization", "Access-Control-Allow-Origin", "Access-Control-Allow-Headers", "Content-Type", "contenttype", "accesscontrolalloworigin", "Access-Control-Allow-Methods"},
		ExposeHeaders:    []string{"Content-Length", "Authorization", "authorization", "Access-Control-Allow-Origin", "Access-Control-Allow-Headers", "Access-Control-Request-Headers", "Content-Type", "accesscontrolalloworigin"},
		AllowCredentials: true,
	}))
	beego.InsertFilter("*", beego.BeforeRouter, middlewares.CheckAuthentication)
	beego.Run()
}
