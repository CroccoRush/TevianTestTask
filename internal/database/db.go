package database

import (
	"TevianTestTask/pkg/watchdog"
	"sync"
)

func Run(globsig chan bool, wg *sync.WaitGroup, onStartWg *sync.WaitGroup) {

	defer wg.Done()
	innerch := make(chan bool)

	onStartWg.Done()
	watchdog.Dog(globsig, innerch, DB, "database")
}
