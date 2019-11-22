package mmio

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"
)

// to use:
// done := make(chan interface{})
// cr := mmio.CallResponse(done)
// go func() {
// 	for {
// 		select {
// 		case <-done:
// 			return
// 		case s := <-cr:
// 			fmt.Println("Echoing: ", s)
// 		}
// 	}
// }()
// <-done

// CallResponse is a pipeline that takes in a user's entry, and return it as a channel. ctrl-c ends the program
// modified from: https://www.reddit.com/r/golang/comments/4hktbe/read_user_input_until_he_press_ctrlc/
func CallResponse(done chan interface{}) <-chan string {
	sigs := make(chan os.Signal) // needed so the program can exit on ctrl-c
	signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM)

	call := make(chan string)
	go func() {
		defer close(call)
		for {
			select {
			case <-done:
				return
			default:
				var s string
				fmt.Scan(&s)
				call <- s
			}
		}
	}()

	response := make(chan string)
	go func() {
		defer close(response)
		for {
			select {
			case <-sigs:
				fmt.Println("got shutdown, exiting")
				close(done)
				return
			case <-done:
				return
			case s := <-call:
				response <- s
			}
		}
	}()
	return response
}
