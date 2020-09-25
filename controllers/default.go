package controllers

import (
	"Beegoproject1/db_mysql"
	"Beegoproject1/models"
	"encoding/json"
	"fmt"
	"github.com/astaxie/beego"
	"io/ioutil"
)

type MainController struct {
	beego.Controller
}

/*func (c *MainController) Get() {
	//1:获取name，age,sex
	user:=c.Ctx.Input.Query("user")
	pwd:=c.Ctx.Input.Query("pwd")
	//2.做对比
	if user != "ganping" || pwd !="616161" {
		c.Ctx.ResponseWriter.Write([]byte("对不起，数据错误"))
		return
	}
	c.Ctx.ResponseWriter.Write([]byte("数据请求成功"))*/


	/*c.Data["Website"] = "www.baidu.com"
	c.Data["Email"] = "ganping6104@gmail.com"
	c.TplName = "index.tpl"8
}*/
//写post请求
func (c *MainController) Post(){
	//接收post请求的参数
	/*name:=c.Ctx.Request.FormValue("name")
	age:=c.Ctx.Request.FormValue("age")
	sex:=c.Ctx.Request.FormValue("sex")
	fmt.Println(name,age,sex)
	//进行数据对比
	if name !="ganping" && age !="18"{
		c.Ctx.ResponseWriter.Write([]byte("数据校验失败"))
		return
	}
	c.Ctx.ResponseWriter.Write([]byte("数据校验成功"))//get请求

	var person models.Person//post请求
	data,err:=ioutil.ReadAll(c.Ctx.Request.Body)
	if err !=nil{
		c.Ctx.WriteString("数据接收错误，重试")
		return
	}
	err =json.Unmarshal(data,&person)
	if err !=nil{
		c.Ctx.WriteString("数据解析错误，重试")
		return
	}
	fmt.Println("姓名:",person.Name)
	fmt.Println("年龄:",person.Age)
	fmt.Println("性别:",person.Sex)
	c.Ctx.WriteString("数据请求成功")*/
	//定义请求接收name，birthday，address，nick
   /* name:=c.Ctx.Request.FormValue("name")
	birthday:=c.Ctx.Request.FormValue("birthday")
	address:=c.Ctx.Request.FormValue("address")
	nick:=c.Ctx.Request.FormValue("nick")
	fmt.Println(name,birthday,address,nick)
	//用得到的数据进行对比
	if name != "ganping" && address !="jinxian"{
		c.Ctx.WriteString("数据接收错误，重试")
	   return
	}
		c.Ctx.WriteString("数据接收成功，恭喜!!")

	var person models.Person
	data,err:=ioutil.ReadAll(c.Ctx.Request.Body)
	if err !=nil{
		c.Ctx.WriteString("数据接收错误")
		return
	}
	err =json.Unmarshal(data,&person)
	if err !=nil{
		c.Ctx.WriteString("数据解析错误")
		return
	}
	fmt.Println("姓名:",person.Name)
	fmt.Println("生日:",person.Birthday)
	fmt.Println("地址:",person.Address)
	fmt.Println("昵称",person.Nick)
	c.Ctx.WriteString("数据请求成功")*/

	data,err:=ioutil.ReadAll(c.Ctx.Request.Body)
	if err !=nil{
		c.Ctx.WriteString("数据接收错误")
		return
	}
	var user models.User
	err =json.Unmarshal(data,&user)
	if err !=nil{
		fmt.Println(err.Error())
		c.Ctx.WriteString("数据解析错误")
		return
	}
	//解析用户的数据，并保存数据
	id, err := db_mysql.InsertUser(user)
	if err !=nil{
		fmt.Println(err.Error())
		c.Ctx.WriteString("用户保存失败")
		return
	}
	fmt.Println(id)

	result := models.ResponseResult{
		Code:    0,
		Message: "保存成功",
		Data:    nil,
	}
	c.Data["json"] = &result
	c.ServeJSON()


}
