/*
** AUTHOR	: Robert James Patterson
** DATE		: 09/04/2019
** SYNOPSIS	: Work thru files for the 'An Introduction To Programming In Go' book
 */
package main

import (
	"fmt"
)

func main() {

	var x [5]int

	x[4] = 100

	fmt.Println("=====================================================")
	fmt.Printf("The array 'x' contains these five values: \n%v\n", x)
	fmt.Println("=====================================================")
}
