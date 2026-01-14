package tickers

import (
	"fmt"
	"time"
)

func Ticker() {
	helloTicker := time.NewTicker(500 * time.Millisecond)
	defer helloTicker.Stop()
	
	worldTicker := time.NewTicker(1000 * time.Millisecond)
	defer worldTicker.Stop()
	
	done := time.After(3 * time.Second)
	
	for {
		select{
		case <- done:
			return
		case <- helloTicker.C:
			fmt.Println("hello")
		case <- worldTicker.C:
			fmt.Println("world")
		}	
	}

}
