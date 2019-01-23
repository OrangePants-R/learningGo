// Package main makes this an excecutable program
package main

import "runtime"
import "fmt"

/*
main_function
Go executes this program using this function

*/
func main() {
	fmt.Println(runtime.NumCPU())
}
