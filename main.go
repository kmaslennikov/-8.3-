package main

import (
	"encoding/json"
	"fmt"
	"log"
)

func main() {
	fmt.Print("Расчёт сдачи.\nВвод:")
	var input string
	_, _ = fmt.Scanln(&input)
	var bills []int
	err := json.Unmarshal([]byte(input), &bills)
	if err != nil {
		fmt.Println(err)
		log.Fatal()
	}

	fmt.Println("Вывод:", lemonadeChange(bills))
}

func lemonadeChange(bills []int) bool {
	var counter_5, counter_10, counter_20 int

	for i := 0; i < len(bills); i++ {
		switch bills[i] {
		case 5:
			counter_5++
			break
		case 10:
			counter_10++
			if counter_5 > 0 {
				counter_5--
			} else {
				return false
			}
		case 20:
			counter_20++
			if counter_10 > 0 && counter_5 > 0 {
				counter_10--
				counter_5--
			} else if counter_5 >= 3 {
				counter_5 -= 3
			} else {
				return false
			}

		}
	}
	return true
}
