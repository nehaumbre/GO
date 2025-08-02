package main

import "fmt"

func sendData(ch chan string) {
	ch <- "Hello from go routine"
}

//Channel Direction:You can restrict a function's channel use:
// func sendOnly(ch chan<- int) {
//     ch <- 42
// }
// func receiveOnly(ch <-chan int) {
//     fmt.Println(<-ch)
// }

//Buffered Channels
// A buffered channel allows sending without an immediate receiver:
// ch: make(chan int ,2) //buffer size 2
// ch <- 1
// ch <- 2
// ch <- 3 // Would block if buffer is full

//Closing Channels:You can close a channel to signal no more values will be sent:
// close(ch)

//Receivers can check if a channel is closed:
// value, ok := <-ch
// if !ok {
//     fmt.Println("Channel closed")
// }

//  Select Statement
// select lets you wait on multiple channel operations:
// select {
// case msg1 := <-ch1:
//     fmt.Println("Received:", msg1)
// case msg2 := <-ch2:
//     fmt.Println("Received:", msg2)
// default:
//     fmt.Println("No messages")

// Loop over a channel until it’s closed: Using Channels with range
//go func(ch chan int) {
//     for i := 0; i < 5; i++ {
//         ch <- i
//     }
//     close(ch)
// }(ch)

// for val := range ch {
//     fmt.Println(val)
// }

func main() {
	//Channels: Safe Communication Between Goroutines
	//They let one goroutine send data to another safely and synchronously (by default).
	// ch := make(chan int) // Create a channel of type int
	// ch <- 10             // Send 10 to the channel
	// value := <-ch        // Receive from the channel
	// fmt.Println(value)   // Prints: 10

	ch := make(chan string) // Create a channel

	go sendData(ch) //start go routine

	message := <-ch      // Receive message (blocks until sent)
	fmt.Println(message) // Prints: Hello from go routine

	unbufferChannel := make(chan int)

	bufferChannel := make(chan string, 5)

	go func() {
		unbufferChannel <- 2
	}()

	go func() {
		bufferChannel <- "buff1"
		bufferChannel <- "buff2"
		bufferChannel <- "buff3"
		bufferChannel <- "buff4"
		bufferChannel <- "buff5"
		close(bufferChannel)
	}()

	valUnbuff := <-unbufferChannel
	fmt.Println(valUnbuff) // Prints: 2

	for msg := range bufferChannel {
		fmt.Println(msg) // Prints: buff1, buff2..
	}
}

// 🔹 Unbuffered Channels in Go
// An unbuffered channel is a channel with no capacity, meaning:
// Send and receive must happen at the same time — otherwise, the goroutine blocks.
// If you try to send to an unbuffered channel when no goroutine is receiving from it
// (i.e., no goroutine is waiting on the channel), the sender will block until a
// receiver is ready to receive from the channel.
