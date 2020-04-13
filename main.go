package main

import (
	"fmt"
	"net/http"
	"filestore/handler"
)

func main(){
	http.HandleFunc("/file/upload",handler.UploadHandler)
	http.HandleFunc("/file/upload/suc",handler.UploadSucHandler)
	http.HandleFunc("/file/meta",handler.GetFileMetaHandler)
	if err:=http.ListenAndServe(":8080",nil);err!=nil{
		fmt.Printf("Failed to start server,err:%s",err.Error())
	}
}