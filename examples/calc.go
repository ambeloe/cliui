package main

import (
	"cliui"
	"fmt"
	"strconv"
)

func main() {
	var menu cliui.UI

	menu.Name = "calc"

	menu.Add("add", add)
	menu.Add("sub", subtract)
	menu.Add("mul", multiply)
	menu.Add("div", divide)

	menu.Run()
}

func add(args []string) {
	var nums = make([]int, 0)
	var sum = 0

	for i, n := range args {
		var err error
		i, err = strconv.Atoi(n)
		if err != nil {
			fmt.Println("Invalid number provided")
			goto end
		}
		nums = append(nums, i)
	}
	for i := 0; i < len(nums); i++ {
		sum += nums[i]
	}
	fmt.Println("Sum: " + strconv.Itoa(sum))
end:
}

func subtract(args []string) {
	var nums = make([]int, 0)

	var sum = atoi(args[0])

	for i := 1; i < len(args); i++ {
		nums = append(nums, atoi(args[i]))
	}
	for i := 0; i < len(nums); i++ {
		sum -= nums[i]
	}
	fmt.Println("Sum: " + strconv.Itoa(sum))
}

func multiply(args []string) {
	var nums = make([]int, 0)

	var prod = atoi(args[0])

	for i := 1; i < len(args); i++ {
		nums = append(nums, atoi(args[i]))
	}
	for i := 0; i < len(nums); i++ {
		prod *= nums[i]
	}
	fmt.Println("Product: " + strconv.Itoa(prod))
}

func divide(args []string) {
	var nums = make([]int, 0)

	var quot = atoi(args[0])

	for i := 1; i < len(args); i++ {
		nums = append(nums, atoi(args[i]))
	}
	for i := 0; i < len(nums); i++ {
		quot /= nums[i]
	}
	fmt.Println("Quotient: " + strconv.Itoa(quot))
}

// atoi parses string to int with error catching; returns 0 if invalid
func atoi(x string) int {
	i, err := strconv.Atoi(x)
	if err != nil {
		fmt.Println("Invalid number provided")
		return 0
	}
	return i
}
