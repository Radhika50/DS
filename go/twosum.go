package main

import "fmt"

//use of map, array and mp.find in go--exists, append function

//Input:=[4,7,,2,9,13,5,3] target =12 , 2 elements from array will be equal to target, return indices

func main() {
	var target int
	var arsize int
	fmt.Print("enter size of array")
	fmt.Scan(&arsize)
	ar := make([]int, arsize)
	fmt.Println()
	fmt.Println("Enter elements")
	for i := 0; i < arsize; i++ {
		fmt.Scan(&ar[i])
	}
	fmt.Println("enter target")
	fmt.Scan(&target)
	//question
	//we need an hashmap and res vector
	res := []int{}
	mp := make(map[int]int)
	for i := 0; i < arsize; i++ {
		var search int = target - ar[i]
		//if  there return
		if idx, exists := mp[search]; exists {
			res = append(res, idx, i)
			break
		} else {
			mp[ar[i]] = i
		}
	}

	fmt.Printf("index values %v,%v gives target %v", res[0], res[1], target)
}

// if mp.find(mp[key])!=mp.end()-cpp

// if val, exists := mp[key]; exists {
// 	fmt.Println("Key found with value:", val)
// }
