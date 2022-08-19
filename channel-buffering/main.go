package main

import "fmt"

func main() {
	messages := make(chan string, 5)

	messages <- "buffered"
	messages <- "channel"
	messages <- "channel"
	messages <- "channel"
	messages <- "yoyo"

	fmt.Println(<-messages)
	fmt.Println(<-messages)
	fmt.Println(<-messages)
	fmt.Println(<-messages)
	fmt.Println(<-messages)

}
