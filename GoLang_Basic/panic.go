package main

import ("fmt")
func handlePanic(){
	if a:=recover();a!=nil{
		fmt.Println("3 .RECOVER from -",a)
	}
	fmt.Println("4.I am stopping here by closing all connection")
}
func entry(language*string,name *string){
	fmt.Println("1-Entery Regualar")
	defer handlePanic()
	defer doCleanUp()
	if language ==nil{
		panic("Error:language cannot be nil")
	}
	if name==nil{
         panic("panic-because of name")
	}
	fmt.Println("language",language)
	fmt.Println("name",name)
}
func doCleanUp(){
	fmt.Println("2.defer-entry-cleanup")
}
func main(){
  defer fmt.Println("i am the first defer")
  defer fmt.Println("6.defer- in main")
  lang:="go lang"
  entry(&lang,nil)
  fmt.Println("5.main- normal",lang)
}