package main

import "fmt"

func main() {
	var myString string // Объявление переменной myString типа string и инициализация пустым значением

	var hello string = "\tHello\n" // Инициализация переменной hello типа string с текстом "    Hello" и переводом строки
	var title rune = '🌍'           // Инициализация переменной title типа rune с значением '🌍' (эмодзи "земля")

	word := "my string" // Инициализация переменной word типа string со значением "my string"

	str := "??"      // Инициализация переменной str типа string со значением "??"
	str = "hiAndrew" // Перезапись переменной str значением "hiAndrew"

	var b byte = 'c' // Объявление переменной b типа byte и инициализация символом 'c' (код символа 'c': 99 )
	var r rune = '🌍' // Объявление переменной r типа rune и инициализация значением '🌍' (эмодзи "земля")

	/*Строка последовательность байтов*/
	str = "some text"                                     // Перезапись переменной str значением "some text"
	fmt.Println("Длина строки some text:", len(str))      // Вывод длины строки (11)
	fmt.Println("Символ some text с индексом 1:", str[1]) // Вывод символа по индексу 1 (111) байтовое представление "о"
	// https://grandidierite.github.io/ASCII-table/

	str = "string" // Перезапись переменной str значением "string"

	fmt.Println("Результаты:")
	fmt.Println("title:", title)
	fmt.Println("myString:", myString)
	fmt.Println("hello:", hello)
	fmt.Println("b:", b)
	fmt.Println("r:", r)
	fmt.Println("word:", word)

	/*Мы можем обращаться к каждому байту по отдельности*/
	str = "12345"                              // Перезапись переменной str значением "12345"
	fmt.Println("Первый байт строки:", str[0]) // Вывод первого байта строки (49) байтовое представление числа 1
	fmt.Println("Второй байт строки:", str[1]) // Вывод второго байта строки (50) байтовое представление числа 2

	// https://autoit-script.ru/docs/appendix/ascii.htm?ysclid=lw61umpy3m515884225

	fmt.Println("Длина строки 12345:", len(str)) // Количество байтов 5

	str = "этоstring"                                // Перезапись переменной str значением "string"
	fmt.Println("Длина строки этоstring:", len(str)) // Вывод длины строки (12)

}
