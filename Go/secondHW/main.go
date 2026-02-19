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

// Practice
package main

import (
	"fmt"
)

func main() {

	averageAndSum()
	maxMin()
	evenOdd()
	calculator()
	checkPassword()
	countDigitsAndSum()
	multiplicationTable()

	// функции
	fmt.Println("add:", add(2, 3))
	fmt.Println("sub:", sub(5, 2))
	fmt.Println("mul:", mul(4, 3))
	fmt.Println("max:", max(7, 3))
	fmt.Println("min:", min(7, 3))

	fmt.Println("isEven:", isEven(6))
	fmt.Println("isPositive:", isPositive(-2))
	fmt.Println("abs:", abs(-10))
	fmt.Println("sumToN:", sumToN(5))
	fmt.Println("factorial:", factorial(5))

	fmt.Println("sumDigits:", sumDigits(123))
	fmt.Println("reverseNumber:", reverseNumber(123))
	fmt.Println("countDigits:", countDigits(12345))
	fmt.Println("isPrime:", isPrime(7))

	if res, ok := divide(10, 2); ok {
		fmt.Println("divide:", res)
	} else {
		fmt.Println("divide: нельзя делить на 0")
	}
}

// 1. Сумма и среднее
func averageAndSum() {
	var n, x int
	sum := 0

	fmt.Print("n: ")
	fmt.Scan(&n)

	for i := 0; i < n; i++ {
		fmt.Scan(&x)
		sum += x
	}

	avg := float64(sum) / float64(n)
	fmt.Printf("Sum: %d\nAverage: %.2f\n\n", sum, avg)
}

// 2. Max Min
func maxMin() {
	var n, x int

	fmt.Print("n: ")
	fmt.Scan(&n)

	fmt.Scan(&x)
	max := x
	min := x

	for i := 1; i < n; i++ {
		fmt.Scan(&x)
		if x > max {
			max = x
		}
		if x < min {
			min = x
		}
	}

	fmt.Println("Max:", max)
	fmt.Println("Min:", min)
	fmt.Println()
}

// 3. Чёт / нечёт
func evenOdd() {
	var n, x int
	even := 0
	odd := 0

	fmt.Print("n: ")
	fmt.Scan(&n)

	for i := 0; i < n; i++ {
		fmt.Scan(&x)
		if x%2 == 0 {
			even++
		} else {
			odd++
		}
	}

	fmt.Println("Even:", even)
	fmt.Println("Odd:", odd)
	fmt.Println()
}

// 4. Калькулятор
func calculator() {
	var a, b, op int

	fmt.Print("a b: ")
	fmt.Scan(&a, &b)

	fmt.Print("1:+ 2:- 3:* 4:/ -> ")
	fmt.Scan(&op)

	switch op {
	case 1:
		fmt.Println(a + b)
	case 2:
		fmt.Println(a - b)
	case 3:
		fmt.Println(a * b)
	case 4:
		if b == 0 {
			fmt.Println("Ошибка: деление на 0")
		} else {
			fmt.Println(a / b)
		}
	}
	fmt.Println()
}

// 5. Пароль
func checkPassword() {
	password := "admin123"
	var input string

	for i := 0; i < 3; i++ {
		fmt.Print("Password: ")
		fmt.Scan(&input)

		if input == password {
			fmt.Println("Доступ разрешён\n")
			return
		}
	}

	fmt.Println("Доступ заблокирован\n")
}

// 6. Цифры и сумма цифр
func countDigitsAndSum() {
	var n int
	fmt.Print("Number: ")
	fmt.Scan(&n)

	if n < 0 {
		n = -n
	}

	count := 0
	sum := 0

	for n > 0 {
		sum += n % 10
		n /= 10
		count++
	}

	fmt.Println("Digits:", count)
	fmt.Println("Sum digits:", sum)
	fmt.Println()
}

// 7. Таблица умножения
func multiplicationTable() {
	var n int
	fmt.Print("Number for table: ")
	fmt.Scan(&n)

	for i := 1; i <= 10; i++ {
		fmt.Printf("%d x %d = %d\n", n, i, n*i)
	}
	fmt.Println()
}

//// Лёгкие

func add(a, b int) int { return a + b }
func sub(a, b int) int { return a - b }
func mul(a, b int) int { return a * b }

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

//// Средние

func isEven(n int) bool { return n%2 == 0 }
func isPositive(n int) bool { return n > 0 }

func abs(n int) int {
	if n < 0 {
		return -n
	}
	return n
}

func sumToN(n int) int {
	sum := 0
	for i := 1; i <= n; i++ {
		sum += i
	}
	return sum
}

func factorial(n int) int {
	res := 1
	for i := 2; i <= n; i++ {
		res *= i
	}
	return res
}

//// Сложнее

func sumDigits(n int) int {
	if n < 0 {
		n = -n
	}
	sum := 0
	for n > 0 {
		sum += n % 10
		n /= 10
	}
	return sum
}

func reverseNumber(n int) int {
	rev := 0
	for n != 0 {
		rev = rev*10 + n%10
		n /= 10
	}
	return rev
}

func countDigits(n int) int {
	if n == 0 {
		return 1
	}
	if n < 0 {
		n = -n
	}
	count := 0
	for n > 0 {
		n /= 10
		count++
	}
	return count
}

func isPrime(n int) bool {
	if n < 2 {
		return false
	}
	for i := 2; i*i <= n; i++ {
		if n%i == 0 {
			return false
		}
	}
	return true
}

func divide(a, b int) (int, bool) {
	if b == 0 {
		return 0, false
	}
	return a / b, true
}