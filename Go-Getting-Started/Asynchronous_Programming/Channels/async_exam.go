package main

/*
  Channel - are used to sharing data between threads ,done by using a machinery that deal/handles with all of the synchronization of sender and receiver thread
  channel ensures that data is safely delivered to receiver
*/
import (
	"fmt"
	"runtime"
	"strings"
	_ "time"
)

/*
func abcGen1(channel chan string) {
	for i := byte('a'); i <= byte('z'); i++ {
		//	go println(string(i), "\t")
		channel <- string(i) //Here <- is called this/receive operator and string(i) is the message to pass  ------Sender - sending the message
	}
	close(channel) //closing the channel - make channel unavailable to receive new Messages
}

func receiverMessagePrinter(channel chan string) {
	for {
		println(<-channel) // receiving the message
	}
}
*/

/*     To Eliminate the infinite for loop*/

func abcGen1(channel chan string) {
	for i := byte('a'); i <= byte('z'); i++ {
		//	go println(string(i), "\t")
		channel <- string(i) //Here <- is called this/receive operator and string(i) is the message to pass  ------Sender - sending the message
	}
	close(channel) //closing the channel - make channel unavailable to receive new Messages
}

func receiverMessagePrinter(channel chan string, doneChannel chan bool) {
	// for {
	// 	println(<-channel) // receiving the message
	// }

	for i := range channel {
		println(i)
	}
	doneChannel <- true

}

func main() {
	/* To Get advantage of multiple cores of cpu  we used runtime*/
	runtime.GOMAXPROCS(8) //this will give unordered result
	fmt.Println("Asynchronous Programming")
	channel := make(chan string) //chan is a keyword that describes that we are creating a channel
	doneChannel := make(chan bool)
	//abcGen1()
	println(strings.Repeat("-", 34))
	println("Using Channels")
	go abcGen1(channel)                             //Sender
	go receiverMessagePrinter(channel, doneChannel) // received in order manner and send true when done using doneChannel
	//	time.sleep(d.time Duration)
	println("This come first")
	//	time.Sleep(100 * time.Millisecond) eliminating time.sleep()
	<-doneChannel
}

/*
To eliminate the infinite loop
we close the channel (when sender done sending the message)  close(channel) then use for-loop to traverse over the channel to get received Messages

to eliminate the sleep- we create another channel

*/
