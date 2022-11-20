package main

import (
	"fmt"
)

func main() {
	fmt.Println("Вместо числа кратного 3 выводится Fizz, а вместо числа кратного 5 выводим Buzz. Если число кратно пятнадцати, то программа должна выводить слово FizzBuzz")

	var number = 1

	for number < 100 {

		switch {
		case number%15 == 0:
			fmt.Println("FizzBuzz")
		case number%3 == 0:
			fmt.Println("Fizz")
		case number%5 == 0:
			fmt.Println("Buzz")
		default:
			fmt.Println(number)
		}
		number++
	}
}
