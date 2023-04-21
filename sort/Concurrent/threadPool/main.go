package main

import (
	"fmt"
	"runtime"
	"time"
)

//领导 — > 工人 ---- > 任务

type Job interface {
	Do()
}

type Worker struct {
	JobQueue chan Job
	Quit     chan bool
}

func NewWork() Worker {
	return Worker{
		JobQueue: make(chan Job),
		Quit:     make(chan bool),
	}
}

func (w Worker) Run(wq chan chan Job) {
	go func() {
		for {
			wq <- w.JobQueue
			select {
			case job := <-w.JobQueue:
				job.Do()
			case <-w.Quit:
				return
			}
		}
	}()
}

type WorkerPoll struct {
	workerlen   int
	JobQueue    chan Job
	WorkerQueue chan chan Job
}

func NewWorkerPool(workerlen int) *WorkerPoll {
	return &WorkerPoll{
		workerlen:   workerlen,
		JobQueue:    make(chan Job),
		WorkerQueue: make(chan chan Job, workerlen),
	}
}

func (wp *WorkerPoll) Run() {
	fmt.Println("初始化worker")
	for i := 0; i < wp.workerlen; i++ {
		worker := NewWork()
		worker.Run(wp.WorkerQueue)
	}

	go func() {
		for {
			select {
			case job := <-wp.JobQueue:
				worker := <-wp.WorkerQueue
				worker <- job
			}
		}
	}()
}

type Dosomething struct {
	Num int
}

func (d *Dosomething) Do() {
	fmt.Println("开启线程数", d.Num)
	time.Sleep(1 * 1 * time.Second)
}

func main() {
	num := 100 * 100 * 20
	p := NewWorkerPool(num)
	p.Run()

	datanum := 100 * 100
	go func() {
		for i := 1; i <= datanum; i++ {
			sc := &Dosomething{Num: i}
			p.JobQueue <- sc
		}
	}()

	for {
		fmt.Println("runtime.NumGoroutine() :", runtime.NumGoroutine())
		time.Sleep(2 * time.Second)
	}
}
