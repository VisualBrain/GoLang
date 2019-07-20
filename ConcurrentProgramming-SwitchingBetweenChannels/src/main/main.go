package main

import (
	"fmt" //allow us to print the outputs
)

type Message struct {
	To      []string
	From    string
	Content string
}

type FailedMessage struct {
	ErrorMessage    string
	OriginalMessage Message
}

// Switching Between Channels
func main() {

	//message channel
	channel01 := make(chan Message, 1) //create a channel of type Message and have buffer -1
	//error channel
	channel02 := make(chan FailedMessage, 1)
	//	msg := Message{
	//		To:      []string{"alvira@mail.com"},
	//		From:    "aviJay@mail.com",
	//		Content: "This is the Format of Tommorrow's Meeting.Be Prepared and Best of Luck ",
	//	}
	//
	//	failedMsg := FailedMessage{
	//		ErrorMessage:    "There is a Error Occured while Sending,Probably because of Network Issues!Please Check your Connection",
	//		OriginalMessage: Message{},
	//	}
	//	//
	//	//   channel01 <- msg //pushing message to channel
	//	//   channel02 <- failedMsg
	//	//
	//	//   fmt.Println("Message ",<-channel01)
	//	//   fmt.Println("Failed Message :", <-channel02)
	//
	//	//Job is to work with both channels simultaneously and work with channel that have message available
	//	//Whichever channel have message available will be executed first
	//	channel01 <- msg
	//	channel02 <- failedMsg
	select { //overloaded to handle multiple programming tasks
	case receivedMessage := <-channel01:
		fmt.Println(receivedMessage)
	case errorMessage := <-channel02:
		fmt.Println(errorMessage)
	default:
		fmt.Println("No Message Received ")
	}
}
