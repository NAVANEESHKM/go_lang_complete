package service

// import ("fullapi/models"
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
// func InsertProduct(product models.Product)(*mongo.InsertOneResult,error){
	
// 	 ctx,_:=context.WithTimeout(context.Background(),10*time.Second)
	 
// 	 result,err:=ProductContext().InsertOne(ctx,product)
// 	 if err!=nil{
// 		fmt.Println(err)
// 	 }
// 	 fmt.Println(result)
// 	 return result,nil
// }

// func InsertProductList(products []interface{})(*mongo.InsertManyResult,error){
	
// 	ctx,_:=context.WithTimeout(context.Background(),10*time.Second)
	
// 	result,err:=ProductContext().InsertMany(ctx,products)
// 	if err!=nil{
// 	   fmt.Println(err)
// 	}
// 	fmt.Println(result)
// 	return result,nil
// }

// func FindProducts()([]*models.Product,error){
// 	ctx,_:=context.WithTimeout(context.Background(),10*time.Second)
// 	filter:=bson.D{{"name","samsung"}}//here condition should be given

// 	result,err:=ProductContext().Find(ctx,filter)
//      if err!=nil{

// 		fmt.Println(err.Error())
// 		return nil,err
// 	 }else{
// 		fmt.Println(result)
// 		//build the array of products for the cursor that we received
//         var products []*models.Product
// 		for result.Next(ctx){
// 			product:=&models.Product{}
// 			err:=result.Decode(product)

// 			if err!=nil{
// 				return nil,err
// 			}
// 			fmt.Println(product)
// 			products=append(products,product)
// 		}
// 	    if err:=result.Err(); err!=nil{
// 		      return nil,err
// 	     }
// 	    if len(products)==0{
// 		      return []*models.Product{},nil
// 	    }
// 	     return products,nil
// 	 }
// }


