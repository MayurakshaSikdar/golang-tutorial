package main

import (
	"fmt"
	"maps"
	"reflect"
)

func main() {
	// ARRAY
	var a [4]int
	var b [4]string
	var c [4]float64
	var d [4]bool
	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)
	fmt.Println(d)

	fmt.Println(len(a))
	fmt.Println(reflect.TypeOf(a))

	nums := [3]int{1, 2, 3}
	fmt.Println(nums)

	dimarr := [2][2]int{{1, 2}, {3, 4}}
	fmt.Println(dimarr)

	// SLICES - dynamic arrays - most used construct in Golang
	var n []int
	fmt.Println(n)
	var n2 = make([]int, 2, 5)
	fmt.Println(n2)
	fmt.Println(cap(n2)) // to find out the capacity of a slice
	fmt.Println(len(n2))

	n2 = append(n2, 1)
	n2 = append(n2, 2)
	n2 = append(n2, 3)
	n2 = append(n2, 4)
	n2 = append(n2, 5)
	fmt.Println(cap(n2))
	fmt.Println(len(n2))

	n3 := []int{}
	n3 = append(n3, 3)
	n3 = append(n3, 1)
	fmt.Println(n3)

	var n4 = make([]int, 2, 5)
	var n5 = make([]int, 2, 5)
	n4 = append(n4, 1)
	n4 = append(n4, 2)
	copy(n5, n4)
	fmt.Println(n5)

	var n6 = make([]int, 2, 5)
	n6 = append(n6, 1)
	n6 = append(n6, 2)
	n6 = append(n6, 3)
	fmt.Println(nums[1:2])

	// MAP
	m := make(map[string]string)
	m["name"] = "Mayuraksha"
	m["surname"] = "Sikdar"
	fmt.Println(m)
	fmt.Println(m["name"])
	fmt.Println(m["non-exist"]) // map key not present. By default it returns zero value

	m2 := map[string]int{"zero": 0, "one": 1}
	fmt.Println(m2["two"])

	delete(m2, "zero")
	fmt.Println(m2)

	v, ok := m2["one"]
	fmt.Println(v, ok)
	if ok {
		fmt.Println("All Ok !")
	} else {
		fmt.Println("All Not Ok !")
	}

	fmt.Println(maps.Equal(m2, m2))

	// RANGE
	nm := []int{1, 2, 3, 4, 5}
	sum := 0
	for _, v := range nm {
		sum += v
	}
	fmt.Println(sum)

	nm2 := map[string]int{"zero": 0, "one": 1}
	for k, v := range nm2 {
		fmt.Println(k, " --> ", v)
	}

	for i, c := range "Golang" {
		fmt.Println(i, string(c))
	}
}
