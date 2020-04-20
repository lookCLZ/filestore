package db

import (
	"database/sql"
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

type TableUser struct {
	UserName string
	UserPwd  sql.NullString
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
	fmt.Println(0)
	stmt, err := mydb.DBConn().Prepare("select user_name,user_pwd from tbl_user where user_name=? limit 1")
	fmt.Println(0.5)
	if err != nil {
		fmt.Println(1)
		fmt.Println(err.Error())
		return false
	}
	defer stmt.Close()

	tuser := TableUser{}
	err = stmt.QueryRow(username).Scan(
		&tuser.UserName,
		&tuser.UserPwd,
	)

	fmt.Printf("%+v", tuser)
	if tuser.UserPwd.String == encpwd {
		fmt.Println(true)
		return true
	}
	return false
}

func UpdateToken(username string, token string) bool {
	stmt, err := mydb.DBConn().Prepare(
		"replace into tbl_user_token (user_name,user_token) value (?,?)")
	if err != nil {
		fmt.Println(err.Error())
		return false
	}
	defer stmt.Close()

	_, err = stmt.Exec(username, token)
	if err != nil {
		fmt.Println(err.Error())
		return false
	}
	return true
}

func GetUserInfo(username string) (User, error) {
	user := User{}
	stmt, err := mydb.DBConn().Prepare(
		"select user_name,signup_at from tbl_user where user_name=? limit 1")
	if err != nil {
		fmt.Println(err.Error())
		return user, err
	}
	defer stmt.Close()

	err = stmt.QueryRow(username).Scan(&user.Username, &user.SignupAt)
	if err != nil {
		return user, err
	}
	if err != nil {
		return user, err
	}
	return user, nil
}
