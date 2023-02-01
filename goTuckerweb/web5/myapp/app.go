package myapp

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"time"

	"github.com/gorilla/mux"
)

//User Struct
type User struct{
	ID int `json:"id"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Email     string `json:"email"`
    CreateAt time.Time `json:"create_at"`
}

var userMap map[int]*User
var lastID int 

func indexHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Hello, World!")
}
func usersHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Get UserInfo by /users/{id}")
}

func getUserInfoHandler(w http.ResponseWriter, r *http.Request) {
//자동 parsing 
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err!= nil {
        w.WriteHeader(http.StatusBadRequest)
		fmt.Fprint(w, err)
		return
	}
	//아이디 존재 여부 확인 
	user, ok := userMap[id]
    if!ok {
		w.WriteHeader(http.StatusOK)
		fmt.Fprint(w, "No User ID:",id)
		return 
	}
	//확인된 아이디 리턴 
	w.Header().Add("content-type", "application/json")
	w.WriteHeader(http.StatusOK)
	data, _ := json.Marshal(user)
	fmt.Fprint(w, string(data))

}

func createUserHandler(w http.ResponseWriter, r *http.Request) {
	user := new(User)
	//json data parsing and input data to struct 
	err := json.NewDecoder(r.Body).Decode(user)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprint(w, err)
		return 
	}

	//유저 추가 정보 입력, map에 추가 
	lastID++
	user.ID = lastID
	user.CreateAt = time.Now()
	userMap[user.ID] = user 

	w.Header().Add("content-type", "application/json")
	w.WriteHeader(http.StatusCreated)
	data, _ := json.Marshal(user)
	fmt.Fprint(w, string(data))
}


func NewHandler() http.Handler{
	userMap = make(map[int]*User)
	lastID = 0
	mux := mux.NewRouter()
	
	mux.HandleFunc("/", indexHandler)
	mux.HandleFunc("/users", usersHandler).Methods("GET")
	mux.HandleFunc("/users", createUserHandler).Methods("POST")
	mux.HandleFunc("/users/{id:[0-9]+}", getUserInfoHandler)
	return mux
}