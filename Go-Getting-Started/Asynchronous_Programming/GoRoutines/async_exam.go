package main

/*

  Concurrent vs Parallel
	 in concurrent - there is only one cpu and GoRoutine serialize the operations by alloting each operation a certain amount of time to do execution
	 in parallel - there are many cpu's available then,schedular will allot each task to single cpu and allow them to execute/run to completion
  To detail with the problem of concurrency ,go provide two tools:
   Goroutines - are the name given to Go's threadlike constructs - used to launch a portion of code run concurrently,
   Go routine by definition required to do Asynchronous Programming,to call a function to goroutine we simply add go ahead of the funtion
  Channel - are used to sharing data between threads ,done by using a machinery that deal/handles with all of the synchronization of sender and receiver thread
  channel ensures that data is safely delivered to receiver
*/
import (
	"fmt"
	"runtime"
	"strings"
	"time"
)

func abcGen1() { /* 26 Go Routines work*/
	for i := byte('a'); i <= byte('z'); i++ {
		go println(string(i), "\t")
	}
}
func main() {
	/* To Get advantage of multiple cores of cpu  we used runtime*/
	runtime.GOMAXPROCS(8) //this will give unordered result
	fmt.Println("Asynchronous Programming")
	abcGen1()
	println(strings.Repeat("-", 34))
	println("Using Go Routine")
	go abcGen1() //Goroutine
	/*main function schedules the abcGen1() func but immediately terminates ,application exit before the function have a chance to run
	it didn't execute fully as main function doesn't give enough time to execute
	to run the function
	 we put the main function to sleep for some time .this allow GoScheduler to start running any pending task,
	 so as to give the function to complete its execution
	*/
	//	time.sleep(d.time Duration)
	println("This come first")
	time.Sleep(100 * time.Millisecond)
}

/*
In this currently ,28 go routines are working
go run - 1
abcGen1() - 26 Go routines' (in loop)
go abcGen1() - 1 Go routine
*/
