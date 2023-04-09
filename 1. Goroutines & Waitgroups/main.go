package main

import (
	"fmt"
	"sync"
)

/* func printSomething(s string, wg *sync.WaitGroup) {

	defer wg.Done()

	fmt.Println(s)
}

func main() {

	var wg sync.WaitGroup

	words := []string{
		"alpha",
		"beta",
		"delta",
		"gamma",
		"pi",
		"zeta",
		"eta",
		"theta",
		"epsilon",
	}

	wg.Add(len(words))

	for i, word := range words {
		go printSomething(fmt.Sprintf("%d: %s", i, word), &wg)
	}

	wg.Wait()

	wg.Add(1)
	printSomething("This is the second thing to be printed!", &wg)
}
*/

var msg string

func updateMessage(s string, wg *sync.WaitGroup) {
	defer wg.Done()
	msg = s
}

func printMessage() {
	fmt.Println(msg)
}

func main() {
	msg = "Hello, world!"

	var wg sync.WaitGroup

	wg.Add(1)
	go updateMessage("Hello, universe!", &wg)
	wg.Wait()
	printMessage()

	wg.Add(1)
	go updateMessage("Hello, cosmos!", &wg)
	wg.Wait()
	printMessage()

	wg.Add(1)
	go updateMessage("Hello, world!", &wg)
	wg.Wait()
	printMessage()
}
