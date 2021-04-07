package main

import "fmt"

func main() {
	var min uint = 0

	if min <= 15 {
		fmt.Println("первая четверть")
	} else if min <= 30 {
		fmt.Println("вторая четверть")
	} else if min <= 45 {
		fmt.Println("третья четверть")
	} else if min <= 59 {
		fmt.Println("четвертая четверть")
	} else {
		fmt.Println("Введите число от 0 - 59")
	}
}
