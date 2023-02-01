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
	if len(userMap) == 0 {
		w.WriteHeader(http.StatusOK)
		fmt.Fprint(w, "No Users")
		return
	}
	users := []*User{}
	for _, u := range userMap {
		users = append(users, u)
	}
	data, _ := json.Marshal(users)
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	fmt.Fprint(w, string(data))
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

func deleteUserHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err!= nil {
        w.WriteHeader(http.StatusBadRequest)
		fmt.Fprint(w, err)
		return
	}
	//아이디 존재 여부 확인 
	_, ok := userMap[id]
	if!ok {
		w.WriteHeader(http.StatusOK)
		fmt.Fprint(w, "No User ID:",id)
		return 
	}
	//Map에서 삭제 
	delete(userMap, id)
	w.WriteHeader(http.StatusOK)
	fmt.Fprint(w, "Deleted User ID:",id)
}

func updateUserHandler(w http.ResponseWriter, r *http.Request) {
	//body에 유저 정보가 전달
	updateUser := new(User)
	err := json.NewDecoder(r.Body).Decode(updateUser)
	if err!= nil {
        w.WriteHeader(http.StatusBadRequest)
		fmt.Fprint(w, err)
        return 
	}
	user, ok := userMap[updateUser.ID]
	if!ok {
		w.WriteHeader(http.StatusOK)
		fmt.Fprint(w, "No User ID:",updateUser.ID)
		return 
	}
	if updateUser.FirstName != "" {
		user.FirstName = updateUser.FirstName
	}
	if updateUser.LastName != "" {
		user.LastName = updateUser.LastName
	}
	if updateUser.Email != "" {
		user.Email = updateUser.Email
	}

	w.Header().Add("content-type", "application/json")
	w.WriteHeader(http.StatusOK)
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
	mux.HandleFunc("/users", updateUserHandler).Methods("PUT")
	mux.HandleFunc("/users/{id:[0-9]+}", getUserInfoHandler).Methods("GET")
	mux.HandleFunc("/users/{id:[0-9]+}", deleteUserHandler).Methods("DELETE")

	return mux
}