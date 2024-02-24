// package name main - executable package -
// other package name - reusable package
package main

import "fmt"

func main() {

	var num int = 200
	var intNum16 int16 = 32767
	var intNum32 int32 = 2147483647

	var floatNum32 float32
	var floatNum64 float64

	// adding number into max int16 will mutate the ouput, but no error
	intNum16 += 1

	fmt.Println(num)
	fmt.Println(intNum16)
	fmt.Println(intNum32)
	fmt.Println(floatNum32)
	fmt.Println(floatNum64)
}
