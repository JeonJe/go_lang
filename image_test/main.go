package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type data struct{
	id int `json:"_id"`
	category string `json:"category,omitempty"`
	name string `json:"name"`
	num int `json:"num"`
	url string `json:"url"`
}

type RequestInfo struct{
	Category string `json:"category,omitempty"`
	Length int `json:"length"`
}

func CheckErr(err error){
	if err != nil {
		log.Fatal(err)
	 }
}

func mongoConn()(client *mongo.Client){
	credential := options.Credential{
	   Username: "jungle",
	   Password: "jungle@123",
	}
	clientOptions := options.Client().ApplyURI("mongodb://3.39.23.91:27017").SetAuth(credential)
	client, err := mongo.Connect(context.TODO(), clientOptions)
	CheckErr(err)
 
	// Check the connection
	err = client.Ping(context.TODO(), nil)
 
	if err != nil {
	   log.Fatal(err)
	}
	fmt.Println("MongoDB Connection Made")
	
	return client
}


func mongoDisConn(client *mongo.Client){

	err := client.Disconnect(context.TODO())

	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Connection to MongoDB closed.")
	
}

func InsertImage(client *mongo.Client){

}

func FindImage(client *mongo.Client){

}



func main() {

	http.ListenAndServe(":3000", newHandler())

}

func getImageHandler (w http.ResponseWriter, r *http.Request){
	
	// request의 body에 들어있는 json 읽어들이기 
	info := new(RequestInfo)
	err := json.NewDecoder(r.Body).Decode(info)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprint(w, err)
		return 
	}
	
	client := mongoConn()
	fmt.Println(client)
	defer mongoDisConn(client)
	
	collection := client.Database("test").Collection("FileInfo")

	numsOfNeed := int64(info.Length)

    // 필터 조건 선언 : "카테고리"로 들어온 데이터들을 n개 확인 
	filter := bson.D{primitive.E{Key : "category", Value: info.Category}}
	opts := options.Find().SetLimit(numsOfNeed).SetSort(bson.D{{Key : "num", Value: 1}})
	cursor, err := collection.Find(context.TODO(), filter, opts) 
	// 개수 

	if err != nil {
		log.Fatal(err)
	}

	// 데이터를 담을 변수 선언
	var results []bson.M
	if err = cursor.All(context.TODO(), &results); err != nil {
		panic(err)
	}

	//n개 정보 json으로 리턴 
	w.Header().Add("Content-Type", "application/x-www-form-urlencoded")
	data, _ := json.Marshal(results)
	w.WriteHeader(http.StatusOK)
	fmt.Fprint(w, string(data))
}


// server image path URL 요청시 파일 리턴
func getImageFileHandler(w http.ResponseWriter, req *http.Request){
	// 프로젝트 폴더부터 경로 시작시키기 위하여 받은 url 경로에 점 추가
	localPath := "." + req.URL.Path

	// 경로상 파일 읽기 및 에러리 404 리턴
	content, err := os.ReadFile(localPath)
	if err != nil {
		w.WriteHeader(404)
		w.Write([]byte(http.StatusText(404)))
		return
	}

	// mime type 헤더에 설정 및 content 리턴
	w.Header().Set("Content-Type", "image/jpeg; charset=utf-8")
	w.Write(content)
}

// 라우터
func newHandler() http.Handler {
	mux := mux.NewRouter()	
	
	mux.HandleFunc("/getImage", getImageHandler).Methods("POST")
	mux.HandleFunc("/image/", getImageFileHandler).Methods("GET")
	return mux
}