package selecttimeout

import (
	"fmt"
	"time"
)

func FunctionOrdered() {
	c1 := make(chan string)
	c2 := make(chan string)
	c3 := make(chan string)
	c4 := make(chan string)
	c5 := make(chan string)
	// TODO:Implement go-routines
	// Read README.md for the instructions
	go func() {
		time.Sleep(1 * time.Millisecond)
		c1 <- "from c1"
	}()
	go func() {
		time.Sleep(3 * time.Millisecond)
		c2 <- "from c2"
	}()
	go func() {
		time.Sleep(5 * time.Millisecond)
		c3 <- "from c3"
	}()
	go func() {
		time.Sleep(2 * time.Millisecond)
		c4 <- "from c4"
	}()
	go func() {
		time.Sleep(4 * time.Millisecond)
		c5 <- "from c5"
	}()
	for i := 0; i < 5; i++ {
		select {
		case msg1 := <-c1:
			fmt.Println(msg1)
		case msg2 := <-c2:
			fmt.Println(msg2)
		case msg3 := <-c3:
			fmt.Println(msg3)
		case msg4 := <-c4:
			fmt.Println(msg4)
		case msg5 := <-c5:
			fmt.Println(msg5)
		}
	}

}
