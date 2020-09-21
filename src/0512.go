package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
)
func refile(w http.ResponseWriter,r *http.Request){
	r.ParseMultipartForm(32<<20)//32kb
	file, header, err :=r.FormFile("upfile")
	if err!=nil{
		fmt.Println("报错")
		//log.Fatal(err)
		return
	}
	filename := header.Filename
	fmt.Printf("\n文件名字%s",filename)
	filesize := header.Size
	fmt.Printf("\n文件大小%d",filesize)
    fmt.Println(file)
    file.Close()
	newFile,err :=os.OpenFile("C:/Users/lei17/Desktop/"+"hello",os.O_CREATE|os.O_RDONLY|os.O_WRONLY,0666)
	if err != nil {                                                        //create rdonly wronly
		log.Fatal(err)
		return
	}
	defer newFile.Close()

	//2、将上传的file文件内容拷贝到新文件当中
	io.Copy(newFile,file)
}

func main(){

	http.HandleFunc("/",refile)
	err:= http.ListenAndServe("127.0.0.1:8080",nil)
	if err!=nil{
		log.Fatal(err)
		return
	}
}
