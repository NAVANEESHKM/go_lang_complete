package main

import("fmt")
//Array file
func main(){
	    var a=[3][3]int{{1,2,3},{4,5,6},{7,8,9}}
		var i,j int
		for i=0;i<3;i++ {
			for j=0;j<3;j++ {
				fmt.Println(a[i][j])
			}
			fmt.Println("\n")
		}

}
