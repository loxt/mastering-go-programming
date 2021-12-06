package hydrachat

import (
	"fmt"
	"io"
	"sync"

	"github.com/loxt/mastering-go-programming/hydra/hlogger"
)

type room struct {
	name    string
	Msgch   chan string
	clients map[chan<- string]struct{}
	Quit    chan struct{}
	*sync.RWMutex
}

func CreateRoom(name string) *room {
	r := &room{
		name:    name,
		Msgch:   make(chan string),
		RWMutex: new(sync.RWMutex),
		clients: make(map[chan<- string]struct{}),
		Quit:    make(chan struct{}),
	}
	r.Run()
	return r
}

// todo 12:00

func (r *room) AddClient(c io.ReadWriteCloser) {
	r.Lock()
	wc, done := StartClient(r.Msgch, c, r.Quit)
	r.clients[wc] = struct{}{}
	r.Unlock()

	// remove client when done is signalled
	go func() {
		<-done
		r.RemoveClient(wc)
	}()
}

func (r *room) Run() {
	logger := hlogger.GetInstance()
	logger.Println("Starting chat  room", r.name)
	go func() {
		for msg := range r.Msgch {
			r.broadcastMsg(msg)
		}
	}()
}

func (r *room) broadcastMsg(msg string) {
	r.RLock()
	defer r.RUnlock()
	fmt.Println("Received message: ", msg)
	for wc := range r.clients {
		go func(wc chan<- string) {
			wc <- msg
		}(wc)
	}
}
