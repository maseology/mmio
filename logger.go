package mmio

// from: https://stackoverflow.com/questions/18361750/correct-approach-to-global-logging-in-golang

import (
	"log"
	"os"
	"sync"
)

type logger struct {
	filename string
	*log.Logger
}

var mmlog *logger
var once sync.Once

// GetInstance start log file outputting to ./mm.log
func GetInstance(fnam string) *logger {
	once.Do(func() {
		mmlog = createLogger(fnam)
	})
	return mmlog
}

func createLogger(fname string) *logger {
	file, _ := os.OpenFile(fname, os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0777)

	return &logger{
		filename: fname,
		Logger:   log.New(file, "", log.Lshortfile),
	}
}

////////////////////////////////////
//// SAMPLE
////////////////////////////////////
// package main

// import (
//     "mmio"
//     "fmt"
//     "net/http"
// )

// func main() {
//     logger := mmio.GetInstance()
//     logger.Println("Starting")

//     http.HandleFunc("/", sroot)
//     http.ListenAndServe(":8080", nil)
// }

// func sroot(w http.ResponseWriter, r *http.Request) {
//     logger := mmio.GetInstance()

//     fmt.Fprintf(w, "welcome")
//     logger.Println("Starting")
// }
