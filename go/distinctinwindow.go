package main

import (
	"fmt"
)

//use of map, array, for each loop

func helper(ar []int, k int) []int {
	//edge
	if k > len(ar) {
		return []int{}
	}
	//unordered map
	mp := make(map[int]int)
	ans := []int{}
	//first window
	for i := 0; i < k; i++ {
		mp[ar[i]]++
	}
	//push_back is append
	ans = append(ans, len(mp))
	//next
	for i := k; i < len(ar); i++ {
		mp[ar[i]]++
		//remove before
		mp[ar[i-k]]--
		if mp[ar[i-k]] == 0 {
			delete(mp, ar[i-k])
		}
		ans = append(ans, len(mp)) //size of map after each iteration
	}
	return ans
}

func main() {
	var n, k int
	fmt.Print("Enter array size: ")
	fmt.Scan(&n)
	//declare array
	ar := make([]int, n)
	fmt.Println("Enter elements")
	for i := 0; i < n; i++ {
		fmt.Scan(&ar[i])
	}
	fmt.Println("Enter window size k")
	fmt.Scan(&k)
	//result is array
	result := helper(ar, k)
	//for each in array
	for _, ele := range result {
		fmt.Print(ele, " ")
	}
	fmt.Println()
}

/*
q: Distinct numbers in a window of size k
vector<int> helper(vector<int> &ar, int k){
unordered_map<int,int> mp;
vector<int> ans;
if(k > ar.size()) return ans;
//first k elements
for(int i=0;i<k;i++){
	mp[ar[i]]++;
}
	ans.push_back(mp.size())
	//rem
	for(int i=k;i<ar.size();i++){
	mp[ar[i]]++;
	mp[ar[i-k]]--;
	if (mp[ar[i-k]] == 0) mp.erase(ar[i-k])
	ans.push_back(mp.size())
	}
	return ans;
}
*/
