package main

import (
	"fmt"
)

func main() {
	var n int
	fmt.Print("Enter size of i/p array: ")
	fmt.Scan(&n)
	//declare a slice of strings
	strs := make([]string, n)
	fmt.Println("Enter strings in the array:")
	for i := 0; i < n; i++ {
		fmt.Scan(&strs[i])
	}
	//common prefix
	ans := ""
	if len(strs) == 0 {
		fmt.Println("Common Prefix:-- ")
		return
	}
	//find minimum string
	mini := strs[0]
	for _, ele := range strs {
		if ele < mini {
			mini = ele
		}
	}
	//ans
	for i := 0; i < len(mini); i++ {
		char := mini[i]
		for _, word := range strs {
			if word[i] != char {
				fmt.Println("Common prefix(longest) :", ans)
				return
			}
		}
		ans += string(char)
	}
	fmt.Println("Common prefix(longest) :", ans)

}
