package workers

import (
	"fmt"
	"math/rand"
	"strconv"
	"sync"
	"time"

	"golang.org/x/sync/errgroup"
)

func ConvToString(inputs []int, workers int) ([]string, error) {

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
					s := convertToString(v, id)
					outChan <- s
				}

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

func convertToString(in, workerId int) string {

	fmt.Printf("worker: %v | CONVERTING: %v\n", workerId, in)

	workDuration := time.Duration(rand.Intn(1000)+500) * time.Millisecond // Wait between 0.5 and 1.5 seconds
	time.Sleep(workDuration)

	fmt.Printf("worker: %v | DONE CONVERTING: %v | TOOK: %v\n", workerId, in, workDuration)

	return strconv.Itoa(in)
}
