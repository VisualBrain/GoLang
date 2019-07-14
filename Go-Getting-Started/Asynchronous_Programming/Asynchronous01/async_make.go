package main

import (
	"fmt"
	"runtime"
	"sync"
	"time"
)

/*
Defer is used to ensure that a function call is performed later in a program’s execution
*/
/*
WaitGroups essentially allow us to tackle the problem of exiting from main without executing all goroutines by blocking the main the until any goroutines within that WaitGroup have successfully executed
*/

func main() {
	runtime.GOMAXPROCS(2)
	var waitgroup sync.WaitGroup
	waitgroup.Add(2) //setting the nuumber of goroutines we need to wait for
	go func() {
		defer waitgroup.Done()
		time.Sleep(5 * time.Second)
		fmt.Println("Hello")
	}()

	go func() {
		defer waitgroup.Done() //ensure that all operation is done before exiting from function
		fmt.Println("Working Walsjs")
	}()
	waitgroup.Wait()
}
