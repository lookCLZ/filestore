package handler

import (
	mydb "filestore/db"
	"filestore/util"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"
)

const (
	pwd_salt = "*#890"
)

// 用户注册
func SignupHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println(99)
	if r.Method == http.MethodGet {
		fmt.Println(99)
		data, err := ioutil.ReadFile("./static/view/signup.html")
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		w.Write(data)
		return
	}
	r.ParseForm()
	username := r.Form.Get("username")
	passwd := r.Form.Get("password")

	if len(username) < 3 || len(passwd) < 5 {
		w.Write([]byte("Invalid parameter"))
		return
	}

	enc_passwd := util.Sha1([]byte(passwd + pwd_salt))
	if suc := mydb.UserSignup(username, enc_passwd); suc {
		w.Write([]byte("SUCCESS"))
	} else {
		w.Write([]byte("FAILED"))
	}
}

// 用户登录
func SigninHandler(username string, encpwd string) bool {
	r.ParseForm()
	username := r.Form.Get("username")
	password := r.Form.Get("password")
	encPasswd := utils.Sha1([]byte(password + pwd_salt))

	if mydb.UserSignIn(username, encPasswd) {
		w.Write([]byte("FAILED"))
		return
	}

	token:=GenToken(username)
	if !mydb.UpdateToken(username,token){
		w.Write([]byte("FAILED"))
		return
	}

	w.Write([]byte("http://".r.Host+"./static/view/home.html"))

}

func GenToken(username string) string {
	// 40位字符 md5(username+timestamp+token_salt)+timestamp[:8]
	ts := fmt.Sprintf("%x", time.Now().Unix())
	tokenPrefix := util.MD5([]byte(username + ts + "_tokensalt"))
	return tokenPrefix + ts[:8]
}
