package main

import "fmt"

func main() {
	var ptr *int // Нулевой указатель

	/*	Проверка на nil перед разыменованием
		if ptr != nil {
			fmt.Println("Значение, на которое указывает ptr:", *ptr)
		} else {
			fmt.Println("ptr равен nil, разыменование невозможно")

		}*/

	fmt.Println("Разыменовываем nil указатель", *ptr)
}

/*Разыменование нулевого указателя в Go приводит к панике и является распространенной ошибкой,
которую можно избежать, проверяя указатели на nil перед их использованием. */
