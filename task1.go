package main

import "fmt"

func myAdd(start int, num int) int {
	answer := start
	for i := 1; i <= num; i++ {
		temp := answer
		for temp != 0 {
			answer += temp % 10
			temp /= 10
		}
	}
	return answer
}

func Fibonnaci(n int) int {
	if n <= 0 {
		return 0
	} else if n == 1 {
		return 1
	} else {
		return Fibonnaci(n-1) + Fibonnaci(n-2)
	}
}

func pyramid(rows int) {
	for i := 1; i <= rows; i++ {
		num := 1

		for s1 := rows - i; s1 != 0; s1-- {
			fmt.Print(" ")
		}

		for m := 1; m <= 1+(i-1)*2; m++ {
			if m%2 == 1 {
				fmt.Print(num)
				num++
			} else {
				fmt.Print(" ")
			}
		}
		fmt.Println()
	}
}

func main() {

	var number int
	var iter int
	fmt.Scan(&number, &iter)

	fmt.Println(myAdd(number, iter))

	fmt.Println(Fibonnaci(12))
	var rows int
	fmt.Scan(&rows)
	pyramid(rows)
}
