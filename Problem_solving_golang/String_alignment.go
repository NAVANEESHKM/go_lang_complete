package main

import "fmt"

func main() {
	var don string
	fmt.Println("Enter the String")
	fmt.Scanln(&don)
	
	var str = []byte(don)
	for i := 0; i < len(str); i++ {

		if i%2 == 0 && i != 0 {

			var temp = str[i+1]
			str[i+1] = str[i]
			str[i] = temp
			i = i + 3
		}

	}
	mstr := string(str)
	fmt.Println(mstr)
}
