package main

import "fmt"

/**
A goroutine is a lightweight thread managed by the Go runtime.
Goroutines run in the same address space, so access to shared memory must be synchronized.
Channels are a typed conduit through which you can send and receive values with the channel operator, <-.
The data flows in the direction of the arrow.
 */

func sumOfArrayElements(arr []int , c chan int) {
	sum := 0
	for _, v := range arr {
		sum += v
	}
	c <- sum
}


func bufferedChannel(){
	ch := make(chan int, 3)
	ch <- 1
	ch <- 2
	ch <- 3
	fmt.Println(<-ch, <-ch , <-ch)
}

func fibonacci(upperBound int, ch chan int) {
	x, y := 0, 1
	for i := 0; i < upperBound; i++ {
		ch <- x
		x, y = y, x+y
	}

	close(ch)
}

func rangeAndClose(){
  ch := make(chan int, 10)
  fibonacci(cap(ch), ch)
  for i := range ch {
  	fmt.Printf("%d ", i )
  }
  fmt.Printf("\n")
}


/**
The select statement lets a goroutine wait on multiple communication operations.
A select blocks until one of its cases can run, then it executes that case. It chooses one at random if multiple are ready.
 */


func selectFibonacci(fibChan, quitCh chan int) {
	x, y := 0, 1

	for {
		select {
		case fibChan <- x:
		     x, y = y, x + y

		case <- quitCh:
		   fmt.Println("quiting the channel")
		   return
		}
	}
}


func fibo2Select(){
	fibch := make(chan int)
	quitch := make(chan int)

	go func(){
		for i := 0 ; i < 10 ; i++ {
			fmt.Printf("%d ", <-fibch)
		}
		quitch <- 0
	}()

	selectFibonacci(fibch, quitch)
}


func main(){
	var arr = []int{1,2,3,4,5,6,7,8,9,1}
	// channel
	ch := make(chan int)
	// goRoutines
	go sumOfArrayElements(arr[:len(arr)/2], ch)
	go sumOfArrayElements(arr[len(arr)/2:], ch)
	x,y := <-ch, <-ch
	fmt.Println("x:", x, " y:", y, " x+y: ", x+y)

	// buffered channel
	bufferedChannel()

	// close channel
	rangeAndClose()

	// select
	fibo2Select()
}