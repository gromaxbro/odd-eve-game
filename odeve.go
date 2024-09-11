package main

import ("fmt"
	"math/rand/v2"
)
var not_done = true
var Your_points int = 0
var Com_points int = 0

func batting(){
	fmt.Println("Match started")
	for true{
		var i int
		var ran = rand.IntN(7)
		fmt.Print("->")
		fmt.Scan(&i)
		fmt.Println(ran)
		if i == ran{
			fmt.Println("OUT")
			fmt.Println("Your points:",Your_points)
			if not_done{}else{
				fmt.Println("Computer WON!!!!")
			}
			break
		} else{
			Your_points=Your_points+i
			if not_done{
			}else{
				if Your_points > Com_points {
					fmt.Println("YOU WON!!")
					break
				}
			}
			
		}
	}
	if not_done{
		not_done = false
		balling()
	}
}

func balling(){
	fmt.Println("Match started")
	for true{
		var i int
		var ran = rand.IntN(7)
		fmt.Print("->")
		fmt.Scan(&i)
		fmt.Println(ran)
		if i == ran{
			fmt.Println("OUT")
			fmt.Println("Computer points:",Com_points)
			if not_done{}else{
				fmt.Println("You WON!!!!")
			}
			break
		} else{
			Com_points=Com_points+ran
			if not_done{
			}else{
				if Your_points < Com_points {
					fmt.Println("Computer Won!!")
					break
				}
			}
			
		}
	}
	if not_done{
		not_done = false
		batting()
	}
}
func main(){
	var i int
	var chos string
	var got string

	// var  string
	var ran = rand.IntN(7)

	fmt.Println("Odd or Eve?")
	fmt.Scan(&chos)
	fmt.Print("-> ")
	fmt.Scan(&i)
	fmt.Printf("Computer : %v \n",ran)
	var mj = ran+i
	var ed int = mj%2
	if ed == 1{
		fmt.Println("Odd")
		got = "Odd"
	} else{
		fmt.Println("Even")
		got = "Eve"
	}
	if chos == got{
			var ch string
			fmt.Println("You Won Toss , Choose")
			fmt.Scan(&ch)
			if ch == "batting"{
				batting()
			}
			if ch == "balling"{
				balling()
			}
		} else{
			fmt.Println("You Lost Toss")
			var ch = rand.IntN(2)
			if ch == 1{
				fmt.Println("Computer choosed balling")
				batting()
			}
			if ch == 0{
				fmt.Println("Computer choosed batting")
				balling()
			}
		}
	// fmt.Println()
	// fmt.Println()
	
}