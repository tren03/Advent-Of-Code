package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"unicode"
)

// for now just takes the start and end of the number and prints it
func check(matrix [][]string, i int, m int, n int, start int, end int) (int, bool) {

	var str string
	for j := start; j < end; j++ {

		fmt.Print(matrix[i][j])
		str += matrix[i][j]
	}
	nos, err := strconv.Atoi(str)
	fmt.Println("nos convertedd = ", nos)
	if err != nil {
		log.Println("error conv to int")
	}

	// check down dow
	if i+1 < m {
		for k := start; k < end; k++ {
			if isSymbol(matrix[i+1][k]) {
				return nos, true
			}
		}
		if start-1 >= 0 {
			if isSymbol(matrix[i+1][start-1]) {
				return nos, true
			}
		}
		if end < n {
			if isSymbol(matrix[i+1][end]) {
				return nos, true
			}
		}

	}

	// check top row
	if i-1 < m && i-1>=0 {
		for k := start; k < end; k++ {
			if isSymbol(matrix[i-1][k]) {
				return nos, true
			}
		}
		if start-1 >= 0 {
			if isSymbol(matrix[i-1][start-1]) {
				return nos, true
			}
		}
		if end < n {
			if isSymbol(matrix[i-1][end]) {
				return nos, true
			}
		}

	}

	if start-1 >= 0 {
		if isSymbol(matrix[i][start-1]) {
			return nos, true
		}
	}

	if end < n {
		if isSymbol(matrix[i][end]) {
			return nos, true
		}
	}
	return nos, false

}

// takes a string and returns true if it is a symbol
func isSymbol(str string) bool {
	for _, ch := range str {
		if !unicode.IsDigit(ch) && ch != '.' {
			return true
		}
	}
	return false
}

func isDot(str string) bool {

	for _, ch := range str {
		if ch == '.' {
			return true
		}
	}
	return false

}

// takes a string (which is 1 element in the matrix) and returns if it is a digit
func isItDigit(str string) bool {
	for _, ch := range str {
		if unicode.IsDigit(ch) {
			return true
		}
	}
	return false
}

// 2x2 matrix in go
var matrix [][]string

func main() {
	data, err := os.ReadFile("1test.txt")
	if err != nil {
		log.Println("read err")
	}
	fmt.Println(string(data))
	sdata := string(data)
	matrix = append(matrix, []string{})
	i := 0
	for _, ch := range sdata {

		if ch == '\n' {
			matrix = append(matrix, []string{})
			i++
			continue
		}
		matrix[i] = append(matrix[i], string(ch))
	}
	matrix = matrix[:len(matrix)-1]

	fmt.Println(matrix)
	m := len(matrix)
	n := len(matrix[0])
	j := 0

	fmt.Println(m)
	fmt.Println(n)
	fmt.Println("starting now")
	sum := 0
	i = 0
	for i < m {
		j = 0
		for j < n {
			if !isItDigit(matrix[i][j]) {
				// do nothin
			} else {
				start := j
				for j < n && isItDigit(matrix[i][j]) {
					j++

				}
				end := j
				nos, flag := check(matrix, i, m, n, start, end)
                fmt.Println("this is the flag  ",flag)
				if flag {
					sum += nos
				}

				fmt.Println("sum = ", sum)
			}

	fmt.Println(" ")
            j++
		}
		i++
	}

}
