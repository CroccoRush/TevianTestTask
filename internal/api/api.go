package api

import (
	"TevianTestTask/pkg/watchdog"
	"log"
	"sync"
)

func Run(globsig chan bool, wg *sync.WaitGroup, onStartWg *sync.WaitGroup) {

	defer wg.Done()
	innerch := make(chan bool)

	go watchdog.Dog(globsig, innerch, app, "server")
	onStartWg.Done()
	log.Println("http server started on :3000")
	if err := app.Listen(":3000"); err != nil {
		log.Fatal(err)
	}
	close(innerch)
}
