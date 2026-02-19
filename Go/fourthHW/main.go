package main

import "fmt"

func main() {

	// 1. Создать массив из 10 чисел и вывести
	arr := [10]int{3, 7, 12, 5, 8, 20, 1, 14, 6, 9}

	fmt.Println("Массив:")
	for i := 0; i < len(arr); i++ {
		fmt.Println(arr[i])
	}

	// 2. Сумма элементов
	sum := 0
	for i := 0; i < len(arr); i++ {
		sum += arr[i]
	}
	fmt.Println("Сумма:", sum)

	// 3. Макс и Мин
	max := arr[0]
	min := arr[0]

	for i := 1; i < len(arr); i++ {
		if arr[i] > max {
			max = arr[i]
		}
		if arr[i] < min {
			min = arr[i]
		}
	}
	fmt.Println("Максимум:", max)
	fmt.Println("Минимум:", min)

	// 4. Количество чётных
	evenCount := 0
	for i := 0; i < len(arr); i++ {
		if arr[i]%2 == 0 {
			evenCount++
		}
	}
	fmt.Println("Чётных чисел:", evenCount)

	// 5. Умножить на 2
	fmt.Println("Массив * 2:")
	for i := 0; i < len(arr); i++ {
		fmt.Println(arr[i] * 2)
	}

	// 6. Создать слайс и добавить элементы
	slice := []int{2, 4, 6, 8, 10}
	slice = append(slice, 12, 14, 16)
	fmt.Println("Слайс после append:", slice)

	// 7. Удалить элемент по индексу (например индекс 2)
	index := 2
	slice = append(slice[:index], slice[index+1:]...)
	fmt.Println("После удаления:", slice)

	// 8. Среднее арифметическое
	sum2 := 0
	for i := 0; i < len(slice); i++ {
		sum2 += slice[i]
	}
	average := float64(sum2) / float64(len(slice))
	fmt.Println("Среднее:", average)

	// 9. Самое короткое имя
	names := []string{"Ramazan", "Ali", "Zara", "Max", "John"}

	shortest := names[0]
	for i := 1; i < len(names); i++ {
		if len(names[i]) < len(shortest) {
			shortest = names[i]
		}
	}
	fmt.Println("Самое короткое имя:", shortest)

	// 10. Новый слайс с числами > 10
	numbers := []int{5, 11, 3, 20, 7, 15, 2, 30}
	var greaterThanTen []int

	for i := 0; i < len(numbers); i++ {
		if numbers[i] > 10 {
			greaterThanTen = append(greaterThanTen, numbers[i])
		}
	}
	fmt.Println("Числа больше 10:", greaterThanTen)
}