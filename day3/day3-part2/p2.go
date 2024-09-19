package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"unicode"
)

var star_pos [][]int

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
				star_pos = append(star_pos, []int{i + 1, k, nos})
				return nos, true
			}
		}
		if start-1 >= 0 {
			if isSymbol(matrix[i+1][start-1]) {
				star_pos = append(star_pos, []int{i + 1, start - 1, nos})

				return nos, true
			}
		}
		if end < n {
			if isSymbol(matrix[i+1][end]) {
				star_pos = append(star_pos, []int{i + 1, end, nos})
				return nos, true
			}
		}

	}

	// check top row
	if i-1 < m && i-1 >= 0 {
		for k := start; k < end; k++ {
			if isSymbol(matrix[i-1][k]) {
				star_pos = append(star_pos, []int{i - 1, k, nos})
				return nos, true
			}
		}
		if start-1 >= 0 {
			if isSymbol(matrix[i-1][start-1]) {

				star_pos = append(star_pos, []int{i - 1, start - 1, nos})
				return nos, true
			}
		}
		if end < n {
			if isSymbol(matrix[i-1][end]) {

				star_pos = append(star_pos, []int{i - 1, end, nos})
				return nos, true
			}
		}

	}

	if start-1 >= 0 {
		if isSymbol(matrix[i][start-1]) {

			star_pos = append(star_pos, []int{i, start - 1, nos})
			return nos, true
		}
	}

	if end < n {
		if isSymbol(matrix[i][end]) {

			star_pos = append(star_pos, []int{i, end, nos})
			return nos, true
		}
	}
	return -1, false

}

// takes a string and returns true if it is a symbol
func isSymbol(str string) bool {
	for _, ch := range str {
		if ch == '*' {
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
    data, err := os.ReadFile("../1test.txt")
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

	// logic
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
				fmt.Println("this is the flag  ", flag)
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

	var vis []int
	for l := 0; l < len(star_pos); l++ {
		vis = append(vis, 0)
		fmt.Println(star_pos[l])

	}
	fmt.Println(vis)
	count := 0
	prod_ans := 0
	var temp int
	for i := 0; i < len(star_pos); i++ {
        count=0
		temp = 0
		// count the number of occurances
		for j := i + 1; j < len(star_pos); j++ {
			if star_pos[i][0] == star_pos[j][0] && star_pos[i][1] == star_pos[j][1] && vis[i] != 1 && vis[j] != 1 {
				count+=2

				fmt.Println(star_pos[i][0], star_pos[j][0])

			}
		}
		// if count == 2, only 2 adjancent, so process those
		if count == 2 {
			itr1 := i

			for itr2 := itr1 + 1; itr2 < len(star_pos); itr2++ {

				if star_pos[itr1][0] == star_pos[itr2][0] && star_pos[itr1][1] == star_pos[itr2][1] && vis[itr1] != 1 && vis[itr2] != 1 {
					fmt.Println("nos 1 = ", star_pos[itr1][2])

					fmt.Println("nos 2 = ", star_pos[itr2][2])
					temp = star_pos[itr1][2] * star_pos[itr2][2]
					vis[itr1] = 1
					vis[itr2] = 1

				}

			}

		}
		prod_ans += temp
	}

	fmt.Println("final ans = ", prod_ans)
	fmt.Println(vis)
}
