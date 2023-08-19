package main

import "fmt"

func help(str string, res string, k int) int {
	for i := 0; i < len(res); i++ {
		if str[k] != res[i] {
			return 0
		}
		k++
	}
	return 1
}
func main() {
	var c = 0
	var str string
	fmt.Println("Enter the String")
	fmt.Scanln(&str)
	var res = "go"
	for i := 0; i < len(str); i++ {
		if str[i] == 'g' {
			fin := help(str, res, i)
			if fin == 1 {
				c++
			}
		}
	}
	fmt.Println(c)
}
