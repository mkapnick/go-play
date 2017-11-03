package hij

import "fmt"

func Hello() {
	hello := make(chan string, 3)

	hello <- "1st"
	hello <- "2nd"
	hello <- "3rd"

	fmt.Println(<-hello)
	fmt.Println(<-hello)
	fmt.Println(<-hello)

	// close the channel
	close(hello)
}
