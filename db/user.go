package db

import (
	mydb "filestore/db/mysql"
	"fmt"
)

type User struct {
	Username     string
	Email        string
	Phone        string
	SignupAt     string
	LastActiveAt string
	Status       int
}

// 用户注册
func UserSignup(username string, passwd string) bool {
	stmt, err := mydb.DBConn().Prepare(
		"insert ignore into tbl_user (`user_name`,`user_pwd`) values (?,?)")
	if err != nil {
		fmt.Println("Failed to insert, err:" + err.Error())
		return false
	}
	defer stmt.Close()

	ret, err := stmt.Exec(username, passwd)
	if err != nil {
		fmt.Println("Failed to insert,err", err.Error())
		return false
	}
	if rowsAffected, err := ret.RowsAffected(); err == nil && rowsAffected > 0 {
		return true
	}
	return false
}

func UserSignIn(username string, encpwd string) bool {
	stmt, err := mydb.DBConn.Prepare("select * from tbl_user where user_name=? limit 1")
	if err != nil {
		fmt.Println(err.Error())
		return false
	}

	rows, err := stmt.Query(username)
	if err != nil {
		fmt.Println(err.Error)
		return false
	}

	pRows := mydb.ParseRows(rows)
	if len(pRows) > 0 && string(pRows[0]["user_pwd"].([]byte)) == encpwd {
		return true
	}
	return false
}

func UpdateToken(username string,token string)  {
	stmt,err:=mydb.DBConn().Prepare(
		"replace into tbl_user_token (username,user_token) value (?,?)"
	)
	if err!=nil{
		fmt.Println(err.Error())
		return false 
	}
	defer stmt.Close()

	_,err:=stmt.Exec(username,token)
	if err!=nil{
		fmt.Println(err.Error())
		return false 
	}
	return true
}