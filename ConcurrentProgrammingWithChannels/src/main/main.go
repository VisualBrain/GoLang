package main 

import (
	"fmt"
	"strings"
)

func main() {
	
	/*
	
	//Basic Channel
	
	/*
	channel := make(chan string) //string data flow through the channel
	fmt.Println(<-channel)  // <- is called receive operator
	//fatal error: all goroutines are asleep - deadlock! as the fmt.Println(<-channel) is waiting for message to receive making the others to wait
	
	channel := make(chan string) //string data flow through the channel
	channel <- "message" //results in deadlock - cause sending the message to channel causing the channel to block until something is able to receive
	fmt.Println(<-channel)  // <- is called receive operator
	----------------------------------->
	channel := make(chan string,1) //string data flow through the channel
	channel <- "message" 
	fmt.Println(<-channel)
	//it will print "message"
	*/
	
	/*
	Buffered Channel
	+Closing a Channel
	
	
	//Printing messages onebyone through channels
	
	phrase := "Package strconv implements conversions to and from string representations of basic data types.\n"
	//phrase2 := "The data flow analysis can track the result of mathematical operations and uses this information to warn you about conditions that are always true or false. We have improved the analysis of many math operations including multiplication, remainder, and bitwise operations.\n"
	
	words := strings.Split(phrase, " ")
	fmt.Println(words,len(words))
	//making channel
	//channel := make(chan string)  //results in deadlock as no buffer space 
	//Buffered Channel
	channel := make(chan string,len(words)) // 13 is the length of message without it results in deadlocks
	for _,word := range words {
		channel <- word // sending the messages through channel
	}
	close(channel) //only close the channel at sender side -> it makes receive believe that no further message is coming
	for i:=0;i<len(words);i++ {
		fmt.Println(<-channel+" ")
	}
	*/
	
	/*
	//Ranging over Channel - status-when receiver stop to receive the message 
	phrase := "Package strconv implements conversions to and from string representations of basic data types.\n"
	
	words := strings.Split(phrase, " ")
	fmt.Println(words,len(words))
	channel := make(chan string,len(words)) 
	for _,word := range words {
		channel <- word 
	}
	close(channel)
	for i:=0;i<len(words);i++ {
		fmt.Println(<-channel+" ") // receiver <- return two things - message and boolean Ok/NotOk(if msg is securely received)
	}
	*/
	phrase := "Package strconv implements conversions to and from string representations of basic data types.\n"
	
	words := strings.Split(phrase, " ")
	fmt.Println(words,len(words))
	channel := make(chan string,len(words)) 
	for _,word := range words {
		channel <- word 
	}
	close(channel)
//	for  {
//		if msg,ok := <-channel;ok {
//			fmt.Println(msg+" ")
//		}else{
//			break
//		}
//	}
    for msg := range channel{
    	fmt.Println(msg+" ")
    }
}

