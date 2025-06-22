package actions

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/tacheraSasi/ellie/styles"
)

type WeatherData struct {
	CurrentCondition []struct {
		FeelsLikeC  string `json:"FeelsLikeC"`
		TempC       string `json:"temp_C"`
		Humidity    string `json:"humidity"`
		WeatherDesc []struct {
			Value string `json:"value"`
		} `json:"weatherDesc"`
		WindspeedKmph string `json:"windspeedKmph"`
	} `json:"current_condition"`
	NearestArea []struct {
		AreaName []struct {
			Value string `json:"value"`
		} `json:"areaName"`
		Country []struct {
			Value string `json:"value"`
		} `json:"country"`
		Region []struct {
			Value string `json:"value"`
		} `json:"region"`
	} `json:"nearest_area"`
}

func Weather() {
	styles.InfoStyle.Println("Fetching weather information...")
	resp, err := http.Get("https://wttr.in/?format=j1")
	if err != nil {
		styles.ErrorStyle.Println("Failed to fetch weather data.")
		return
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	fmt.Println(&resp.Body)
	if err != nil {
		styles.ErrorStyle.Println("Failed to read weather data.")
		return
	}

	var weather WeatherData
	err = json.Unmarshal(body, &weather)
	if err != nil {
		styles.ErrorStyle.Println("Failed to parse weather data.")
		return
	}

	if len(weather.CurrentCondition) > 0 && len(weather.NearestArea) > 0 {
		current := weather.CurrentCondition[0]
		area := weather.NearestArea[0]

		location := fmt.Sprintf("%s, %s, %s", area.AreaName[0].Value, area.Region[0].Value, area.Country[0].Value)
		description := current.WeatherDesc[0].Value
		temp := fmt.Sprintf("%s°C (Feels like %s°C)", current.TempC, current.FeelsLikeC)
		humidity := fmt.Sprintf("%s%%", current.Humidity)
		wind := fmt.Sprintf("%s km/h", current.WindspeedKmph)

		styles.SuccessStyle.Println("Current Weather in " + location)
		fmt.Println("Description:", description)
		fmt.Println("Temperature:", temp)
		fmt.Println("Humidity:", humidity)
		fmt.Println("Wind Speed:", wind)
	} else {
		styles.ErrorStyle.Println("Could not retrieve weather information.")
	}
}
