package main

import (
	"fmt"
	"log"
	"math"
	"os"
	"strconv"
	"strings"
	"unicode"
)

// the main logic is that -> use map for both set of numbers and check how many match

/*func getCards(data []byte) [][][]string {
	fmt.Println(data)
    var cards [][][] string
    for _,ch := range data{
        if ch == '\n'{
            cards = append(cards, [][]string{{}})
        }



    }
}*/

func parse_cards(cards string) int{
	numbers := strings.Split(cards, ":")
	fmt.Println(numbers[1])
	numberset := strings.Trim(numbers[1], " ")
	fmt.Println(numberset)

	index := 0
	flag := true

	//sets to store numbers
	my_nos := make(map[int]bool)
	other_nos := make(map[int]bool)
	for index < len(numberset) {
		if rune(numberset[index]) == '|' {
			flag = false
		}
		if unicode.IsDigit(rune(numberset[index])) {
			if index+1 < len(numberset) && unicode.IsDigit(rune(numberset[index+1])) {
				str := string(numberset[index]) + string(numberset[index+1])
				str_int, err := strconv.Atoi(str)
				if err != nil {
					log.Println("err conversion")
				}

				if flag {
					my_nos[str_int] = true
				} else {
					other_nos[str_int] = true
				}

				index += 2
				continue
			} else {
				str := string(numberset[index])
				str_int, err := strconv.Atoi(str)
				if err != nil {
					log.Println("err conversion")
				}

				if flag {
					my_nos[str_int] = true
				} else {
					other_nos[str_int] = true
				}

			}

		}
        index++
	}
    fmt.Println("my nos",my_nos)
    fmt.Println("other nos",other_nos)

    nos_matches:=0
    
    for key,_ := range my_nos{
        if _,exists := other_nos[key];exists{
            nos_matches++
        }
    }


    fmt.Println("nos matches",nos_matches)
    sum := math.Pow(2,float64(nos_matches)-1)
    fmt.Println("sum",sum)

    return int(sum)
    


    
}
func main() {
    data, err := os.ReadFile("../real_input.txt")
	if err != nil {
		log.Println("err reading file")
	}
	fmt.Println(string(data))
	sdata := string(data)

    
    
    ind := 0
    final_ans := 0
	cards := strings.Split(sdata, "\n")
    for ind<len(cards)-1{
        temp:=parse_cards(cards[ind])
        final_ans += temp 
        ind++
    }

   fmt.Println(final_ans) 




}
