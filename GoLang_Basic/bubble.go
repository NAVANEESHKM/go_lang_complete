package main

import("fmt")

func main(){
	var arr=[]int{34,12,67,1}
	var i=0
	var j=0
	for i=0;i<len(arr);i++ {
		for j=i;j<len(arr);j++ {
			if(arr[i]<arr[j]){
				var temp=arr[i]
				arr[i]=arr[j]
				arr[j]=temp
			}
		}
	}
	for i=0;i<len(arr);i++ {
		fmt.Println(arr[i])
	}
}