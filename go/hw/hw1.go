/*
Задача:
Дан слайс. Необходимо заменить четные числа на 1 и посчитать сумму чисел в новом слайсе
numbers := []int{3, 5, 7, 2, 7, 8, 6, 4, 7, 0, 1, 7, 4, 8, 10, 3, 6, 8, 5, 4, 12, 3}
*/

// package main

// import "fmt"

// func main() {
// 	numbers := []int{3, 5, 7, 2, 7, 8, 6, 4, 7, 0, 1, 7, 4, 8, 10, 3, 6, 8, 5, 4, 12, 3}
// 	newSlice := make([]int, len(numbers))
// 	sum := 0

// 	for i, num := range numbers {
// 		if num%2 == 0 {
// 			newSlice[i] = 1
// 		} else {
// 			newSlice[i] = num
// 		}
// 		sum += newSlice[i]
// 	}

// 	fmt.Println("Исходный слайс:", numbers)
// 	fmt.Println("Новый слайс:", newSlice)
// 	fmt.Println("Сумма чисел в новом слайсе:", sum)
// }

/*
Создайте слайс целых чисел и заполните его числами от 1 до 10.
Используя цикл, пройдите по слайсу и увеличьте каждое значение на 5, используя указатель.
Выведите измененный слайс.
*/
// package main

// import "fmt"

// func main() {
// 	numbers := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

// 	for i := range numbers {
// 		ptr := &numbers[i]
// 		*ptr += 5
// 	}

// 	// Выводим измененный слайс
// 	fmt.Println("Modified slice:", numbers)
// }

/*
Дан слайс чисел, необходимо найти минимальное и максимальное значение, которое делится на 2 без остатка.
numbers := []int{8, 44, 3, 5, 11, 8, 2, 10, 6, 77, 15, 12}
*/
// package main

// import (
// 	"fmt"
// 	"math"
// )

// func main() {
// 	numbers := []int{8, 44, 3, 5, 11, 8, 2, 10, 6, 77, 15, 12}

// 	var minEven, maxEven int
// 	minEven = math.MaxInt
// 	maxEven = math.MinInt

// 	for _, num := range numbers {
// 		if num%2 == 0 {
// 			if num < minEven {
// 				minEven = num
// 			}
// 			if num > maxEven {
// 				maxEven = num
// 			}
// 		}
// 	}

// 	fmt.Println("Минимальное четное значение:", minEven)
// 	fmt.Println("Максимальное четное значение:", maxEven)
// }
