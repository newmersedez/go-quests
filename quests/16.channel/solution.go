package channel

import "fmt"

// TODO:Implement ReadFromBoth function
// Read README.md for the instructions
func ReadFromBoth(ch1 chan string, ch2 chan string) string {
	val1 := <- ch1
	val2 := <- ch2
	return fmt.Sprintf("read: %s & %s", val1, val2)
}

// TODO:Implement ReadFromBoth function
// Read README.md for the instructions
func WriteToBoth(ch1 chan string, ch2 chan string, msg string) {
	go func () {
		ch1 <- fmt.Sprintf("write: %s", msg)
	}()

	go func() {
		ch2 <- fmt.Sprintf("write: %s", msg)
	}()
}

func ReadThenWrite(input chan string, output chan string) {
	val := <- input
	output <- fmt.Sprintf("transform: %s", val)
}
