/*
Напишите программу, которая принимает строки через канал и подсчитывает количество слов в каждой строке. Используйте несколько горутин для
обработки строк и один канал для передачи результатов.

Условно, на вход строка
"Всем привет!
Следующая лекция в среду
Увидимся на лекции!
Результат
Word count: 2
Word count: 4
Word count: 3
*/

// package main

// import (
// 	"fmt"
// 	"strings"
// 	"sync"
// )

// func wordCountWorker(input <-chan string, results chan<- int, wg *sync.WaitGroup) {
// 	defer wg.Done()
// 	for line := range input {
// 		wordCount := len(strings.Fields(line))
// 		results <- wordCount
// 	}
// }

// func main() {
// 	lines := []string{
// 		"Всем привет!",
// 		"Следующая лекция в среду",
// 		"Увидимся на лекции!",
// 	}

// 	input := make(chan string, len(lines))
// 	results := make(chan int, len(lines))
// 	var wg sync.WaitGroup

// 	// Запускаем несколько воркеров
// 	numWorkers := 3
// 	for i := 0; i < numWorkers; i++ {
// 		wg.Add(1)
// 		go wordCountWorker(input, results, &wg)
// 	}

// 	// Отправляем строки в канал
// 	go func() {
// 		for _, line := range lines {
// 			input <- line
// 		}
// 		close(input)
// 	}()

// 	// Ожидаем завершения воркеров
// 	go func() {
// 		wg.Wait()
// 		close(results)
// 	}()

// 	// Читаем результаты
// 	for wordCount := range results {
// 		fmt.Println("Word count:", wordCount)
// 	}
// }

/*
Реализуйте простую модель многопользовательского чата, где несколько пользователей могут отправлять сообщения в общий канал. Каждое сообщение
должно содержать имя отправителя и текст сообщения. Создайте несколько горутин для имитации пользователей, которые отправляют сообщения.

примерный вывод:
[User3]: Message 1 from User3
[User1]: Message 1 from User1
[User2]: Message 1 from User2
[User2]: Message 2 from User2
[User1]: Message 2 from User1
[User3]: Message 2 from User3
[User3]: Message 3 from User3
[User2]: Message 3 from User2
[User1]: Message 3 from User1
[User2]: Message 4 from User2
[User3]: Message 4 from User3
[User1]: Message 4 from User1
[User1]: Message 5 from User1
[User3]: Message 5 from User3
[User2]: Message 5 from User2
*/
// package main

// import (
// 	"fmt"
// 	"math/rand"
// 	"sync"
// 	"time"
// )

// type Message struct {
// 	User    string
// 	Content string
// }

// func user(name string, chat chan<- Message, wg *sync.WaitGroup, numMessages int) {
// 	defer wg.Done()
// 	for i := 1; i <= numMessages; i++ {
// 		chat <- Message{
// 			User:    name,
// 			Content: fmt.Sprintf("Message %d from %s", i, name),
// 		}
// 		time.Sleep(time.Millisecond * time.Duration(rand.Intn(500)))
// 	}
// }

// func main() {
// 	rand.Seed(time.Now().UnixNano())

// 	chat := make(chan Message, 10)
// 	var wg sync.WaitGroup

// 	// Запускаем пользователей
// 	users := []string{"User1", "User2", "User3"}
// 	numMessages := 5
// 	for _, userName := range users {
// 		wg.Add(1)
// 		go user(userName, chat, &wg, numMessages)
// 	}

// 	// Запускаем горутину для закрытия канала после завершения пользователей
// 	go func() {
// 		wg.Wait()
// 		close(chat)
// 	}()

// 	// Читаем из канала и выводим сообщения
// 	for msg := range chat {
// 		fmt.Printf("[%s]: %s\n", msg.User, msg.Content)
// 	}
// }
