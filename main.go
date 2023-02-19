package main

import (
	"fmt"
	"strconv"
)

func summa(a, b int) int {
	return a + b
}
func razn(a, b int) int {
	return a - b
}
func proiz(a, b int) int {
	return a * b
}
func delen(a, b int) int {
	return a / b
}

var prover1 int = 0

var roman = []string{
	"I", "II", "III", "IV", "V", "VI", "VII", "VIII", "IX", "X", "XI", "XII", "XIII", "XIV", "XV", "XVI", "XVII",
	"XVIII", "XIX", "XX", "XXI", "XXII", "XXIII", "XXIV", "XXV", "XXVI", "XXVII", "XXVIII", "XXIX", "XXX", "XXXI",
	"XXXII", "XXXIII", "XXXIV", "XXXV", "XXXVI", "XXXVII", "XXXVIII", "XXXIX", "XL", "XLI", "XLII",
	"XLIII", "XLIV", "XLV", "XLVI", "XLVII", "XLVIII", "XLIX", "L", "LI", "LII", "LIII", "LIV", "LV", "LVI", "LVII",
	"LVIII", "LIX", "LX", "LXI", "LXII", "LXIII", "LXIV", "LXV", "LXVI", "LXVII", "LXVIII", "LXIX", "LXX", "LXXI", "LXXII",
	"LXXIII", "LXXIV", "LXXV", "LXXVI", "LXXVII", "LXXVIII", "LXXIX", "LXXX", "LXXXI", "LXXXII", "LXXXIII", "LXXXIV",
	"LXXXV", "LXXXVI", "LXXXVII", "LXXXVIII", "LXXXIX", "XC", "XCI", "XCII", "XCIII", "XCIV", "XCV", "XCVI", "XCVII",
	"XCVIII", "XCIX", "C",
}

func rimch(s string) string {

	for k := 0; k < 20; k++ {
		if s == roman[k] {
			s = strconv.Itoa(k + 1)
			prover1++
		}

	}

	return s
}

func main() {

	var chislo1 string
	var oper string
	var chislo2 string
	for {
		prover1 = 0
		fmt.Println("Ведите мат. операцию")
		fmt.Scan(&chislo1, &oper, &chislo2)
		chislo1 = rimch(chislo1)
		chislo2 = rimch(chislo2)
		ch1, _ := strconv.Atoi(chislo1)
		ch2, _ := strconv.Atoi(chislo2)
		if ch1 < 1 || ch1 > 10 || ch2 < 1 || ch2 > 10 {
			fmt.Println("Неверное значение чисел")
			break
		}
		if prover1 == 1 {
			fmt.Println("Числа из разных систем")
			break
		}
		if prover1 == 2 {
			if oper == "+" {
				fmt.Println("Сумма равна", roman[summa(ch1, ch2)-1])
			} else if oper == "-" {
				razpol := razn(ch1, ch2)
				if razpol < 1 {
					fmt.Println("В римской системе нет отрецательных чисел")
					break
				}
				fmt.Println("Разница равна", roman[razpol-1])
			} else if oper == "*" {
				fmt.Println("Произведение равно", roman[proiz(ch1, ch2)-1])
			} else if oper == "/" {
				fmt.Println("Деление равно", roman[delen(ch1, ch2)-1])
			} else {
				fmt.Println("Неверное значение операции")
				break
			}
		} else if oper == "+" {
			fmt.Println("Сумма равна", summa(ch1, ch2))
		} else if oper == "-" {
			fmt.Println("Разница равна", razn(ch1, ch2))
		} else if oper == "*" {
			fmt.Println("Произведение равно", proiz(ch1, ch2))
		} else if oper == "/" {
			fmt.Println("Деление равно", delen(ch1, ch2))
		} else {
			fmt.Println("Неверное значение операции")
			break
		}

	}
}
