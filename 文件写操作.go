package main

import (
	"io"
	"log"
	"os"
)

func main() {
	logfile, err := os.OpenFile("/tmp/test", os.O_CREATE|os.O_RDWR|os.O_APPEND, 0644)
	logfile, err = os.OpenFile("/tmp/test", os.O_CREATE|os.O_TRUNC|os.O_RDWR, 0644)

	defer logfile.Close()
	_, err = io.WriteString(logfile, "hello world~")

	if err != nil {
		log.Fatal("there is a issue while open the files!")
	}
}
