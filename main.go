package main

import (
	"fmt"
	"log"
	"os"
)

type Data struct {
	Low    int64
	Medium int64
	High   int64
}

func main() {
	logFile1, err1 := os.OpenFile("alarmErrorLogs.txt", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)
	if err1 != nil {
		os.Exit(1)
	}
	log.SetOutput(logFile1)
	log.Println("Alarm error program initiated.")

	jsonContent, err2 := os.ReadFile("./alarmConfig.json")
	if err2 != nil {
		log.Fatal(err2)
		log.SetOutput(logFile1)
		log.Println("Could not open alarmConfig.json")
	}
	fmt.Println(jsonContent)
}
