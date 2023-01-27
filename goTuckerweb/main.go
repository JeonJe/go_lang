package main

import (
	"net/http"

	"github.com/JeonJe/goTuckerweb/myapp"
)
func main() {
	//웹서버 실행, 라우터 등록 
	http.ListenAndServe(":3000", myapp.NewHttpHandler())
}
  