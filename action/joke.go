package actions

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"

	"github.com/tacheraSasi/ellie/styles"
	"github.com/tacheraSasi/ellie/utils"
)

const (
	jokeAPIEndpoint = "https://icanhazdadjoke.com/"
	requestTimeout  = 10 * time.Second
)

type JokeResponse struct {
	Joke string `json:"joke"`
}

// Joke fetches and displays a random dad joke from the API
func Joke() {
	done := make(chan bool)
	go utils.ShowLoadingSpinner("Fetching a hilarious joke for you...", done)

	joke, err := fetchJoke()
	done <- true

	if err != nil {
		styles.GetErrorStyle().Printf("\n❌ %s\n\n", err)
		return
	}

	if joke == "" {
		styles.GetErrorStyle().Println("\n❌ Could not retrieve a joke. Try again later!")
		return
	}

	styles.GetSuccessStyle().Printf("\n😂 %s\n\n", joke)
}

// fetchJoke handles the API request and response processing
func fetchJoke() (string, error) {
	client := &http.Client{Timeout: requestTimeout}
	
	req, err := http.NewRequest(http.MethodGet, jokeAPIEndpoint, nil)
	if err != nil {
		return "", fmt.Errorf("failed to create request: %v", err)
	}
	req.Header.Set("Accept", "application/json")

	resp, err := client.Do(req)
	if err != nil {
		return "", fmt.Errorf("failed to fetch joke. Please check your internet connection: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return "", fmt.Errorf("API request failed with status: %s", resp.Status)
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", fmt.Errorf("failed to read joke data: %v", err)
	}

	var jokeData JokeResponse
	if err := json.Unmarshal(body, &jokeData); err != nil {
		return "", fmt.Errorf("failed to parse joke data: %v", err)
	}

	return jokeData.Joke, nil
}