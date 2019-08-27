package main

import (
	"fmt"
	"sort"
)

//排序map
func main() {
	m1 := map[string]int{"cai": 1, "ao": 2, "oh": 3}
	fmt.Println("unsort:")
	for i, j := range m1 {
		fmt.Printf("key:%v  value:%v\n", i, j)
	}

	s1 := make([]string, len(m1))
	i := 0
	for k, _ := range m1 {
		s1[i] = k
		i++
	}

	sort.Strings(s1)
	fmt.Println("sorted:")
	for _, j := range s1 {
		fmt.Printf("key:%v  value:%v\n", j, m1[j])
	}

	t := 5
	fmt.Print(t)
}
