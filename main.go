package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
)

const maxNumber = 12307

func main() {
	input := bufio.NewScanner(os.Stdin)
	var number int

	if input.Scan() {
		var err error
		number, err = strconv.Atoi(input.Text())
		if err != nil {
			fmt.Println("Invalid input")
			return
		}
	}

	for number < maxNumber {
		if number < 0 {
			number *= -1
		} else if number%7 == 0 {
			number *= 39
		} else if number%9 == 0 {
			number = number*13 + 1
			continue
		} else {
			number = (number + 2) * 3
		}

		if number%13 == 0 && number%9 == 0 {
			fmt.Println("Service error")
			break
		} else {
			number++
		}
	}

	fmt.Printf("Final number is %d - ", number)
	printNumberToWords(number)
}

func printNumberToWords(number int) {
	if number == 0 {
		fmt.Println("ноль")
		return
	}
	for i := len(strconv.Itoa(number)) - 1; i >= 0; i-- {
		sign := number / int(math.Pow(10, float64(i))) % 10
		if sign == 0 {
			continue
		}
		switch i {
		case 5:
			fmt.Printf("%s ", tripleDigit[sign])
		case 4:
			if sign == 1 {
				nextSign := number / int(math.Pow(10, float64(i-1))) % 10
				fmt.Printf("%s тысяч ", doubleDigitWithOne[nextSign])
				i--
			} else {
				fmt.Printf("%s ", doubleDigitWithoutOne[sign])
			}
		case 3:
			if sign == 1 {
				fmt.Print("одна тысяча ")
			} else if sign == 2 {
				fmt.Print("две тысячи ")
			} else if sign >= 2 && sign <= 4 {
				fmt.Printf("%s тысячи ", singleDigit[sign])
			} else {
				fmt.Printf("%s тысяч ", singleDigit[sign])
			}
		case 2:
			fmt.Printf("%s ", tripleDigit[sign])
		case 1:
			if sign == 1 {
				nextSign := number / int(math.Pow(10, float64(i-1))) % 10
				fmt.Printf("%s ", doubleDigitWithOne[nextSign])
				i--
			} else {
				fmt.Printf("%s ", doubleDigitWithoutOne[sign])
			}
		case 0:
			fmt.Printf("%s ", singleDigit[sign])
		}
	}
}

var singleDigit = map[int]string{
	1: "один", 2: "два", 3: "три",
	4: "четыре", 5: "пять", 6: "шесть",
	7: "семь", 8: "восемь", 9: "девять",
}
var doubleDigitWithoutOne = map[int]string{
	2: "двадцать", 3: "тридцать", 4: "сорок",
	5: "пятьдесят", 6: "шестьдесят", 7: "семьдесят",
	8: "восемьдесят", 9: "девяносто",
}
var doubleDigitWithOne = map[int]string{
	0: "десять", 1: "одиннадцать", 2: "двенадцать", 3: "тринадцать",
	4: "четырнадцать", 5: "пятнадцать", 6: "шестнадцать",
	7: "семнадцать", 8: "восемнадцать", 9: "девятнадцать",
}
var tripleDigit = map[int]string{
	1: "сто", 2: "двести", 3: "триста",
	4: "четыреста", 5: "пятьсот", 6: "шестьсот",
	7: "семьсот", 8: "восемьсот", 9: "девятьсот",
}
