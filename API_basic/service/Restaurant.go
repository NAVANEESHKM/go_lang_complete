package service

// import ("fullapi/models"
// 	"go.mongodb.org/mongo-driver/mongo/options"
// 	"fullapi/config"
// "go.mongodb.org/mongo-driver/mongo"
// "go.mongodb.org/mongo-driver/bson"
// "time"
// "fmt"
// "context"


// )
// func ProductContext() *mongo.Collection{
// 	client,_:=config.ConnectDataBase()
// 	return config.GetCollection(client,"sample_restaurants","restaurants")
// }
// func InsertProduct(product models.Restaurant)(*mongo.InsertOneResult,error){
	
// 	 ctx,_:=context.WithTimeout(context.Background(),100*time.Second)  //Sets deadline
	 
// 	 result,err:=ProductContext().InsertOne(ctx,product)  //Insert one data to the database
// 	 if err!=nil{
// 		fmt.Println(err)
// 	 }
// 	 fmt.Println(result)
// 	 return result,nil
// }

// func InsertProductList(products []interface{})(*mongo.InsertManyResult,error){
	
// 	ctx,_:=context.WithTimeout(context.Background(),100*time.Second)
	
// 	result,err:=ProductContext().InsertMany(ctx,products) //Inserts many data by storing it in array
// 	if err!=nil{
// 	   fmt.Println(err)
// 	}
// 	fmt.Println(result)
// 	return result,nil
// }

// func FindProducts()([]*models.Restaurant,error){
// 	ctx,_:=context.WithTimeout(context.Background(),100*time.Second)
// 	filter:=bson.D{}//here condition should be given
//      limit:=options.Find().SetLimit(10)
// 	result,err:=ProductContext().Find(ctx,filter,limit)
//      if err!=nil{

// 		fmt.Println(err.Error())
// 		return nil,err
// 	 }else{
// 		// fmt.Println(result)
// 		//build the array of products for the cursor that we received
//         var products []*models.Restaurant
// 		for result.Next(ctx){
// 			product:=&models.Restaurant{}
// 			err:=result.Decode(product)

// 			if err!=nil{
// 				return nil,err
// 			}
// 			//fmt.Println(product)
// 			products=append(products,product)
// 		}
// 	    if err:=result.Err(); err!=nil{
// 		      return nil,err
// 	     }
// 	    if len(products)==0{
// 		      return []*models.Restaurant{},nil
// 	    }
// 	     return products,nil
// 	 }
// }


