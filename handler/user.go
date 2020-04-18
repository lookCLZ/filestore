package handler

import (
	mydb "filestore/db"
	"filestore/util"
	"fmt"
	"io/ioutil"
	"net/http"
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
