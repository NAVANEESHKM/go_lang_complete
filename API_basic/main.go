package main

import(
	"fmt"
	// "fullapi/models"
	"fullapi/service"
	"go.mongodb.org/mongo-driver/mongo"
	// "go.mongodb.org/mongo-driver/bson/primitive"
	
)
var (
	
	mongoClient * mongo.Client  //act as global 
)
func main(){
      
	//   client,_:=config.ConnectDataBase()
    //   collection:=config.GetCollection(client,"sample_training","zips")
    
    fmt.Println("MongoDB successfully ")
	// product:=models.Product{ID:primitive.NewObjectID(),Name:"Nokia",Price:100000,Description:"Good Product"}
	// product:=[]interface{}{models.Product{ID: primitive.NewObjectID(),Name: "jio",Price: 14000,Description: "Good"},
    //     models.Product{ID: primitive.NewObjectID(),Name: "Nokia",Price: 11000,Description: "Good"}}
	// service.InsertProductList(product)

	// service.Find_Transaction()
	
	  
}