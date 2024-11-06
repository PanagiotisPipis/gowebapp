package watcher

import (
	"sync"
	"encoding/json"

	"github.com/google/uuid"
)

type Watcher struct {
	id          string         // Watcher ID.
	inCh        chan string    // Input channel.
	outCh       chan []byte  // Updates to counter will notify this channel.
	counter     *Counter       // The counter.
	counterLock *sync.RWMutex  // Lock for counter.
	quitChannel chan struct{}  // Quit.
	running     sync.WaitGroup // Run, Amy, Run!
}

func New() *Watcher {
	w := Watcher{}
	w.id = uuid.NewString()
	w.inCh = make(chan string, 1)
	w.outCh = make(chan []byte, 1)
	w.counter = &Counter{Iteration: 0}
	w.counterLock = &sync.RWMutex{}
	w.quitChannel = make(chan struct{})
	w.running = sync.WaitGroup{}
	return &w
}

// Start watcher in another Go routine, Stop() must be called at the end.
func (w *Watcher) Start() error {
	w.running.Add(1)
	go func(wg *sync.WaitGroup) {
		defer wg.Done()
		for {
			select {
			case str := <-w.inCh:
				w.counter.Iteration += 1
				w.counter.HexString = str
				data, _ := json.Marshal(w.counter)
				select {
				case w.outCh <- data:
				case <-w.quitChannel:
					return
				}
			case <-w.quitChannel:
				return
			}
		}
	}(&w.running)

	return nil
}

func (w *Watcher) Stop() {
	w.counterLock.Lock()
	defer w.counterLock.Unlock()

	close(w.quitChannel)
	w.running.Wait()
}

func (w *Watcher) GetWatcherId() string { return w.id }

func (w *Watcher) Send(str string) { w.inCh <- str }

func (w *Watcher) Recv() <-chan []byte { return w.outCh }

func (w *Watcher) ResetCounter() {
	w.counterLock.Lock()
	defer w.counterLock.Unlock()

	w.counter.Iteration = 0
	data, _ := json.Marshal(w.counter)
	select {
	case w.outCh <- data:
	case <-w.quitChannel:
		return
	}
}
