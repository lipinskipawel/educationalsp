package main

import (
	"bufio"
	"log"
	"os"

	"github.com/lipinskipawel/educationalsp/rpc"
)

func main() {
	logger := getLogger("/home/pawel/project/lipinskipawel/educationalsp/log.txt")
	logger.Println("Hey, I started")

	scanner := bufio.NewScanner(os.Stdin)
	scanner.Split(rpc.Split)

	for scanner.Scan() {
		message := scanner.Text()
		handleMessage(logger, message)
	}
}

func handleMessage(logger *log.Logger, message any) {
	logger.Println(message)
}

func getLogger(filename string) *log.Logger {
	logfile, err := os.OpenFile(filename, os.O_CREATE|os.O_TRUNC|os.O_WRONLY, 0666)
	if err != nil {
		panic("hey, you didn't give me a good file")
	}

	return log.New(logfile, "[educationalsp]", log.Ldate|log.Ltime|log.Lshortfile)
}
