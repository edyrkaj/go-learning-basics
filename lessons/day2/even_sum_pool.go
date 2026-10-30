package main

import (
	"context"
	"fmt"
	"sync"
	"sync/atomic"
)

// Day 2 lesson: sum even numbers with a worker pool.
//
// Day 1 did this in one loop. This file does the same job with several
// goroutines, then checks the answer against the closed formula.
//
// Concepts used together:
//   - a small struct with a constructor and a method
//   - goroutines bounded by a worker count
//   - channels for jobs and partial sums
//   - select so cancellation can interrupt a send or receive
//   - sync.WaitGroup to know when every worker has returned
//   - sync.Once so startup work runs one time no matter which worker arrives first
//   - atomic.Int64 for a counter that many goroutines update

// span is one inclusive slice of the 1..limit range.
type span struct {
	start int
	end   int
}

// Result is what a finished SumEven call reports.
type Result struct {
	Sum           int
	Jobs          int64
	FirstWorkerID int
}

// Pool runs a fixed number of workers. chunk is how many integers each job covers.
type Pool struct {
	workers int
	chunk   int
}

func NewPool(workers, chunk int) (*Pool, error) {
	if workers < 1 {
		return nil, fmt.Errorf("workers must be >= 1, got %d", workers)
	}
	if chunk < 1 {
		return nil, fmt.Errorf("chunk must be >= 1, got %d", chunk)
	}
	return &Pool{workers: workers, chunk: chunk}, nil
}

// SumEven returns the sum of even integers in 1..limit.
// A cancelled context returns the partial sum and ctx.Err().
func (p *Pool) SumEven(ctx context.Context, limit int) (Result, error) {
	if limit < 0 {
		return Result{}, fmt.Errorf("limit must be >= 0, got %d", limit)
	}
	if err := ctx.Err(); err != nil {
		return Result{}, err
	}
	if limit < 2 {
		return Result{}, nil
	}

	jobs := make(chan span)
	partials := make(chan int)

	var (
		wg          sync.WaitGroup
		startOnce   sync.Once
		firstWorker atomic.Int64
		jobsDone    atomic.Int64
	)

	wg.Add(p.workers)
	for id := 1; id <= p.workers; id++ {
		go func(id int) {
			defer wg.Done()

			// Same idea as once.go: many goroutines call Do, the function runs once.
			// The once.Do call synchronizes with later reads of firstWorker.
			startOnce.Do(func() {
				firstWorker.Store(int64(id))
			})

			for {
				var job span
				select {
				case <-ctx.Done():
					return
				case got, ok := <-jobs:
					if !ok {
						return
					}
					job = got
				}

				partial := sumEvenInRange(job.start, job.end)
				jobsDone.Add(1)

				select {
				case <-ctx.Done():
					return
				case partials <- partial:
				}
			}
		}(id)
	}

	// Close partials only after every worker has left, so range in the
	// collector cannot miss a send or block forever on an open channel.
	go func() {
		wg.Wait()
		close(partials)
	}()

	go func() {
		defer close(jobs)
		for start := 1; start <= limit; start += p.chunk {
			end := start + p.chunk - 1
			if end > limit {
				end = limit
			}
			select {
			case <-ctx.Done():
				return
			case jobs <- span{start: start, end: end}:
			}
		}
	}()

	total := 0
	for partial := range partials {
		total += partial
	}

	result := Result{
		Sum:           total,
		Jobs:          jobsDone.Load(),
		FirstWorkerID: int(firstWorker.Load()),
	}
	// Workers may have stopped early. A closed partials channel is not proof
	// that every number was visited.
	if err := ctx.Err(); err != nil {
		return result, err
	}
	return result, nil
}

func sumEvenInRange(start, end int) int {
	sum := 0
	for i := start; i <= end; i++ {
		if i%2 == 0 {
			sum += i
		}
	}
	return sum
}

// evenSumFormula is 2+4+...+2k = k*(k+1), where k = limit/2.
// For limit 10, k is 5 and the sum is 30, the same answer as day 1.
func evenSumFormula(limit int) int {
	if limit < 2 {
		return 0
	}
	k := limit / 2
	return k * (k + 1)
}

func main() {
	pool, err := NewPool(4, 250)
	if err != nil {
		fmt.Println("pool:", err)
		return
	}

	// Same input as lessons/day1/sum_even_numbers.go.
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	day1, err := pool.SumEven(ctx, 10)
	if err != nil {
		fmt.Println("day1 check:", err)
		return
	}
	fmt.Printf("limit 10: sum=%d formula=%d jobs=%d firstWorker=%d\n",
		day1.Sum, evenSumFormula(10), day1.Jobs, day1.FirstWorkerID)
	if day1.Sum != evenSumFormula(10) {
		fmt.Println("day1 check failed")
		return
	}

	larger, err := pool.SumEven(ctx, 1000)
	if err != nil {
		fmt.Println("limit 1000:", err)
		return
	}
	fmt.Printf("limit 1000: sum=%d formula=%d jobs=%d firstWorker=%d\n",
		larger.Sum, evenSumFormula(1000), larger.Jobs, larger.FirstWorkerID)
	if larger.Sum != evenSumFormula(1000) {
		fmt.Println("limit 1000 check failed")
		return
	}

	// Cancel before any work. SumEven should return ctx.Err() and a partial sum.
	cancelled, stop := context.WithCancel(context.Background())
	stop()
	partial, err := pool.SumEven(cancelled, 1000)
	fmt.Printf("cancelled: partial=%d jobs=%d err=%v\n", partial.Sum, partial.Jobs, err)
}
