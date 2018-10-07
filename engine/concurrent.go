package engine

import (
	"github.com/GerryLon/go-crawler/config"
	"github.com/GerryLon/go-crawler/filter"
	"log"
)

type ConcurrentEngine struct {
	Scheduler   Scheduler
	WorkerCount int
}

type Scheduler interface {
	Submit(Request)
	ConfigWorkerChan(chan Request)
}

func (e *ConcurrentEngine) Run(seeds ...Request) {
	in := make(chan Request)
	out := make(chan ParseResult)
	e.Scheduler.ConfigWorkerChan(in)

	for i := 0; i < e.WorkerCount; i++ {
		e.createWorker(in, out)
	}

	// 将seeds送给scheduler
	for _, request := range seeds {
		if isDuplicate(request.Url) {
			log.Printf("#%d: %s is duplicate", dedupFilter.Len(), request.Url)
			continue
		}
		e.Scheduler.Submit(request)
	}

	// 为item计数
	count := 0

	for {

		// worker会生成ParseResult， 在这里收
		r := <-out

		for _, item := range r.Items {
			count++
			log.Printf("Got #%d item: %v", count, item)
		}

		// 将worker生成的ParseResult中的Requests送给scheduler
		for _, request := range r.Requests {
			if isDuplicate(request.Url) {
				log.Printf("#%d: %s is duplicate", dedupFilter.Len(), request.Url)
				continue
			}
			e.Scheduler.Submit(request)
		}
	}
}

func (e *ConcurrentEngine) createWorker(in chan Request, out chan ParseResult) {
	go func() {
		for {
			request := <-in
			result, err := worker(request)

			if err != nil {
				continue
			}

			out <- result
		}
	}()
}

var dedupFilter = filter.MemoryDedupFilter{}

func isDuplicate(url string) bool {
	if config.WillDeDup {
		if dedupFilter.Has(url) {
			return true
		} else {
			dedupFilter.Set(url)
			return false
		}
	}
	return false
}
