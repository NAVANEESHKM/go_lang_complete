package main

import ("fmt"
        "sync")

func routine(i int,wg*sync.WaitGroup){
	fmt.Println("Started Routine",i)
	fmt.Printf("Routine %d ended \n",i)
	wg.Done()
}
func main(){
	no:=3
	var wg sync.WaitGroup
	for i:=1;i<no;i++{
		wg.Add(1)
		go routine(i,&wg)
	}
	wg.Wait()
	fmt.Println("all routines are finished")
}