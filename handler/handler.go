package handler

import (
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
)

func UploadHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		// 返回 html页面
		data, err := ioutil.ReadFile("./static/view/index.html")
		if err != nil {
			io.WriteString(w, "internet server error")
			return
		}
		w.Write(data)
	} else if r.Method == "POST" {
		// 接收文件流及存储到本地目录
		file, head, err := r.FormFile("file")
		if err != nil {
			fmt.Printf("Failed to get data,err:%s\n", err.Error())
			return
		}
		defer file.Close()

		newFile, err := os.Create("/Users/liuhongrui/Desktop/temp/" + head.Filename)
		if err != nil {
			fmt.Printf("Failed to create file,err:%s\n", err.Error())
			return
		}
		defer newFile.Close()

		if _, err := io.Copy(newFile, file); err != nil {
			fmt.Println("Failed to save data into file,err:%s\n", err.Error())
			return
		}

		http.Redirect(w, r, "/file/upload/suc", http.StatusFound)
	}
}

func UploadSucHandler(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "Upload finished!")
}
