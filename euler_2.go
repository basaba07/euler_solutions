package main

import "fmt"


func main(){

	a := 1
	b := 1
	i := 4000000
	sum := 0




	for {
	// check if the next fib number is less than 4M	
		if (a+b) < i{

			// check if next fib number is an even number
			if(a+b)%2 == 0{

				// compute sum 
				sum = sum + (a+b)
			}
			
			// holder for next fib num
			next_b := a + b
			// holder for previous fib num
			next_a := b
			// move to next fib num
			a = next_a

			// move to next fib num
			b = next_b 
		} else {
			break
		}
		}
	fmt.Println(sum)	
		
	}