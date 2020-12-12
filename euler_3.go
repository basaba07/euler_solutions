package main

import "fmt"


func main (){

num := 600851475143

n := 2

for num != 1 {
	if num%n == 0{

		//fmt.Println("Prime Factor:",n)
		num = num/n
	} else {

		n = n+1

		//fmt.Println(n)
	}

}

fmt.Println("Last Prime Factor: ",n)

}