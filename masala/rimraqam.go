package main

import "fmt"

func main() {
	var s string = "III"
	var m map[string]int 
	m = make(map[string]int)
	m["I"] = 1
	m["V"] = 5
	m["X"] = 10
	m["L"] = 50
	m["C"] = 100
	m["D"] = 500
	m["M"] = 1000
	var n, res int = len(s), 0
	slice := s[:]
	for i := 1; i < n; i ++{
		if m[string(slice[i-1])] >= m[string(slice[i])] {
			res += m[string(slice[i-1])]
		}else{
			res -=  m[string(slice[i-1])]
		}
	}
	res += m[string(slice[n-1])]
	fmt.Print(res)
}
