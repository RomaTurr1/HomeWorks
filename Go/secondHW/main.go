package main

import "fmt"

// 1
func inRange(x int, min int, max int) bool {
	return x >= min && x <= max
}

// 2
func ageCategory(age int) string {
	if age < 12 {
		return "Child"
	} else if age <= 17 {
		return "Teen"
	} else if age <= 64 {
		return "Adult"
	}
	return "Senior"
}

// 3
func compareThree(a int, b int, c int) string {
	if a == b && b == c {
		return "all equal"
	} else if a == b || a == c || b == c {
		return "two equal"
	}
	return "all different"
}

// 4
func countDivisors(n int) int {
	count := 0
	for i := 1; i <= n; i++ {
		if n%i == 0 {
			count++
		}
	}
	return count
}

// 5
func firstEven() int {
	var x int
	for {
		fmt.Print("Enter number: ")
		fmt.Scan(&x)
		if x%2 == 0 {
			return x
		}
	}
}

// 6
func averageUntilZero() float64 {
	var x int
	sum := 0
	count := 0

	for {
		fmt.Print("Enter number (0 to stop): ")
		fmt.Scan(&x)
		if x == 0 {
			break
		}
		sum += x
		count++
	}

	if count == 0 {
		return 0
	}
	return float64(sum) / float64(count)
}

// 7
func readPositive() int {
	var x int
	for {
		fmt.Print("Enter positive number: ")
		fmt.Scan(&x)
		if x > 0 {
			return x
		}
	}
}

// 8
func safeCompare(a int, b int) (string, bool) {
	if a == b {
		return "equal", true
	} else if a > b {
		return "a greater", true
	}
	return "b greater", true
}

// 9 MENU
func menu() {
	for {
		fmt.Println("\n--- MENU ---")
		fmt.Println("1 - Проверка диапазона")
		fmt.Println("2 - Категория возраста")
		fmt.Println("3 - Количество делителей")
		fmt.Println("4 - Среднее до нуля")
		fmt.Println("0 - Выход")

		var choice int
		fmt.Print("Choose: ")
		fmt.Scan(&choice)

		switch choice {

		case 1:
			var x, min, max int
			fmt.Print("Enter x, min, max: ")
			fmt.Scan(&x, &min, &max)
			fmt.Println(inRange(x, min, max))

		case 2:
			var age int
			fmt.Print("Enter age: ")
			fmt.Scan(&age)
			fmt.Println(ageCategory(age))

		case 3:
			var n int
			fmt.Print("Enter number: ")
			fmt.Scan(&n)
			fmt.Println(countDivisors(n))

		case 4:
			fmt.Println("Average =", averageUntilZero())

		case 0:
			fmt.Println("Exit")
			return

		default:
			fmt.Println("Invalid option")
		}
	}
}

func main() {
	menu()
}