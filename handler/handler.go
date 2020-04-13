package handler

import (
	"encoding/json"
	"filestore/meta"
	"filestore/util"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"time"
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

		fileMeta := meta.FileMeta{
			FileName: head.Filename,
			Location: "/tmp/" + head.Filename,
			UploadAt: time.Now().Format("2006-01-02 15:04:05"),
		}

		newFile, err := os.Create(fileMeta.Location)
		if err != nil {
			fmt.Printf("Failed to create file,err:%s\n", err.Error())
			return
		}
		defer newFile.Close()

		fileMeta.FileSize, err = io.Copy(newFile, file)
		if err != nil {
			fmt.Println("Failed to save data into file,err:%s\n", err.Error())
			return
		}

		newFile.Seek(0, 0)
		fileMeta.FileSha1 = util.FileSha1(newFile)
		meta.UpdateFileMeta(fileMeta)

		fmt.Printf("%+v", fileMeta)
		http.Redirect(w, r, "/file/upload/suc", http.StatusFound)
	}
}

func UploadSucHandler(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "Upload finished!")
}

// 获取文件元信息
func GetFileMetaHandler(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	// Form contains the parsed form data, including both the URL field's query parameters and the PATCH, POST, or PUT form data. This field is only available after ParseForm is called. The HTTP client ignores Form and uses Body instead
	// 表单包含已解析的表单数据，包括URL字段的查询参数和修补程序、POST或PUT表单数据。此字段仅在调用ParseForm后可用。HTTP客户端忽略Form，而使用Body
	fmt.Printf("%+v", r.Form["filehash"])
	filehash := r.Form["filehash"][0]
	fMeta := meta.GetFileMeta(filehash)
	data, err := json.Marshal(fMeta)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.Write(data)
}

func DownloadHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("DownloadHandler")
	r.ParseForm()
	// Get gets the first value associated with the given key. If there are no values associated with the key, Get returns the empty string. To access multiple values, use the map directly.
	// 获取与给定键关联的第一个值。如果没有与键关联的值，Get将返回空字符串。
	fsha1 := r.Form.Get("filehash")
	fm := meta.GetFileMeta(fsha1)

	fmt.Printf("%+v", fm)
	f, err := os.Open(fm.Location)
	if err != nil {
		fmt.Println("error1", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	defer f.Close()

	data, err := ioutil.ReadAll(f)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
	}
	w.Header().Set("Content-Type", "application/octect-stream")
	w.Header().Set("Content-disposition", "attachment;filename=\""+fm.FileName+"\"")
	w.Write(data)
}
