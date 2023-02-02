package main

import (
	"context"
	"fmt"
	"log"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Info struct{
	Name string `json:"name"`
	Age int `json:"age"`
	Place string `json:"place"`
}
func mongoConn()(client *mongo.Client){
	credential := options.Credential{
	   Username: "jungle",
	   Password: "jungle@123",
	}
	clientOptions := options.Client().ApplyURI("mongodb://3.39.23.91:27017").SetAuth(credential)
	client, err := mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
	   log.Fatal(err)
	}
 
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
	return 
}

func main() {
    client := mongoConn()
	fmt.Println(client)
	test(client)
	mongoDisConn(client)
}


func Insert(client *mongo.Client){
	collection := client.Database("test").Collection("tcollections")

	// Data to insert
	ash := Info{"Ash2", 10, "Pallet Town"}
	insertResult, err := collection.InsertOne(context.TODO(), ash)

	//To insert multiple documents	
	//trainers := []interface{}{misty, brock}
	//insertManyResult, err := collection.InsertMany(context.TODO(), trainers)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Inserted a single document: ", insertResult.InsertedID)
}

func Update(client *mongo.Client){
	updateResult, err := collection.UpdateOne(context.TODO(), filter, update)
if err != nil {
    log.Fatal(err)
}

fmt.Printf("Matched %v documents and updated %v documents.\n", updateResult.MatchedCount, updateResult.ModifiedCount)
}

func Find(client *mongo.Client){

}

func Delete(client *mongo.Client){

}
