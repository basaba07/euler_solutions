package main

import "fmt"


func main (){

num := 121

n := 2

for num != 1 {
	if num%n == 0{

		fmt.Println(n)
		num = num/n
	} else {

		n = n+1
	}

}

}