package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

// Структура
type Student struct {
	FullName         string
	MathScore        int
	InformaticsScore int
	EnglishScore     int
}

// функция для создания студентов
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
	time.Sleep(1 * time.Second) // Дать серверу время на запуск

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

	fmt.Println("Поступившие студенты:")
	for _, student := range admitted {
		fmt.Printf(" - %s\n", student.FullName)
	}
}
