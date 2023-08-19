package main

import ("fmt"
"net/http"
"log"
)
func main(){
	http.HandleFunc("/",handleHello)
		 http.HandleFunc("/products",handleProducts)
		 http.HandleFunc("/contacts",handleContacts)
		
        log.Fatal(http.ListenAndServe(":3000",nil))
		
}

func handleHello(w http.ResponseWriter,req *http.Request){
	fmt.Println(req.Method)
	fmt.Println(req.URL.Path)
   
   if req.URL.Path!="/"{
	    w.WriteHeader(http.StatusNotFound) //status code is important for front end development
		//StatusOk(200) StatusNotFound(404) StatusInternalServerError(500) StatusForbidden(403)
	    fmt.Fprintf(w,"Page Not Found")
   }else{
	      fmt.Fprintf(w,"<h1>welcome to HTTP</h1>")
   }
   
}
func handleProducts(w http.ResponseWriter,req *http.Request){
	fmt.Println(req.Method)
	fmt.Println(req.URL.Path)
   fmt.Fprintf(w,"<h1>products are here</h1>")
   
}
func handleContacts(w http.ResponseWriter,req *http.Request){
	fmt.Println(req.Method)
	fmt.Println(req.URL.Path)
   fmt.Fprintf(w,"<h1>navaneeshmuthusamy@gmail.com</h1>")
   
}