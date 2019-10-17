package main

import "fmt"

func finished() {
	fmt.Println("Finished finding largest")
}

func largest(nums []int) {
	defer finished() //Defer语句用于存在defer语句的函数返回之前执行函数调用

	fmt.Println("Started finding largest")

	max := nums[0]
	for _, v := range nums {
		if v > max {
			max = v
		}
	}

	fmt.Println("Largest number in", nums, "is", max)
}

func main() {
	nums := []int{78, 99, 234, 4, 35, 5234}
	largest(nums)
}
