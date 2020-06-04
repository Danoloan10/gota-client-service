package main

import (
	"os"
    "os/signal"
    "syscall"
	"sync"

	"github.com/danoloan10/gota-client-service/ui"
)

const (
	CNNHostname = "cnn.com"
)
func main() {
	sigs := make(chan os.Signal, 1)
	signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM)

	serviceWg := &sync.WaitGroup{}
	srv := ui.StartUIServer(serviceWg)
	// parar srv
	<-sigs
	srv.Close()
	serviceWg.Done()
}
