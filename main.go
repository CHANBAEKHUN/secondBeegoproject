package main

import (
	_ "Beegoproject1/routers"
	_"Beegoproject1/db_mysql"
	"github.com/astaxie/beego"
)

func main() {

	//appName:=config.String("appname")
	//fmt.Println("名称",appName)
	//port,err:=config.Int("httpport")
	//if err !=nil{
	//	panic("项目配置解析失败")
	//}

	beego.Run()
}

