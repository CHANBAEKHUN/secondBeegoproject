package db_mysql

import (
	"Beegoproject1/models"
	"crypto/md5"
	"database/sql"
	"encoding/hex"
	"fmt"
	"github.com/astaxie/beego"
	_ "github.com/go-sql-driver/mysql"
)
var Db*sql.DB
func init(){
	fmt.Println("连接数据库")
	config :=beego.AppConfig
	dbDriver:=config.String("db_driverName")
	dbUser:=config.String("db_user")
	dbPassword:=config.String("db_password")
	dbIp:=config.String("db_ip")
	dbName:=config.String("db_name")

	connUrl:=dbUser+":"+dbPassword+"@tcp("+dbIp+")/"+dbName+"?charset=utf8"
	db,err:=sql.Open(dbDriver,connUrl)
	if err !=nil{
		panic("错误")
	}
	Db=db
}
func InsertUser(user models.User)(int64,error){
	hashMd5:=md5.New()
	hashMd5.Write([]byte(user.Password))
	bytes:=hashMd5.Sum(nil)
	user.Password =hex.EncodeToString(bytes)
	fmt.Println("将要保存的用户名:",user.Nick,"密码:",user.Password,"生日:",user.Birthday,"地址:",user.Address,"姓名:",user.My)
	result,err:=Db.Exec("insert into user(nick,password,birthday,address,my ) values(?,?,?,?,?)",user.Nick,user.Password,user.Birthday,user.Address,user.My)
	if err !=nil{
		return -1, err
	}
	id,err :=result.RowsAffected()
		if err !=nil{
			return -1,err
	}
return id,nil
}
//查询用户
//func QueryUser(){
	//Db.QueryRow("select * from")
//}