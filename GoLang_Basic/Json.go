package main

import ("fmt"
"net/http"
"log"
"io"
"encoding/json" 
)

type Article struct{
	Id string `json:"Id"`
	Title string `json:"Title"`
	Content string `json:"Content"`
	Summary string `json:"Summary"`
	Author string `json:"Author"`
}//golang to json mapping

func main(){
	http.HandleFunc("/json",handleArticle)
	http.HandleFunc("/getarticle",getArticles)
	log.Fatal(http.ListenAndServe(":3000",nil))
}
func handleArticle(w http.ResponseWriter,req *http.Request){
	if req.Method=="POST"{
		reqBody,_:=io.ReadAll(req.Body)
		var post Article
		err:=json.Unmarshal(reqBody,&post)
		post.Author="John"
		if err!=nil{
			fmt.Println("Got an error")
			fmt.Fprintf(w,err.Error())
		}else{
			json.NewEncoder(w).Encode(post)
			// fmt.Fprintf(w,string(newData))
		}
	}else{
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w,"unable to process the request")
	}
}
func getArticles(w http.ResponseWriter,req *http.Request){
	
	if req.Method=="POST"{
		reqBody,_:=io.ReadAll(req.Body)
		var articles[] Article
		err:=json.Unmarshal(reqBody,&articles)
		if err!=nil{
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w,err.Error())

	}else{
		fmt.Println(articles)
		articles=append(articles,Article{Id:"3",Title:"MIB"})
		json.NewEncoder(w).Encode(articles)
	}
}
}


