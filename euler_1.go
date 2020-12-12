package main

import "fmt"


func main(){

	// obtain max number that can be multipled by 3,5,15 that results in a value less than 1000

	n_3 := 1000/3
	n_5 := (1000/5) - 1
	n_15 := 1000/15


	// compute the sums for multiples for each number * max number

	sum_3_multiples := ((n_3*(n_3+1))/2)*3
	sum_5_multiples := ((n_5*(n_5+1))/2)*5
	sum_15_multiples := ((n_15*(n_15+1))/2)*15


	// remove repeated factors that 3 and 5 share

	result := (sum_3_multiples + sum_5_multiples) - sum_15_multiples



	fmt.Println(result) 
}