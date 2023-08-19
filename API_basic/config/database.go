package config

import (
	"context"
	"log"
	"fullapi/constants"
	
	"time"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)
func ConnectDataBase()(*mongo.Client,error){
	ctx,_:=context.WithTimeout(context.Background(),10*time.Second) //It is used to set the deadlines
	mongoconnection:=options.Client().ApplyURI(constants.ConnectionString)  //Used to create instance of MongoDB driver
	mongoClient,err:=mongo.Connect(ctx,mongoconnection)  //connects to mongodb
	if err!=nil{
		log.Fatal(err.Error())
		return nil,err
	}
	if err:=mongoClient.Ping(ctx,readpref.Primary()); err!=nil{  //It checks the connection
		return nil,err
	}
	return mongoClient,nil

}

func GetCollection(client*mongo.Client,dbName string,collectionName string)*mongo.Collection{
	// client,err:=ConnectDataBase()
	// if err!=nil{
	// 	fmt.Println(err)
	// 	panic(err)
	// }
	collection:=client.Database(dbName).Collection(collectionName)  //points to a database where the collection is stored
	return collection
}