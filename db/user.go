package db 

import (
	mydb "filestore/db/mysql"
	"fmt"
)

type User struct {
	Username string 
	Email string 
	Phone string 
	SignupAt string 
	LastActiveAt string 
	Status int
}

// 用户注册
func UserSignup(username string,passwd string)bool{
	stmt,err:=mydb.DBConn().Prepare(
		"insert ignore into tbl_user (`user_name`,`user_pwd`) values (?,?)")
	if err!=nil{
		fmt.Println("Failed to insert, err:" + err.Error())
		return false
	}
	defer stmt.Close()

	ret,err:=stmt.Exec(username,passwd)
	if err!=nil{
		fmt.Println("Failed to insert,err",err.Error())
		return false
	}
	if rowsAffected,err:=ret.RowsAffected();err==nil&&rowsAffected>0{
		return true 
	}
	return false
}