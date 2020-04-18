package handler

import (
	mydb "filestore/db"
	"filestore/util"
	"fmt"
	"io"
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
func SigninHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		// 返回 html页面
		data, err := ioutil.ReadFile("./static/view/signin.html")
		if err != nil {
			io.WriteString(w, "internet server error")
			return
		}
		w.Write(data)
	} else if r.Method == "POST" {
		r.ParseForm()
		username := r.Form.Get("username")
		password := r.Form.Get("password")
		encPasswd := util.Sha1([]byte(password + pwd_salt))

		if !mydb.UserSignIn(username, encPasswd) {
			fmt.Println("aaa")
			w.Write([]byte("FAILED"))
			return
		}
		fmt.Println("ccc")
		token := GenToken(username)
		fmt.Println("dddd")
		if !mydb.UpdateToken(username, token) {
			w.Write([]byte("FAILED"))
			return
		}
		fmt.Println("EEE")
		w.Write([]byte("http://" + r.Host + "/static/view/home.html"))
	}
}

func GenToken(username string) string {
	// 40位字符 md5(username+timestamp+token_salt)+timestamp[:8]
	ts := fmt.Sprintf("%x", time.Now().Unix())
	tokenPrefix := util.MD5([]byte(username + ts + "_tokensalt"))
	return tokenPrefix + ts[:8]
}
