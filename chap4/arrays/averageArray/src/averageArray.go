/*
** AUTHOR	: Robert James Patterson
** DATE		: 09/05/2019
** SYNOPSIS	: Work thru files for the 'An Introduction To Programming In Go' book
 */
package main

import (
	"fmt"
)

func main() {

	x := [5]float64{98, 93, 77, 82, 83}
	var total float64 = 0

	// In other traditional languages, we would write the 'for' loop as such:
	//
	// for i := 0; i < len(x); i++ {
	// 	total += x[i]
	//
	// } And this is how we would ay the same in GO:
	for _, value := range x {
		total += value
	}

	fmt.Println("============================")
	fmt.Printf("The average score is : %v\n", total/float64(len(x)))
	fmt.Println("============================")
}
