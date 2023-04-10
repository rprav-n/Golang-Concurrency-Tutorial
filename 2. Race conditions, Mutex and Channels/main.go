package main

import (
	"fmt"
	"math/rand"
	"time"

	"github.com/fatih/color"
)

/*
	sync.Mutex
	- Mutex = "mutual exclusion" - allows us to deal with race condition
	- Easy to use
	- Dealing with shared resources and concurrent/parallel goroutines
	- Lock/Unlock

	Race condition - occurs when multiple GoRoutines try to access the same data
	Can be difficult to spot when reading code
	Go allows us to check for them when running a program, or when testing out code with go test

	Channels
	- Means of having GoRoutines share data
	- Talk to each other
	- Go's philosophy - Share memory by communicating, rather than communicating by sharing memory
	- Producer/Consumer problem

*/

// Race Condition Example ------------------------------------------------

/* var msg string

var wg sync.WaitGroup

func updateMessage(s string, m *sync.Mutex) {
	defer wg.Done()

	// A thread safe operations
	m.Lock()
	msg = s
	m.Unlock()
}

func main() {
	msg = "Hello, world!"

	var mutex sync.Mutex

	wg.Add(2)

	go updateMessage("Hello, universe!", &mutex)
	go updateMessage("Hello, cosmos!", &mutex)
	wg.Wait()

	fmt.Println(msg)
}
*/

// Complex Example ------------------------------------------------

/* var wg sync.WaitGroup

type Income struct {
	Source string
	Amount int
}

func main() {
	// variable for bank balance
	var bankBalance int

	var balance sync.Mutex

	// print out starting values
	fmt.Printf("Initial account balance: $%d.00", bankBalance)
	fmt.Println()

	// define weely revenue
	incomes := []Income{
		{Source: "Main Job", Amount: 500},
		{Source: "Gifts", Amount: 10},
		{Source: "Part time job", Amount: 50},
		{Source: "Investments", Amount: 100},
	}

	wg.Add(len(incomes))

	// loop through 52 weeks and print out how much is made; keep a running total
	for i, income := range incomes {
		go func(i int, income Income) {
			defer wg.Done()
			for week := 1; week <= 52; week++ {

				balance.Lock()
				temp := bankBalance
				temp += income.Amount
				bankBalance = temp
				balance.Unlock()

				fmt.Printf("On week %d, you earned $%d.00 from %s\n", week, income.Amount, income.Source)
			}
		}(i, income)
	}

	wg.Wait()

	// print out final balance
	fmt.Printf("Final bank balance: $%d.00", bankBalance)
	fmt.Println()
} */

// Producer Consumer - Channenls ------------------------------------------------

const NumberOfPizzas = 10

var pizzasMade, pizzasFailed, total int

// Producer is a type of struct that holds two channels: one for pizzas, with all
// information for a given pizza order including whether it was made
// successfully, and another to handle end of processing (when we quit the channel)
type Producer struct {
	data chan PizzaOrder
	quit chan chan error
}

// PizzaOrder is a type of struct that describes a given pizza order. It has the order
// number, a message indicating what happened to the order, and a boolean
// indicating if the order was successfully completed.
type PizzaOrder struct {
	pizzaNumber int
	message     string
	success     bool
}

// Close is simply a method of closing the channel when we are done with it (i.e.
// something is pushed to quit channel)
func (p *Producer) Close() error {

	ch := make(chan error)
	p.quit <- ch
	return <-ch
}

// makePizza attempts to make a pizza. We generate a random number from 1-12
// and put in two cases where we can't make the pizza in time. Otherwise,
// we make the pizza without the issue. To make things interesting, each pizza
// will take a different length of time to produce (some pizzas are harder than others).
func makePizza(pizzaNumber int) *PizzaOrder {
	pizzaNumber++
	if pizzaNumber <= NumberOfPizzas {
		delay := rand.Intn(5) + 1
		fmt.Printf("Received order #%d!\n", pizzaNumber)

		rnd := rand.Intn(12) + 1
		msg := ""
		success := false

		fmt.Printf("Making pizza #%d. It will take %d seconds...\n", pizzaNumber, delay)

		// delay for a bit
		time.Sleep(time.Duration(delay) * time.Second)

		if rnd < 5 {
			pizzasFailed++
		} else {
			pizzasMade++
		}

		total++

		if rnd <= 2 {
			msg = fmt.Sprintf("*** We ran out of ingredients for pizza #%d!", pizzaNumber)
		} else if rnd <= 4 {
			msg = fmt.Sprintf("*** The cook quit while making pizza #%d!", pizzaNumber)
		} else {
			msg = fmt.Sprintf("Pizza order #%d is ready!", pizzaNumber)
			success = true
		}

		p := PizzaOrder{
			pizzaNumber: pizzaNumber,
			message:     msg,
			success:     success,
		}
		return &p

	}

	return &PizzaOrder{
		pizzaNumber: pizzaNumber,
	}
}

func pizzeria(pizzaMaker *Producer) {
	// keep track of which pizza we are making
	i := 0

	// run forever or until we receive a quit notification
	// this loop will continue to execute, trying to make pizzas,
	// until the quit channel receives something
	for {
		// try to make a pizza
		currentPizza := makePizza(i)
		if currentPizza != nil {
			i = currentPizza.pizzaNumber
			select { // Only used for channels
			// we tried to make a pizza (we sent something to the data channel)
			case pizzaMaker.data <- *currentPizza:

			// we want to quit, so send pizzaMaker.quit to quitChan (a chan error)
			case quitChan := <-pizzaMaker.quit:
				// close channels (Golden Rule: Always close channel when done with them)
				close(pizzaMaker.data)
				close(quitChan)
				return
			}
		}
		// decision
	}

}

func main() {
	// seed the random number generator
	rand.Seed(time.Now().UnixNano())

	// print out a message
	color.Cyan("The Pizzeria is open for business")
	color.Cyan("----------------------------------")

	// create a producer
	pizzaJob := &Producer{
		data: make(chan PizzaOrder),
		quit: make(chan chan error),
	}

	// run the producer in the background
	go pizzeria(pizzaJob)

	// create and run consumer
	for i := range pizzaJob.data {
		if i.pizzaNumber <= NumberOfPizzas {
			if i.success {
				color.Green(i.message)
				color.Green("Order #%d is out for delivery!\n", i.pizzaNumber)
			} else {
				color.Red(i.message)
				color.Red("The customer is really mad!")
			}
		} else {
			fmt.Println("Done making pizzas...")
			err := pizzaJob.Close()
			if err != nil {
				fmt.Println("*** Error closing channel!", err)
			}
		}
	}

	// print out the ending message
	color.Cyan("----------------------------------")
	color.Cyan("Done for the day.")

	color.Cyan("We made %d pizzas, but failed to make %d, with %d attempts in total", pizzasMade, pizzasFailed, total)

	switch {
	case pizzasFailed > 9:
		color.Red("It was an awful day...")
	case pizzasFailed >= 6:
		color.Red("It was not a very good day...")
	case pizzasFailed >= 4:
		color.Yellow("It was an okay day...")
	case pizzasFailed >= 2:
		color.Yellow("It was a pretty good day...")
	default:
		color.Green("It was a great day...")
	}
}
