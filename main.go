package main

import (
	"fmt"
	"github.com/pedrolopesme/citta-server/core"
	"log"
	"net/http"
	"time"
)

func main() {

	broker := core.NewServer()

	go func() {
		for {
			time.Sleep(time.Millisecond * 500)
			eventString := fmt.Sprintf("the time is %v", time.Now())
			log.Println("Receiving event REALLY?")
			broker.Notifier <- []byte(eventString)
		}
	}()

	log.Fatal("HTTP server error: ", http.ListenAndServe(":3030", broker))

}
