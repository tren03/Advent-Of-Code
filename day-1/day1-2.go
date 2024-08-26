package main

import (
	"fmt"
	"os"
	"unicode"
	"strings"
)

type nos struct{
	first int
	first_index int
	second int
	second_index int
}

func main(){
	data,_ := os.ReadFile("day1.txt")
	var str string = string(data)
	// fmt.Println(str)
	first:=0
	second:=0
	ans:=0
	// test:="eightwothree"
	// fmt.Println(substr_parse(test))
	// fmt.Println(number_parse(test))
	final_str:= strings.Split(str,"\n")
	for _,val := range final_str{
		fmt.Println(val)
		n1:=substr_parse(val)
		n2:=number_parse(val)

		// fmt.Println(n1)
		// fmt.Println(n2)

		//get first
		if n1.first==0 && n1.first_index==0{
			first=n2.first
		}else if n2.first==0 && n2.first_index==0{
			first=n1.first
		}else{
			if n1.first_index < n2.first_index{
				first=n1.first
			}else{
				first=n2.first
			}

		}
		
		
		//get second
		if n1.second==0 && n1.second_index==0{
			second=n2.second
		}else if n2.second==0 && n2.second_index==0{
				second=n1.second
		}else{
			if n1.second_index > n2.second_index{
				second=n1.second
	
			}else{
					second=n2.second
	
				
			}

		}

		temp:=first*10+second
		ans+=temp
		

	}
	fmt.Println(ans)
		

		


		
}
	



	// n_1:=substr_parse(test)
	// n_2:=number_parse(test)

	
	

func number_parse(str string)nos{
	n:=nos{first: 0,second: 0,first_index: 0,second_index: 0}
	flag:=false
	for index,val := range str{
		if unicode.IsDigit(val){
			if flag==false{
				n.first = int(val-'0')
				n.second=int(val-'0')
				n.first_index=index
				n.second_index=index
				flag=true
			}else{
				n.second=int(val-'0')
				n.second_index=index
			}

		}
	}
	return n

}

func substr_parse(str string) nos {
	m:=map[string]int{
		"one":1,
		"two":2,
		"three":3,
		"four":4,
		"five":5,
		"six":6,
		"seven":7,
		"eight":8,
		"nine":9,
	}

	flag := false
	n:=nos{first: 0,second: 0,first_index: 0,second_index: 0}
	for index,val := range str{
		substr:=""		
		if !unicode.IsDigit(val){
			for i:=index;i<index+5; i++ {
				if i<len(str){
					substr+=string(str[i])
				}
				
				_, exists:= m[substr]
				if exists && i < len(str){ 
					// fmt.Println(m[substr])
					// fmt.Println(index)
					if(flag==false){// this is first nos
						n.first=m[substr]
						n.second=m[substr]
						n.first_index=index
						n.second_index=index
						flag=true
					}else{
						n.second=m[substr]
						n.second_index=index
					
					}
				}
				
	
			}

		}
	}
	return n

}
