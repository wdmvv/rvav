package handlers

// for reporting jobs that were added
import (
	"bytes"
	"encoding/json"
	"io"
	"net/http"
	"sync"
)

var Jobs *JobsInfo
var once sync.Once

func JobsHandler(w http.ResponseWriter, r *http.Request) {
	WriteStruct(Jobs, w, r)
}

// made for /addexpr specifically
func JobsLogger(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		rec := &StatusRecorder{ResponseWriter: w}
		rec.Code = http.StatusOK //assuming that default is 200
		body, err := io.ReadAll(r.Body)
		if err != nil {
			http.Error(w, "failed to read request body", http.StatusInternalServerError)
			return
		}
		r.Body = io.NopCloser(bytes.NewBuffer(body))

		data := AddExprReqIn{}
		err = json.Unmarshal(body, &data)

		if err == nil {
			addJob(data)
		}

		next.ServeHTTP(rec, r)

		if rec.Code == http.StatusOK {
			moveCompleted(data)
		} else {
			moveFailed(data)
		}
	})
}

func addJob(data AddExprReqIn) {
	Jobs.Lock.Lock()
	defer Jobs.Lock.Unlock()
	Jobs.Running[data.Id] = data.Expr
}

func moveFailed(data AddExprReqIn) {
	Jobs.Lock.Lock()
	defer Jobs.Lock.Unlock()
	delete(Jobs.Running, data.Id)
	Jobs.Failed[data.Id] = data.Expr
}

func moveCompleted(data AddExprReqIn) {
	Jobs.Lock.Lock()
	defer Jobs.Lock.Unlock()
	delete(Jobs.Running, data.Id)
	Jobs.Completed[data.Id] = data.Expr
}

func StartJobs() {
	once.Do(func() {
		Jobs = &JobsInfo{sync.Mutex{}, make(map[int]string), make(map[int]string), make(map[int]string)}
	})
}
