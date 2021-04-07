package main

import "fmt"

func main()  {
	var result string

	fmt.Println(timeOfTheYear(result))
}

func timeOfTheYear(result string) string{

	var num int = 3

	if num == 1 {
		result = "Зима"
 	} else if num == 2 {
 		result = "Весна"
 	} else if num == 3 {
 		result = "Лето"
	} else if num == 4 {
		result = "Осень"
	} else {
		fmt.Println("Введите число от 1 до 4")
	}

	return result
}