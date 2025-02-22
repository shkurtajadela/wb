/*
Цель: Создать HTTP-сервер на языке Go, который будет обрабатывать заявки студентов на поступление в университет. Сервер должен принимать
данные о студентах, проверять их баллы и выводить список поступивших студентов.

Задачи:
Создание структуры данных:
Определите структуру Student, которая будет содержать следующие поля:
FullName (строка) — полное имя студента.
MathScore (целое число) — балл по математике.
InformaticsScore (целое число) — балл по информатике.
EnglishScore (целое число) — балл по английскому языку.

Создание HTTP-сервера:
Реализуйте HTTP-сервер, который будет слушать на порту 8080.

Обработчик для поступления:
Создайте обработчик для POST-запросов на маршрут /apply, который будет принимать JSON с данными студента.
В обработчике проверьте, если сумма баллов по трем предметам (математика, информатика, английский) больше или равна 14, то добавьте студента в список поступивших. В противном случае, верните сообщение о том, что студент не поступил.

Обработчик для вывода поступивших студентов:
Создайте новый маршрут /admitted, который будет возвращать список всех студентов, которые поступили. Список должен быть представлен в формате JSON.

Создание студентов:
Создайте трех студентов (клиентов) с заранее определенными баллами:
Два студента должны иметь общую сумму баллов >= 14.
Один студент должен иметь общую сумму баллов < 14.
*/

package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"sync"
	"time"
)

// Структура
type Student struct {
	FullName         string
	MathScore        int
	InformaticsScore int
	EnglishScore     int
}

type Server struct {
	admittedStudents []Student
	mutex            sync.Mutex
}

// обработчик для поступления
func (s *Server) applyHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
		return
	}

	var student Student
	if err := json.NewDecoder(r.Body).Decode(&student); err != nil {
		http.Error(w, "Invalid JSON", http.StatusBadRequest)
		return
	}

	totalScore := student.MathScore + student.InformaticsScore + student.EnglishScore
	if totalScore >= 14 {
		s.mutex.Lock()
		s.admittedStudents = append(s.admittedStudents, student)
		s.mutex.Unlock()
		w.WriteHeader(http.StatusOK)
		fmt.Fprintf(w, "Congratulations, %s! You have been admitted.\n", student.FullName)
	} else {
		w.WriteHeader(http.StatusOK)
		fmt.Fprintf(w, "Sorry, %s. You did not meet the admission criteria.\n", student.FullName)
	}
}

// обработчик для вывода поступивших студентов
func (s *Server) admittedHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
		return
	}

	s.mutex.Lock()
	defer s.mutex.Unlock()
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(s.admittedStudents)
}

func createStudent(fullName string, math, informatics, english int) {
	student := Student{
		FullName:         fullName,
		MathScore:        math,
		InformaticsScore: informatics,
		EnglishScore:     english,
	}

	studentJSON, err := json.Marshal(student)
	if err != nil {
		fmt.Println("Error encoding student:", err)
		return
	}

	resp, err := http.Post("http://localhost:8080/apply", "application/json", bytes.NewBuffer(studentJSON))
	if err != nil {
		fmt.Println("Error sending request:", err)
		return
	}
	defer resp.Body.Close()

	fmt.Printf("Sent student: %s\n", fullName)
}

func main() {
	// Запускаем сервер в отдельной горутине
	go func() {
		server := &Server{
			admittedStudents: []Student{},
		}

		http.HandleFunc("/apply", server.applyHandler)
		http.HandleFunc("/admitted", server.admittedHandler)

		fmt.Println("Server is running on port 8080...")
		if err := http.ListenAndServe(":8080", nil); err != nil {
			fmt.Println("Error starting server:", err)
		}
	}()

	// Ждём, чтобы сервер успел запуститься
	// (в реальном проекте лучше использовать синхронизацию, например, sync.WaitGroup)
	select {
	case <-time.After(1 * time.Second):
	}

	// Создаём трёх студентов
	createStudent("Иван Иванов", 5, 5, 4)   // Сумма = 14 (поступил)
	createStudent("Мария Петрова", 6, 5, 4) // Сумма = 15 (поступила)
	createStudent("Олег Сидоров", 3, 4, 3)  // Сумма = 10 (не поступил)

	// Получение списка поступивших
	resp, err := http.Get("http://localhost:8080/admitted")
	if err != nil {
		fmt.Println("Error getting admitted students:", err)
		return
	}
	defer resp.Body.Close()

	var admitted []Student
	if err := json.NewDecoder(resp.Body).Decode(&admitted); err != nil {
		fmt.Println("Error decoding response:", err)
		return
	}

	fmt.Println("Admitted students:")
	for _, student := range admitted {
		fmt.Printf(" - %s\n", student.FullName)
	}

}
