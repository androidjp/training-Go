package main

import (
	"encoding/json"
	"log"
	"net/http"
)

type openWeatherMap struct{}

func (w openWeatherMap) temperature(city string) (float64, error) {
	resp, err := http.Get("https://samples.openweathermap.org/data/2.5/weather?q=" + city + "&appid=439d4b804bc8187953eb36d2a8c26a02")
	if err != nil {
		return 0, err
	}
	defer resp.Body.Close()
	// 解析 json 里头的内容，并存储到这个 d 结构对象上
	var d struct {
		Main struct {
			Kelvin float64 `json:"temp"`
		} `json:"main"`
	}
	if err := json.NewDecoder(resp.Body).Decode(&d); err != nil {
		return 0, err
	}
	log.Printf("openWeatherMap: %s: %.2f", city, d.Main.Kelvin)
	return d.Main.Kelvin, nil
}

func main() {
	weatherMap := openWeatherMap{}
	weatherMap.temperature("Zhuhai,zh")
	weatherMap.temperature("Shanghai,zh")
}
