package main

import "fmt"
import "time"

func main() {

	//make a channel
	box := make(chan string)

	// create an anoymous goroutine 
	// and continue running main()
	go func() {
		fmt.Println("Working...")
		time.Sleep(5 * time.Second)
		// sending "Finished!" in the box
		box <- "Finished!"
	}()

	//start running
	fmt.Println("Starting...")

	// wait for box to arrive
	fmt.Println(<-box)
	
	// The End
	fmt.Println("Closing")
}