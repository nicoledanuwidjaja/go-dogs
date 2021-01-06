package main

import (
	"fmt"
	"time"
)

const (
	logInfo    = "INFO"
	logWarning = "WARNING"
	logError   = "ERROR"
)

type logEntry struct {
	time     time.Time
	severity string
	message  string
}

var logCh = make(chan logEntry, 50)
var doneCh = make(chan struct{}) // empty struct requires no memory allocation

func main() {
	go logger()
	//	defer func() {
	//		close(logCh)
	//	}()
	logCh <- logEntry{time.Now(), logInfo, "Starting up app"}
	time.Sleep(1000 * time.Millisecond)
	logCh <- logEntry{time.Now(), logInfo, "Shutting down app"}
	doneCh <- struct{}{}
}

func logger() {
	//	for entry := range logCh {
	//		fmt.Printf("%v - [%v]%v\n", entry.time.Format("2006-01-02T15:04:05"), entry.severity, entry.message)
	//	}
	for {
		select {
		case entry := <-logCh:
			fmt.Printf("%v - [%v]%v\n", entry.time.Format("2006-01-02T15:04:05"), entry.severity, entry.message)
		case <-doneCh:
			break
		default:
		}
	}
}
