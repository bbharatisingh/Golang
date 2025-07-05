package concurrency

import (
	"fmt"
	"reflect"
	"sync"
)

func BasicChanOp(nums []int) {
	// fmt.Println("basic op", nums)
	inputChan := make(chan int)

	wc := 2
	i := 0
	var wg sync.WaitGroup

	opList := make([]<-chan int, wc)

	for i < wc {
		wg.Add(1)
		opList[i] = readNprint(inputChan, &wg)
		i += 1
	}

	go func() {
		defer close(inputChan)
		for _, num := range nums {
			inputChan <- num
		}
	}()

	for _, opchn := range opList {
		for op := range opchn {
			fmt.Println(op)
		}

	}
	wg.Wait() //
	//why wait is here bcz we have added wait.Done() once we returns the output, so go routine should wait to consume the outputalso
	// else it will be called prematurly closing of op
}

func readNprint(inputChan <-chan int, wg *sync.WaitGroup) <-chan int {
	opChan := make(chan int)
	go func() {
		defer close(opChan)
		defer wg.Done()
		for input := range inputChan {
			fmt.Println(input, reflect.TypeOf(input))
			opChan <- input * input
		}
	}()
	return opChan
}
