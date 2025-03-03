/*
Задача 1:
Описание задачи:
Создайте программу для управления библиотекой. Каждый экземпляр книги должен иметь название, автора, год издания и статус (доступна или на руках у читателя).
Добавьте возможность добавления новой книги, поиска книги по названию, выдачи книги читателю и возврата книги.

Требования:
Используйте структуры для представления книги.
Напишите методы для добавления книги, выдачи книги читателю и возврата книги.
Используйте функции для поиска книги по названию.
Используйте циклы для вывода списка всех книг.

Подсказка:
Метод Issue меняет статус книги на "на руках у читателя".
Метод Return меняет статус книги на "доступна".
*/
// package main

// import "fmt"

// // Структура книги
// type Book struct {
// 	Title  string
// 	Author string
// 	Year   int
// 	Status string
// }

// // Библиотека
// var library []Book

// // Добавление новой книги
// func addBook(title, author string, year int) {
// 	book := Book{Title: title, Author: author, Year: year, Status: "доступна"}
// 	library = append(library, book)
// 	fmt.Println("Книга добавлена:", book)
// }

// // Поиск книги по названию
// func findBook(title string) *Book {
// 	for i := range library {
// 		if library[i].Title == title {
// 			return &library[i]
// 		}
// 	}
// 	return nil
// }

// // Выдача книги читателю
// func (b *Book) Issue() {
// 	if b.Status == "доступна" {
// 		b.Status = "на руках у читателя"
// 		fmt.Println("Книга выдана:", b.Title)
// 	} else {
// 		fmt.Println("Книга уже на руках у читателя:", b.Title)
// 	}
// }

// // Возврат книги
// func (b *Book) Return() {
// 	if b.Status == "на руках у читателя" {
// 		b.Status = "доступна"
// 		fmt.Println("Книга возвращена:", b.Title)
// 	} else {
// 		fmt.Println("Книга уже доступна:", b.Title)
// 	}
// }

// // Вывод всех книг
// func listBooks() {
// 	fmt.Println("Список книг:")
// 	for _, book := range library {
// 		fmt.Printf("Название: %s, Автор: %s, Год: %d, Статус: %s\n", book.Title, book.Author, book.Year, book.Status)
// 	}
// }

// // пример вызова
// func main() {
// 	addBook("Гарри Поттер", "Дж. К. Роулинг", 1997)
// 	addBook("Война и мир", "Л. Н. Толстой", 1869)
// 	addBook("Мастер и Маргарита", "М. А. Булгаков", 1967)

// 	listBooks()

// 	book := findBook("Гарри Поттер")
// 	if book != nil {
// 		book.Issue()
// 	}

// 	listBooks()

// 	book.Return()
// 	listBooks()
// }

/*
Задача 2:

Описание задачи:

Создайте программу, которая поможет пользователям учитывать свои ежемесячные расходы.
Используйте карту, где ключом будет название категории расходов (например, "Продукты", "Транспорт", "Развлечения"), а значением - сумма расходов по этой категории.

Требования:
Используя функции
Пользователь должен иметь возможность добавлять новые категории и записывать расходы по каждой из них.
Также добавьте функцию для подсчета общей суммы расходов и вывода ее на экран
*/

// package main

// import "fmt"

// func main() {
// 	expenses := make(map[string]float64)

// 	// Добавление категории и расходов
// 	addExpense := func(category string, amount float64) {
// 		expenses[category] += amount
// 		fmt.Printf("Добавлено %.2f в категорию %s\n", amount, category)
// 	}

// 	// Подсчет общей суммы расходов
// 	calculateTotal := func() {
// 		total := 0.0
// 		for _, amount := range expenses {
// 			total += amount
// 		}
// 		fmt.Printf("Общая сумма расходов: %.2f\n", total)
// 	}

// 	addExpense("Продукты", 150.50)
// 	addExpense("Транспорт", 80.0)
// 	addExpense("Развлечения", 120.0)

// 	calculateTotal()
// }

/*
Задача 3:
Описание задачи:
1. Объяснить вывов при defer tryTest()()
2. Объяснить вывод при defer tryTest()

Объяснение:
defer откладывает выполнение указанного выражения до момента, когда выполнение функции, в которой он находится, завершится. При этом аргументы функции,
переданной в defer, вычисляются сразу же в момент объявления defer.
Функция tryTest выполняет fmt.Println("tryTest") и возвращает функцию, которая при вызове напечатает tryTest2.

В первом случае при вызове defer tryTest()() - tryTest() вызывается сразу же в момент объявления defer tryTest()(). Это напечатает trytest и возвращает функцию func()
{ fmt.Println("tryTest2") }. Затем эта возвращённая функция вызывается немедленно через defer, но её выполнение откладывается до завершения main.
Поэтому вызов этой функции произойдёт ближе к концу работы программы, после остальных deferred-выражений. Таким образом, последовательность deferred-выражений
выполняется в обратном порядке их объявления:
1. trytest
2. defer fmt.Println("Второе время", ...) → Выполняется первой, потому что была объявлена последней.
3. Затем вызывается отложенная функция tryTest2.
4. Наконец, выполняется defer fmt.Println("Первое время", ...).

Во втором случае при defer tryTest() результат выполнения tryTest() (возвращённая функция) не вызывается, так как она передаётся в defer как есть, т.е не два ().
Это значит, что в конце работы main вызов возвращённой функции не произойдёт, а только зарегистрируется её возврат.
Последовательность deferred-выражений:
- Сначала выполняется defer fmt.Println("Второе время", ...).
- trytest
- Затем defer fmt.Println("Первое время", ...).

*/

// package main

// import (
// 	"fmt"
// 	"time"
// )

// func tryTest() func() {
// 	fmt.Println("tryTest")
// 	return func() {
// 		fmt.Println("tryTest2")
// 	}
// }

// func main() {
// 	defer fmt.Println("Первое время:", time.Now())
// 	defer tryTest()()
// 	time.Sleep(2 * time.Second)

// 	defer fmt.Println("Второе время", time.Now())
// }