package main

import "fmt"
import "time"

func main() {

	box := make(chan string)

	go func() {
		fmt.Println("Working...")
		time.Sleep(5 * time.Second)
		box <- "Finished!"
	}()

	fmt.Println("Starting...")

	fmt.Println(<-box)

	fmt.Println("Closing")
}
