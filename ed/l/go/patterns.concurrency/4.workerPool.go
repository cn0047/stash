package main

import (
	"context"
	"errors"
	"fmt"
	"math/rand"
	"os"
	"os/signal"
	"runtime"
	"sync"
	"syscall"
	"time"
)

const (
	CtxDoneMsg     = "\033[35mctx done\033[0m"
	WorkerStartMsg = "\033[32mstpn worker\033[0m"
	WorkerStopMsg  = "\033[36mstpn worker\033[0m"
	WorkerHaltMsg  = "\033[31mhalt worker\033[0m"
	Tick           = 15000 * time.Microsecond
	JobTick        = 750000 * time.Microsecond
)

type Job struct {
	Val int
}

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	go func() {
		signals := make(chan os.Signal, 1)
		signal.Notify(signals, syscall.SIGINT, syscall.SIGTERM)
		<-signals
		cancel()
		fmt.Print("\n Exit.")
	}()

	errorMessages := run(ctx)
	fmt.Printf("\n All jobs done, errors count: %v", len(errorMessages))
	for _, err := range errorMessages {
		fmt.Printf("\n\t error: %v", err)
	}
}

func run(ctx context.Context) []error {
	jobsCount := 13
	workersCount := runtime.NumCPU()

	jobs := jobsDispatcher(ctx, jobsCount, workersCount)
	wg, errs := workerPool(ctx, jobs, jobsCount, workersCount)
	er := readErrors(errs, jobsCount)

	wg.Wait()
	close(errs)

	return er.getErrors()
}

type ErrorsReader struct {
	captured []error
	sync.WaitGroup
}

func (e *ErrorsReader) getErrors() []error {
	e.Wait()
	return e.captured
}

func readErrors(errs <-chan error, size int) *ErrorsReader {
	er := &ErrorsReader{captured: make([]error, 0, size)}
	er.Add(1)

	go func() {
		defer func() {
			er.Done()
		}()

		for {
			select {
			case err, ok := <-errs:
				if !ok {
					return
				}
				er.captured = append(er.captured, err)
			default:
				time.Sleep(Tick)
			}
		}
	}()

	return er
}

func jobsDispatcher(ctx context.Context, jobsCount int, consumersCount int) <-chan Job {
	jobs := make(chan Job, consumersCount)

	go func() {
		defer func() {
			close(jobs)
		}()

		for v := 0; v < jobsCount; v++ {
			err := produceJob(ctx, jobs, v)
			if errors.Is(err, context.Canceled) {
				fmt.Printf("\n %s for dispatched, err: %v.", CtxDoneMsg, ctx.Err())
				return
			}
		}
	}()

	return jobs
}

func produceJob(ctx context.Context, jobs chan<- Job, v int) error {
	j := Job{Val: v}
	for {
		select {
		case jobs <- j:
			fmt.Printf("\n dispatched job %v.", j.Val)
			return nil
		case <-ctx.Done():
			return ctx.Err()
		default:
			time.Sleep(Tick)
			fmt.Print("w")
		}
	}
}

func workerPool(
	ctx context.Context, jobs <-chan Job, jobsCount int, workersCount int,
) (*sync.WaitGroup, chan error) {
	errs := make(chan error, workersCount)
	wg := &sync.WaitGroup{}
	wg.Add(jobsCount)

	for i := 0; i < workersCount; i++ {
		go worker(ctx, wg, jobs, errs)
	}

	return wg, errs
}

func RandId() string {
	letterRunes := []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")
	n := 6
	b := make([]rune, n)
	for i := range b {
		b[i] = letterRunes[rand.Intn(len(letterRunes))]
	}
	return string(b)
}

func worker(ctx context.Context, wg *sync.WaitGroup, jobs <-chan Job, errs chan<- error) {
	id := RandId()
	fmt.Printf("\n %s %s.", WorkerStartMsg, id)

	for {
		select {
		case j, ok := <-jobs:
			if !ok {
				fmt.Printf("\n %s %s, jobs chan is empty or closed.", WorkerStopMsg, id)
				return
			}
			err := execJob(ctx, id, j)
			if err != nil {
				er := sendError(ctx, errs, err)
				if errors.Is(er, context.Canceled) {
					return
				}
			}
			wg.Done()
		case <-ctx.Done():
			fmt.Printf("\n %s %s, err: %v.", WorkerHaltMsg, id, ctx.Err())
			return
		default:
			time.Sleep(Tick)
			fmt.Print("r")
		}
	}
}

func sendError(ctx context.Context, errs chan<- error, err error) error {
	for {
		select {
		case errs <- err:
			return nil
		case <-ctx.Done():
			return ctx.Err()
		default:
			time.Sleep(Tick)
		}
	}
}

func execJob(ctx context.Context, workerId string, j Job) error {
	fmt.Printf("\n\t rcvd job %v worker %s.", j.Val, workerId)

	select {
	case <-ctx.Done():
		fmt.Printf("\n\t %s for job executor for worker %s, err: %v.", CtxDoneMsg, workerId, ctx.Err())
	default:
		fmt.Print("*")
	}

	fmt.Printf("\n\t strt job %v worker %s.", j.Val, workerId)

	done := time.After(JobTick)
	for {
		select {
		case <-done:
			fmt.Printf("\n\t FNSH job %v worker %s.", j.Val, workerId)
			return fmt.Errorf("failed to exec job: %v, error: exec error", j.Val)
		default:
			time.Sleep(Tick)
			print(".")
		}
	}
}
