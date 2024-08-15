package main

import (
	"TevianTestTask/internal/api"
	"TevianTestTask/internal/database"
	"log"
	"os"
	"os/signal"
	"sync"
)

var (
	STAGE        string
	Gwg          sync.WaitGroup
	onStartWg    sync.WaitGroup
	globsig      chan os.Signal
	interruption chan bool
	//serverConf   conf.ServerConfig
)

func init() {
	STAGE = "development"
	Gwg.Add(30)
	if err := os.Setenv("SERV", ":3000"); err != nil {
		log.Fatal(err)
	}
	if err := os.Setenv("NAME_VER", "base"); err != nil {
		log.Fatal(err)
	}
	globsig = make(chan os.Signal, 1)
	interruption = make(chan bool)
	signal.Notify(globsig, os.Interrupt)
	//serverConf.Read("SERV")
}

func main() {
	//err := facecloud.Detect("./storage/uploads/Glasses.jpeg")
	//if err != nil {
	//	log.Fatal(err)
	//}
	//return
	log.Println("server starting...")
	onStartWg.Add(1)
	go api.Run(interruption, &Gwg, &onStartWg)
	onStartWg.Wait()
	onStartWg.Add(1)
	go database.Run(interruption, &Gwg, &onStartWg)
	onStartWg.Wait()

watchdog:
	for {
		select {
		case <-globsig:
			log.Println("...interrupting")
			close(globsig)
			close(interruption)
			onStartWg.Wait()
			Gwg.Wait()
			log.Println("server stopped...")
			break watchdog
		}
	}
}
