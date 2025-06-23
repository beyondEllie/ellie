package actions

import (
	"encoding/json"
	"io"
	"net/http"
	"time"

	"github.com/tacheraSasi/ellie/styles"
	"github.com/tacheraSasi/ellie/utils"
)

type JokeData struct {
	Joke string `json:"joke"`
}

func Joke() {
	done := make(chan bool)
	go utils.ShowLoadingSpinner("Fetching a hilarious joke for you...", done)

	client := &http.Client{Timeout: 10 * time.Second}
	req, err := http.NewRequest("GET", "https://icanhazdadjoke.com/", nil)
	if err != nil {
		done <- true
		styles.GetErrorStyle().Println("❌ Failed to create request.")
		return
	}
	req.Header.Add("Accept", "application/json")

	resp, err := client.Do(req)
	if err != nil {
		done <- true
		styles.GetErrorStyle().Println("❌ Failed to fetch joke. Please check your internet connection!")
		return
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		done <- true
		styles.GetErrorStyle().Println("❌ Failed to read joke data.")
		return
	}

	var joke JokeData
	err = json.Unmarshal(body, &joke)
	done <- true
	if err != nil {
		styles.GetErrorStyle().Println("❌ Failed to parse joke data.")
		return
	}
	if joke.Joke != "" {
		styles.GetSuccessStyle().Printf("\n😂 %s\n\n", joke.Joke)
	} else {
		styles.GetErrorStyle().Println("❌ Could not retrieve a joke. Try again later!")
	}
}
