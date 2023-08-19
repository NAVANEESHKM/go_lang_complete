package main

import ("fmt")

func main(){
	// var a=[5]string{"Alpha","beta","gama","delta","epp"}
	// var s[]string=a[0:4]
    // fmt.Println(a)
	// fmt.Println("slice",s)
	// fmt.Println(cap(s))
	// fmt.Println(len(s))


// 	a:=[7]string{"mon","tue","Wed","Thur","Fri","Sun"}
// 	slice1:=a[1:]
// 	slice2:=a[3:]
// 	fmt.Println("**Before modification**")
// 	fmt.Println(a)
//      fmt.Println(slice1)
// 	 fmt.Println(slice2)
// fmt.Println("**After modification**")
// 	 slice1[0]="fun"
// 	 slice2[0]="joy"
// 	 fmt.Println(a)
// 	 fmt.Println(slice1)
// 	 fmt.Println(slice2)


s:=[]int{10,20,30,40,50}
fmt.Println(s,len(s),cap(s))

s=s[1:5]
fmt.Println(s,len(s),cap(s))

s=s[:3]
fmt.Println(s,len(s),cap(s))

s=s[2:]
fmt.Println(s,len(s),cap(s))



// s1:=[]int{23,445,2,2,23}
// s2:=append(s1,11,111)
// fmt.Println(s1,len(s1),cap(s1))
// fmt.Println(s2,len(s2),cap(s2))

}