package goroutine

import "fmt"


func SendRequest(input string) string {
	responses := make(chan string)
	defer close(responses)
	
	go Server(input, responses)
	
	response := <- responses
	return response

}

func Server(request string, response chan string) {
	response <- fmt.Sprintf("processed: %s", request)
}
