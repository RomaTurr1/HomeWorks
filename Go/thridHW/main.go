package main

import "fmt"

func task1() {
	a := [5]int{1, 2, 3, 4, 5}
	fmt.Println("Array:", a)
	fmt.Println("Length:", len(a))
}

func task2() {
	a := [4]int{10, 20, 30, 40}
	fmt.Println("Before:", a)
	a[1] = 200
	a[3] = 400
	fmt.Println("After:", a)
}

func task3() {
	a := [3]int{1, 2, 3}
	b := a
	a[0] = 100
	fmt.Println("a:", a)
	fmt.Println("b:", b)
	fmt.Println("b changed? -> No, arrays copy by value")
}

func task4() {
	var s []int
	fmt.Println("Value:", s)
	fmt.Println("len:", len(s))
	fmt.Println("cap:", cap(s))
	fmt.Println("is nil:", s == nil)
}

func task5() {
	s := []int{1, 2, 3}
	fmt.Println("Before len,cap:", len(s), cap(s))
	s = append(s, 4)
	fmt.Println("After len,cap:", len(s), cap(s))
}

func task6() {
	s := make([]int, 3, 5)
	s[0], s[1], s[2] = 10, 20, 30
	fmt.Println("Before append cap:", cap(s))
	s = append(s, 40, 50)
	fmt.Println("After append cap:", cap(s))
}

func task7() {
	a := [6]int{1, 2, 3, 4, 5, 6}
	s := a[2:5]
	s[0] = 999
	fmt.Println("Array:", a)
	fmt.Println("Slice:", s)
	fmt.Println("Array changed because slice references same underlying array")
}

func task8() {
	a := [6]int{1, 2, 3, 4, 5, 6}
	s1 := a[2:5]
	s2 := make([]int, len(s1))
	copy(s2, s1)

	s2[0] = 777
	fmt.Println("Array:", a)
	fmt.Println("s1:", s1)
	fmt.Println("s2:", s2)
	fmt.Println("Array unchanged because s2 is independent copy")
}

func modifySlice(s []int) {
	s[0] = 555
}

func task9() {
	s := []int{1, 2, 3}
	modifySlice(s)
	fmt.Println("Modified slice:", s)
}

func appendInside(s []int) {
	s = append(s, 999)
	fmt.Println("Inside function:", s)
}

func task10() {
	s := make([]int, 3, 3)
	s[0], s[1], s[2] = 1, 2, 3
	appendInside(s)
	fmt.Println("Outside function:", s)
	fmt.Println("Outside unchanged because append created new underlying array")
}

func main() {
	fmt.Println("Task 1"); task1()
	fmt.Println("\nTask 2"); task2()
	fmt.Println("\nTask 3"); task3()
	fmt.Println("\nTask 4"); task4()
	fmt.Println("\nTask 5"); task5()
	fmt.Println("\nTask 6"); task6()
	fmt.Println("\nTask 7"); task7()
	fmt.Println("\nTask 8"); task8()
	fmt.Println("\nTask 9"); task9()
	fmt.Println("\nTask 10"); task10()
}

// Practice

package main

import "fmt"

func main() {

	// 1 printInfo
	s1 := []int{1, 2, 3}
	printInfo(s1)

	// 2 setFirstZero
	setFirstZero(s1)
	fmt.Println("after setFirstZero:", s1)

	// 3 addNumber
	s1 = addNumber(s1, 99)
	fmt.Println("after addNumber:", s1)

	// 4 make + функция
	s2 := make([]int, 0, 5)
	s2 = addNumber(s2, 10)
	s2 = addNumber(s2, 20)
	s2 = addNumber(s2, 30)
	fmt.Println("make slice:", s2, "len:", len(s2), "cap:", cap(s2))

	// 5 append без return
	testAppendNoReturn(s2)
	fmt.Println("after appendNoReturn:", s2)

	// 6 slice из массива
	arr := [5]int{1, 2, 3, 4, 5}
	part := arr[1:4]
	changeSlice(part)
	fmt.Println("array after change:", arr)
	fmt.Println("slice:", part)

	// 7 рост capacity
	growCap()

	// 8 append 3 значения
	s3 := []int{1}
	s3 = addThree(s3)
	fmt.Println("after addThree:", s3)

	// 9 пустой slice
	var s4 []int
	s4 = addNumber(s4, 7)
	fmt.Println("empty slice after add:", s4)

	/* 10 Ответ:
	append может создать новый массив, если capacity не хватает,
	тогда функция меняет копию slice, а не оригинал из main. */
}

// 1
func printInfo(nums []int) {
	fmt.Println(nums, "len:", len(nums), "cap:", cap(nums))
}

// 2
func setFirstZero(nums []int) {
	if len(nums) > 0 {
		nums[0] = 0
	}
}

// 3
func addNumber(nums []int, value int) []int {
	return append(nums, value)
}

// 5
func testAppendNoReturn(nums []int) {
	nums = append(nums, 100)
	fmt.Println("inside function:", nums)
}

// 6
func changeSlice(s []int) {
	if len(s) > 0 {
		s[0] = 999
	}
}

// 7
func growCap() {
	s := make([]int, 0, 2)

	s = append(s, 1)
	fmt.Println("len:", len(s), "cap:", cap(s))

	s = append(s, 2)
	fmt.Println("len:", len(s), "cap:", cap(s))

	s = append(s, 3)
	fmt.Println("len:", len(s), "cap:", cap(s))

	s = append(s, 4)
	fmt.Println("len:", len(s), "cap:", cap(s))
}

// 8
func addThree(s []int) []int {
	return append(s, 10, 20, 30)
}