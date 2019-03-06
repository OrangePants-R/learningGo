package main

import (
	"fmt"
	"time"
)

func printHello() {
	fmt.Println("Hello Go")
}

func main() {
	fmt.Println("Starting...")

	go printHello()

	time.Sleep(10 * time.Millisecond)
	fmt.Println("Ending...")

}
