package main

import (
	"fmt"
	"os"
)

func logToStdout(message string) {
	fmt.Print(message)
}

func logToFile(message string) {
	file, err := os.Create("log.txt")
	if err != nil {
		fmt.Println("Error creating file:", err)
		return
	}
	defer file.Close()
	file.Write([]byte(message))
}

func main() {
	logToStdout("This is a log message to stdout\n")
	logToFile("This is a log message to a file\n")
}

/*В приведенном примере без использования интерфейсов у вас есть две функции: logToStdout и logToFile,
которые выполняют логгирование в стандартный вывод и в файл соответственно.
Если вы захотите добавить, например, логгирование в базу данных, вам придется написать новую функцию logToDatabase (например),
которая будет осуществлять запись в базу данных.

Использование интерфейсов в таких случаях позволяет абстрагировать различные способы логгирования за счет определения общего интерфейса (например, Logger) с методом для записи лога.
Тогда разные реализации этого интерфейса могут представлять разные способы логгирования (в файл, в базу данных, в сетевой сервис и т.д.),
но ваш основной код будет работать с интерфейсом Logger, не завися от конкретной реализации.

Каждая из этих функций обрабатывает логгирование в своей среде, и если вам нужно добавить "новую цель для логгирования",
то это означает добавление функции, которая будет логгировать в другое место или использует другой способ вывода:

Примеры "новых целей для логгирования" могут включать:

Логгирование в базу данных: Например, функция, которая сохраняет логи в базу данных вместо записи в файл или вывода в консоль.
Логгирование в сеть: Функция, которая отправляет логи по сети на удаленный сервер.
Логгирование в другие форматы: Например, функция, которая форматирует логи в JSON и сохраняет их в файл или отправляет по сети.
В примере без использования интерфейсов каждая "цель для логгирования" требует отдельной функции,
потому что каждая цель имеет свою уникальную реализацию. Это означает, что при добавлении новой цели нужно писать новую функцию,
что может привести к дублированию кода и усложнению поддержки при изменениях.*/
