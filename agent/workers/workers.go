package workers

// for controlling workers

import (
	"agent/logging"
	"context"
	"os"
	"strconv"
	"sync"
	"time"

	"golang.org/x/sync/semaphore"
)

type WorkersInfo struct {
	workers int64              `json:"-"`
	limit   semaphore.Weighted `json:"-"`
	Current map[string]string  `json:"current"`
	lock    sync.Mutex         `json:"-"`
}

var Info WorkersInfo

func Set() {
	env := os.Getenv("MAX_WORKERS")
	workers, err := strconv.ParseInt(env, 10, 64)
	if err != nil {
		logging.ReportAction("did not find env MAX_WORKERS, setting default 10")
		workers = 10
	}
	Info = WorkersInfo{workers, *semaphore.NewWeighted(workers), make(map[string]string), sync.Mutex{}}
}

func (w *WorkersInfo) Task(expr string) {
	w.limit.Acquire(context.Background(), 1)
	start := time.Now().Format("01/02 - 03:04:05")
	w.lock.Lock()
	defer w.lock.Unlock()
	w.Current[expr] = start
}

func (w *WorkersInfo) Expire(expr string) {
	w.limit.Release(1)
	w.lock.Lock()
	defer w.lock.Unlock()
	delete(w.Current, expr)
}
