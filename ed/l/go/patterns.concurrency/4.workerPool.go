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

	jobsCount := 20
	workersCount := runtime.NumCPU()

	jobs := jobsDispatcher(ctx, jobsCount, workersCount)
	wg := workerPool(ctx, jobs, jobsCount, workersCount)

	go func() {
		signals := make(chan os.Signal, 1)
		signal.Notify(signals, syscall.SIGINT, syscall.SIGTERM)
		<-signals
		cancel()
		fmt.Print("\n Exit.")
		return
	}()

	wg.Wait()
	fmt.Print("\n All jobs done.")
}

func jobsDispatcher(ctx context.Context, jobsCount int, consumersCount int) <-chan Job {
	jobs := make(chan Job, consumersCount)

	go func() {
		for v := 0; v < jobsCount; v++ {
			err := produceJob(ctx, jobs, v)
			if errors.Is(err, context.Canceled) {
				fmt.Printf("\n %s for dispatched, err: %v.", CtxDoneMsg, ctx.Err())
				return
			}
		}
		close(jobs)
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

func workerPool(ctx context.Context, jobs <-chan Job, jobsCount int, workersCount int) *sync.WaitGroup {
	wg := &sync.WaitGroup{}
	wg.Add(jobsCount)

	for i := 0; i < workersCount; i++ {
		go worker(ctx, wg, jobs)
	}

	return wg
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

func worker(ctx context.Context, wg *sync.WaitGroup, jobs <-chan Job) {
	id := RandId()
	fmt.Printf("\n %s %s.", WorkerStartMsg, id)

	for {
		select {
		case j, ok := <-jobs:
			if ok {
				execJob(ctx, id, j)
				wg.Done()
			} else {
				fmt.Printf("\n %s %s, jobs chan is empty.", WorkerStopMsg, id)
				return
			}
		case <-ctx.Done():
			fmt.Printf("\n %s %s, err: %v.", WorkerHaltMsg, id, ctx.Err())
			return
		default:
			time.Sleep(Tick)
			fmt.Print("r")
		}
	}
}

func execJob(ctx context.Context, workerId string, j Job) {
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
			return
		default:
			time.Sleep(Tick)
			print(".")
		}
	}
}
