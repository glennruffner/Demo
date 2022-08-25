package workers

import (
	"fmt"
	"math/rand"
	"strconv"
	"sync"
	"time"

	"golang.org/x/sync/errgroup"
)

// This is an example of channel approach to workers, includes error hnndling
func ConvToString(inputs []int, workers int) ([]string, error) {
	fmt.Println("")
	fmt.Println("ConvToString:")
	var output []string

	inChan, outChan := make(chan int), make(chan string)

	go func() {
		defer close(inChan)
		for _, v := range inputs {
			inChan <- v
		}
	}()

	var g errgroup.Group

	for i := 0; i < workers; i++ {
		func(id int) {
			g.Go(func() error {
				for v := range inChan {
					s := convertToString(v)
					outChan <- s
				}
				//		return errors.New("this is a new error")
				return nil
			})
		}(i + 1)
	}

	var drain sync.WaitGroup

	drain.Add(1)

	go func() {
		defer drain.Done()
		for s := range outChan {
			output = append(output, s)
		}
	}()

	err := g.Wait()
	close(outChan)
	if err != nil {
		return nil, err
	}
	drain.Wait()

	return output, nil

}

func convertToString(in int) string {

	// fmt.Printf("worker: %v | CONVERTING: %v\n", workerId, in)

	workDuration := time.Duration(rand.Intn(100)+50) * time.Millisecond // Wait between 0.5 and 1.5 seconds
	time.Sleep(workDuration)

	// fmt.Printf("worker: %v | DONE CONVERTING: %v | TOOK: %v\n", workerId, in, workDuration)

	return strconv.Itoa(in)
}

func SimpleWorkerConvertToString(in []int, workerCnt int) []string {
	fmt.Println("")
	fmt.Println("Processing SimpleWorkerConvertToString:")
	var out []string
	var workBuffer = make(chan bool, workerCnt)
	var outChan = make(chan string)

	var drain sync.WaitGroup

	if len(in) < 1 {
		return []string{}
	}

	drain.Add(1)

	// launch output reader first so it doesn't deadlock
	go func() {
		defer drain.Done()
		for s := range outChan {
			out = append(out, s)
		}
	}()

	for _, v := range in {
		workBuffer <- true
		go func(id int) {
			defer func() { <-workBuffer }()
			outChan <- convertToString(id)
		}(v)
	}
	// clear buffer
	for i := 0; i < workerCnt; i++ {
		workBuffer <- true
	}

	close(outChan)

	drain.Wait()

	return out
}
