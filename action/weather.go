package actions

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"
	"time"

	"github.com/tacheraSasi/ellie/styles"
)

type WeatherData struct {
	CurrentCondition []struct {
		FeelsLikeC    string `json:"FeelsLikeC"`
		TempC         string `json:"temp_C"`
		Humidity      string `json:"humidity"`
		WeatherDesc   []struct {
			Value string `json:"value"`
		} `json:"weatherDesc"`
		WindspeedKmph string `json:"windspeedKmph"`
		Visibility    string `json:"visibility"`
		Pressure      string `json:"pressure"`
		CloudCover    string `json:"cloudcover"`
		UVIndex       string `json:"uvIndex"`
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

// getWeatherIcon returns an emoji icon based on weather description
func getWeatherIcon(description string) string {
	desc := strings.ToLower(description)
	
	if strings.Contains(desc, "clear") || strings.Contains(desc, "sunny") {
		return "☀️"
	} else if strings.Contains(desc, "partly cloudy") {
		return "⛅"
	} else if strings.Contains(desc, "cloudy") || strings.Contains(desc, "overcast") {
		return "☁️"
	} else if strings.Contains(desc, "rain") || strings.Contains(desc, "drizzle") {
		return "🌧️"
	} else if strings.Contains(desc, "thunder") || strings.Contains(desc, "storm") {
		return "⛈️"
	} else if strings.Contains(desc, "snow") {
		return "❄️"
	} else if strings.Contains(desc, "mist") || strings.Contains(desc, "fog") {
		return "🌫️"
	} else if strings.Contains(desc, "wind") {
		return "💨"
	}
	return "🌤️"
}

func Weather() {
	// Display header with animation
	styles.GetHeaderStyle().Println("\n🌍 Weather Information")
	styles.GetInfoStyle().Println("━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━")
	styles.GetInfoStyle().Println("📡 Fetching current weather data...")
	
	// Create HTTP client with timeout
	client := &http.Client{
		Timeout: 15 * time.Second,
	}
	
	resp, err := client.Get("https://wttr.in/?format=j1")
	if err != nil {
		styles.GetErrorStyle().Println("\n❌ Failed to fetch weather data")
		styles.DimText.Println("   Error: Unable to connect to weather service")
		styles.DimText.Println("   Please check your internet connection and try again")
		return
	}
	defer resp.Body.Close()

	// Check response status
	if resp.StatusCode != http.StatusOK {
		styles.GetErrorStyle().Printf("\n❌ Weather service returned error: %s\n", resp.Status)
		return
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		styles.GetErrorStyle().Println("\n❌ Failed to read weather data")
		styles.DimText.Println("   Error: Unable to read response from weather service")
		return
	}

	var weather WeatherData
	err = json.Unmarshal(body, &weather)
	if err != nil {
		styles.GetErrorStyle().Println("\n❌ Failed to parse weather data")
		styles.DimText.Println("   Error: Invalid response format from weather service")
		return
	}

	if len(weather.CurrentCondition) > 0 && len(weather.NearestArea) > 0 {
		current := weather.CurrentCondition[0]
		area := weather.NearestArea[0]

		// Format location
		location := fmt.Sprintf("%s, %s, %s", 
			area.AreaName[0].Value, 
			area.Region[0].Value, 
			area.Country[0].Value)
		
		description := current.WeatherDesc[0].Value
		icon := getWeatherIcon(description)

		// Display weather information with improved formatting
		fmt.Println()
		styles.GetSuccessStyle().Printf("📍 Location: %s\n\n", location)
		
		// Main weather condition
		styles.Highlight.Printf("%s  %s\n\n", icon, description)
		
		// Temperature section
		styles.GetHeaderStyle().Println("🌡️  Temperature")
		fmt.Printf("   Current:   %s°C\n", current.TempC)
		fmt.Printf("   Feels Like: %s°C\n\n", current.FeelsLikeC)
		
		// Additional details
		styles.GetHeaderStyle().Println("💨  Conditions")
		fmt.Printf("   Wind Speed: %s km/h\n", current.WindspeedKmph)
		fmt.Printf("   Humidity:   %s%%\n", current.Humidity)
		
		// Optional fields (only show if available)
		if current.Visibility != "" {
			fmt.Printf("   Visibility: %s km\n", current.Visibility)
		}
		if current.Pressure != "" {
			fmt.Printf("   Pressure:   %s mb\n", current.Pressure)
		}
		if current.CloudCover != "" {
			fmt.Printf("   Cloud Cover: %s%%\n", current.CloudCover)
		}
		if current.UVIndex != "" {
			fmt.Printf("   UV Index:   %s\n", current.UVIndex)
		}
		
		fmt.Println()
		styles.GetInfoStyle().Println("━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━")
		styles.DimText.Println("Data provided by wttr.in")
		fmt.Println()
	} else {
		styles.GetErrorStyle().Println("\n❌ Could not retrieve weather information")
		styles.DimText.Println("   The weather service did not return complete data")
		styles.DimText.Println("   Please try again later")
	}
}
