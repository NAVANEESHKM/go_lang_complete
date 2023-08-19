package models

import ("go.mongodb.org/mongo-driver/bson/primitive"
       "time"
)


type Trans struct{
	ID primitive.ObjectID `json:"_id" bson:"_id"`
	Account_ID int `json:"account_id bson:"account_id,required"`
	Transaction_Count int `json:"transaction_count" bson:"transaction_count",required`
	Bucket_Start_Date primitive.DateTime `json:"bucket_start_date" bson:"bucket_start_date",required`
	Bucket_End_Date primitive.DateTime  `json:"bucket_end_date" bson:"bucket_end_date"`
	//Transactions []Transaction `json:"transactions" bson:"transactions",required`
}

type Transaction struct{
    Date time.Time `json:"date" bson:"date",required`
	Amount int `json:"amount" bson:"amount",required`
	Transaction_Code string `json:"transaction_code" bson:"tansaction_code",required`
	Symbol string `json:"symbol" bson:"symbol",required`
	Price string `json:"price" bson:"price",required`
	Total string `json:"total" bson:"total",required`
}
