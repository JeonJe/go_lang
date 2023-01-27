package myapp

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

type fooHandler struct{}

type User struct{
	// json에서 어떻게 쓰는지 어노테이션을 붙여준다.
	FristName string 	`json:"first_name"`
	LastName  string 	`json:"last_name"`
	Email string		`json:"email"`
	CreatedAt time.Time `json:"created_at"`
}

func indexHandler(w http.ResponseWriter, r *http.Request){
	fmt.Fprint(w, "Hello, world!")
}

func (f *fooHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	
	// 인스턴스 생성 
	user := new(User)
	// request의 body에 들어있는 json 읽어들이기 
	err := json.NewDecoder(r.Body).Decode(user)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprint(w, err)
		return 
	}

	user.CreatedAt = time.Now()

	//user를 json으로 변환하여 응답
	w.Header().Add("Content-Type", "application/json")
	data, _ := json.Marshal(user)
	w.WriteHeader(http.StatusCreated)
	fmt.Fprint(w, string(data))
	
}

func barHandler (w http.ResponseWriter, r *http.Request) {
	name := r.URL.Query().Get("name")
	if name == "" {
        name = "World"
    }
	fmt.Fprintf(w, "Hello, %s!", name)
}

//Newhttp Handler
func NewHttpHandler() http.Handler{
	// 라우터 클래스를 만들어서 경로를 등록, 이제 http에 정적으로 등록하지 않음 
	mux := http.NewServeMux()
	//page path 와 핸들러 등록 
	mux.HandleFunc("/", indexHandler)

	// 핸들러 인스러 인스턴스를 만들어서 달아줌 
	mux.HandleFunc("/bar", barHandler)
	
	
	// Handler 는 인터페이스이고 함수를 하나 가지고 있음 
	// foo handler는 Serve HTTP를 구현하고 있는 인스턴스 
	mux.Handle("/foo",&fooHandler{})
	return mux
}