package main
import (
	"fmt"
	"strings"
	"os"
	"strconv"
)
// sample string to parse :        Game 1: 1 green, 4 blue; 1 blue, 2 green, 1 red; 1 red, 1 green, 2 blue; 1 green, 1 red; 1 green; 1 green, 1 blue, 1 red
func main(){
	data,_ := os.ReadFile("day2.txt")
	str_full := string(data)
	temp:=0
	// str:="Game 27: 9 blue, 5 red; 15 red, 12 blue, 3 green; 12 red, 12 blue, 1 green"
	ans:=0
	str_full_array := strings.Split(str_full,"\n")
	for _,line := range str_full_array{
		fmt.Println(line)
		// flag:=true
		// gets game id
		result:=strings.Split(line,":")
		// id_temp:=strings.Split(result[0]," ")
		// id:= id_temp[1]	


		// parsing string
		game_contents:= result[1]
		all_games:=strings.Split(game_contents,";")


		one_game := all_games[0]
		one_game=strings.Trim(one_game," ") // removes leading and trailing " "


		
		max_r,max_g,max_b := 0,0,0
		for _,one_game := range all_games{
			one_game=strings.Trim(one_game," ") // removes leading and trailing " "
			// fmt.Printf("%v\n",one_game)
			one_game_array:= strings.Split(one_game,",")
			// fmt.Println("this is one game array",one_game_array)
		
			for _,iter := range one_game_array{
				iter = strings.Trim(iter," ")
				nos := strings.Split(iter," ")
				// fmt.Printf("%v\n",nos)
				
				num, _ := strconv.Atoi(nos[0])
				// fmt.Println(num)


				if nos[1][0]=='r'{
					max_r = max(max_r,num)					
				}
				if nos[1][0]=='g'{
					max_g = max(max_g,num)		
				}
				if nos[1][0]=='b'{
					max_b = max(max_b,num)			
				}			
			
			}
			

		}
		temp=max_b*max_r*max_g
		fmt.Println(temp)
		ans+=temp
	}
	fmt.Println("ANS = ",ans)

}

