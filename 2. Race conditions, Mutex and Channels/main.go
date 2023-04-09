package main

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

type Producer struct {
	data chan PizzaOrder
	quit chan chan error
}

type PizzaOrder struct {
	pizzaNumber int
	message     string
	success     bool
}

func main() {
	// seed the random number generator

	// print out a message

	// create a producer

	// run the producer in the background

	// create and run consumer

	// print out the ending message
}
