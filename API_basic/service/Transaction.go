package service

import(
"fullapi/config"
"go.mongodb.org/mongo-driver/mongo"
"go.mongodb.org/mongo-driver/bson"
"encoding/json"
"time"
"fmt"
"context"
)
func help()*mongo.Collection{
           client,_:=config.ConnectDataBase()
           return config.GetCollection(client,"sample_analytics","transactions")
}
func Find_Transaction(){
	   ctx,_:=context.WithTimeout(context.Background(),100 *time.Second)
	    
	  
	//    result,err:=help().Find(ctx,filter,limit)
	   matchStage:=bson.D{{"$match",bson.D{{"transaction_count",100}}}}//talks about searching data
	   groupStage:=bson.D{
		{"$group",bson.D{
			{"_id","$account_id"},
			{"total_count",bson.D{{"$sum","$transactions.amount"}}},
		}}}  
		result,err:=help().Aggregate(ctx,mongo.Pipeline{matchStage,groupStage})
	 if err!=nil{
		  fmt.Println(err.Error())
		 
	 }else{
		var showsWithInfo[]bson.M
		if err=result.All(ctx,&showsWithInfo); err!=nil{
			  panic(err)
		}
		// fmt.Println(showsWithInfo)
		formatted_data,err:=json.MarshalIndent(showsWithInfo,""," ")
		if err!=nil{
			fmt.Println(err.Error())
		}else{
			fmt.Println(string(formatted_data))
		}
	 }
	 func UpdateTransaction(initialValue int,newValue int)(*mongo.UpdateResult,error){
		ctx,_ := context.WithTimeout(context.Background(),100*time.Second)
		filter := bson.D{{"accont_id",initialValue}}
		update := bson.D{{"$set",bson.D{{"account_id",newValue}}}}
		result,err := help().UpdateOne(ctx,filter,update)
		if err!=nil{
			fmt.Println(err.Error())
			return nil,err
		}
		return result, nil
	}
// 	   }else{
// 		var store[]*models.Trans
// 		      for result.Next(ctx){ //cursor is a build in function
//                     var tran=&models.Trans{}
// 					err:=result.Decode(tran)
// 					if err!=nil{
// 						return nil,err
// 					}
//                     store=append(store,tran)
// 			  }
// 			  return store,nil
// 	   }
	   

}