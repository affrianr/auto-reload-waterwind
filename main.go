package main

import (
	"bytes"
	"encoding/json"
	"log"
	"math/rand"
	"net/http"
	"time"
)

type Data struct {
	Water int
	Wind  int
}

func main() {

	for {
		mainFunc()
		time.Sleep(15 * time.Second)
	}
}

func mainFunc() {

	water := rand.Intn(20) + 1
	wind := rand.Intn(20) + 1

	data := Data{
		Water: water,
		Wind:  wind,
	}

	jsonData, err := json.Marshal(data)

	if err != nil {
		log.Printf("Error marshaling data: %v", err)
		return
	}

	url := "http://localhost:8080/waterwind/1"

	req, err := http.NewRequest(http.MethodPut, url, bytes.NewBuffer(jsonData))

	if err != nil {
		log.Printf("error: %v", err)
	}

	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)

	if err != nil {
		log.Printf("error: %v", err)
	}

	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		log.Printf("Error updating data: %s", resp.Status)
	}

	waterStatus := getWaterStatus(water)
	windStatus := getWindStatus(wind)

	log.Printf("Data updated successfully")
	log.Printf("water: %d, status: %s", water, waterStatus)
	log.Printf("wind: %d, status: %s", wind, windStatus)

}

func getWaterStatus(water int) string {
	switch {
	case water < 5:
		return "aman"
	case water >= 6 && water <= 8:
		return "siaga"
	default:
		return "bahaya"
	}
}

func getWindStatus(wind int) string {
	switch {
	case wind < 6:
		return "aman"
	case wind >= 7 && wind <= 15:
		return "siaga"
	default:
		return "bahaya"
	}
}
