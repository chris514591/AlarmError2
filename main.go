package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"log"
	"os"
	"time"
)

type Config struct {
	Low    int `json:"Low"`
	Medium int `json:"Medium"`
	High   int `json:"High"`
	Panic  int `json:"NO NOT THE GERMANS AGAIN"`
}

func main() {
	var count int
	flag.IntVar(&count, "count", 0, "number of times to sound the alarm")
	flag.Parse()

	// Open a file for logging errors
	errorLogFile, err := os.OpenFile("alarmErrorLogs.log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		log.Fatalf("error opening error log file: %v", err)
	}
	defer errorLogFile.Close()

	log.SetOutput(errorLogFile)

	// Create a logger that writes to the error log file
	errorLogger := log.New(errorLogFile, "", log.LstdFlags)

	// Load config from file
	configFile, err := os.Open("alarmConfig.json")
	if err != nil {
		errorLogger.Printf("error opening config file: %v", err)
		log.Fatalf("error opening config file: %v", err)
	}
	defer configFile.Close()

	var config Config
	err = json.NewDecoder(configFile).Decode(&config)
	if err != nil {
		errorLogger.Printf("error decoding config file: %v", err)
		log.Fatalf("error decoding config file: %v", err)
	}

	// Choose alarm level based on count
	var level string
	if count >= config.Panic {
		level = "Panic"
	} else if count >= config.High {
		level = "High"
	} else if count >= config.Medium {
		level = "Medium"
	} else {
		level = "Low"
	}

	// Log alarm level and count
	log.Printf("Alarm level: %s, Count: %d", level, count)

	// Sound the alarm
	for i := 1; i <= count; i++ {
		time.Sleep(10 * time.Millisecond)
		fmt.Println("ALARM!", i)
	}
}
