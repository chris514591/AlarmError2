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
}

func main() {
	var count int
	flag.IntVar(&count, "count", 0, "number of times to sound the alarm")
	flag.Parse()

	// load config from file
	configFile, err := os.Open("alarmConfig.json")
	if err != nil {
		log.Fatalf("error opening config file: %v", err)
	}
	defer configFile.Close()

	var config Config
	err = json.NewDecoder(configFile).Decode(&config)
	if err != nil {
		log.Fatalf("error decoding config file: %v", err)
	}

	// choose alarm level based on count
	var level string
	if count >= config.High {
		level = "High"
	} else if count >= config.Medium {
		level = "Medium"
	} else {
		level = "Low"
	}

	// log alarm level and count
	log.Printf("Alarm level: %s, Count: %d", level, count)

	// sound the alarm
	for i := 1; i <= count; i++ {
		time.Sleep(100 * time.Millisecond)
		fmt.Println("ALARM!", i)
	}
}
