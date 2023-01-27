package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

func uploadsHandler(w http.ResponseWriter, r *http.Request) {
	uploadFile, header, err := r.FormFile("upload_file")
	
	if err!= nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprint(w, err)
		return
	}
	//upload 파일을 읽는 부분 
	defer uploadFile.Close()
	dirname := "./uploads"
	os.Mkdir(dirname,0777)
	filepath := fmt.Sprintf("%s/%s", dirname, header.Filename)
	file, err := os.Create(filepath)
	
	if err!= nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprint(w,err)
		return
	}
	defer file.Close()
	//업로드 파일을 write하는 부분
	io.Copy(file, uploadFile)
	w.WriteHeader(http.StatusOK)
	fmt.Fprint(w, filepath)

}

func main(){
	http.HandleFunc("/uploads", uploadsHandler)
	http.Handle("/", http.FileServer(http.Dir("public")))
	http.ListenAndServe(":3000", nil)
}