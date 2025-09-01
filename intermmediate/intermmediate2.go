package main

import "fmt"

func main() {

	// FUNCTIONS
	fmt.Println(add(1, 3))
	fmt.Println(substract(1, 3))
	fmt.Println(languages())
	x, y, z := languages()
	print(x, " | ", y, " | ", z)
	// fn := func(a int) int {
	// 	return a
	// }
	fmt.Println(processit())

	// Variadic Functions
	fmt.Println(summation(1, 2, 3, 4, 5, 6, 7, 8, 9))
	q := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	fmt.Println(summation(q...))

	// Closures - outer function variable doesn't get reset from call stack
	increment := counter()
	fmt.Println(increment())
	fmt.Println(increment())
	fmt.Println(increment())
	fmt.Println(increment())
}

func add(a int, b int) int {
	return a + b
}

func substract(a, b int) int { // only single time Int is needed.
	return a - b
}

func languages() (string, string, string) {
	return "golang", "python", "JS"
}

func process(fn func(a int) int) {
	fn(1)
}

func processit() func(a int) int {
	return func(a int) int {
		return 4
	}
}

func summation(nums ...int) int {
	var total int
	total = 0
	for _, n := range nums {
		total += n
	}
	return total
}

func counter() func() int {
	var count int = 0
	return func() int {
		count += 1
		return count
	}
}
