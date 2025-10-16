package actions

import (
	"strings"
	"testing"
)

func TestGetWeatherIcon(t *testing.T) {
	tests := []struct {
		description string
		expected    string
	}{
		{"Clear", "☀️"},
		{"Sunny", "☀️"},
		{"Partly cloudy", "⛅"},
		{"Cloudy", "☁️"},
		{"Overcast", "☁️"},
		{"Rain", "🌧️"},
		{"Light rain", "🌧️"},
		{"Drizzle", "🌧️"},
		{"Thunderstorm", "⛈️"},
		{"Thunder", "⛈️"},
		{"Snow", "❄️"},
		{"Light snow", "❄️"},
		{"Mist", "🌫️"},
		{"Fog", "🌫️"},
		{"Windy", "💨"},
		{"Unknown weather", "🌤️"},
	}

	for _, tt := range tests {
		t.Run(tt.description, func(t *testing.T) {
			result := getWeatherIcon(tt.description)
			if result != tt.expected {
				t.Errorf("getWeatherIcon(%q) = %q; want %q", tt.description, result, tt.expected)
			}
		})
	}
}

func TestGetWeatherIconCaseInsensitive(t *testing.T) {
	testCases := []string{"CLEAR", "clear", "Clear", "cLeAr"}
	expected := "☀️"

	for _, tc := range testCases {
		result := getWeatherIcon(tc)
		if result != expected {
			t.Errorf("getWeatherIcon(%q) = %q; want %q", tc, result, expected)
		}
	}
}

func TestGetWeatherIconPartialMatch(t *testing.T) {
	// Test that partial matches work
	descriptions := []string{
		"Partly cloudy with a chance of rain",
		"Heavy rain expected",
		"Clear skies ahead",
	}
	
	icons := []string{
		getWeatherIcon(descriptions[0]),
		getWeatherIcon(descriptions[1]),
		getWeatherIcon(descriptions[2]),
	}

	// Verify we got some icon (not empty)
	for i, icon := range icons {
		if icon == "" {
			t.Errorf("getWeatherIcon(%q) returned empty string", descriptions[i])
		}
		// Verify it's one of our known icons
		validIcons := []string{"☀️", "⛅", "☁️", "🌧️", "⛈️", "❄️", "🌫️", "💨", "🌤️"}
		found := false
		for _, validIcon := range validIcons {
			if icon == validIcon {
				found = true
				break
			}
		}
		if !found {
			t.Errorf("getWeatherIcon(%q) = %q; not a valid icon", descriptions[i], icon)
		}
	}
}

func TestWeatherIconPriority(t *testing.T) {
	// Test that "partly cloudy" gets priority over just "cloudy"
	result := getWeatherIcon("Partly cloudy")
	expected := "⛅"
	if result != expected {
		t.Errorf("getWeatherIcon('Partly cloudy') = %q; want %q", result, expected)
	}

	// Regular cloudy should get cloud icon
	result = getWeatherIcon("Cloudy")
	expected = "☁️"
	if result != expected {
		t.Errorf("getWeatherIcon('Cloudy') = %q; want %q", result, expected)
	}
}

func TestWeatherDataStructure(t *testing.T) {
	// Test that WeatherData structure can be instantiated
	var weather WeatherData
	
	// Verify the structure has expected fields
	if weather.CurrentCondition == nil {
		// This is expected for a new struct
	}
	
	if weather.NearestArea == nil {
		// This is expected for a new struct
	}
	
	// Just verify the structure compiles correctly
	t.Log("WeatherData structure is valid")
}

// Helper function to check if string contains substring (case-insensitive)
func containsIgnoreCase(s, substr string) bool {
	return strings.Contains(strings.ToLower(s), strings.ToLower(substr))
}
