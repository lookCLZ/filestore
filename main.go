package main

import (
	"filestore/handler"
	"fmt"
	"net/http"
	"os"
	"path/filepath"
)

func main() {
	pwd, _ := os.Getwd()
	fmt.Println("os.Getwd", pwd)
	http.Handle("/static/", http.FileServer(http.Dir(filepath.Join(pwd, "./"))))

	http.HandleFunc("/file/upload", handler.UploadHandler)
	http.HandleFunc("/file/upload/suc", handler.UploadSucHandler)
	http.HandleFunc("/file/meta", handler.GetFileMetaHandler)
	http.HandleFunc("/file/download", handler.DownloadHandler)
	http.HandleFunc("/file/update", handler.FileMetaUpdateHandler)
	http.HandleFunc("/file/delete", handler.FileDeleteHandler)
	http.HandleFunc("/user/signup", handler.SignupHandler)
	http.HandleFunc("/user/signin", handler.SigninHandler)
	http.HandleFunc("/user/info",
		handler.HTTPInterceptor(handler.UserInfoHandler),
	)

	if err := http.ListenAndServe(":8080", nil); err != nil {
		fmt.Printf("Failed to start server,err:%s", err.Error())
	}
}
