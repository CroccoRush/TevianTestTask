package watchdog

import (
	"fmt"
	"io"
	"log"
)

// Dog wrapper for tracking the start and end of the goroutine execution
func Dog(globsig chan bool, innerch chan bool, object io.Closer, objectName string) {
	log.Println(fmt.Sprintf("%s was started...", objectName))
watchdog:
	for {
		select {
		case <-globsig:
			err := object.Close()
			if err != nil {
				log.Fatal(err)
			}
			break watchdog
		case <-innerch:
			break watchdog
		}
	}
	log.Println(fmt.Sprintf("%s was stopped...", objectName))
}
