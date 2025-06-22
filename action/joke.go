package actions

import (
	"encoding/json"
	"io"
	"net/http"

	"github.com/tacheraSasi/ellie/styles"
)

type JokeData struct {
	Joke string `json:"joke"`
}

func Joke() {
	styles.InfoStyle.Println("Fetching a random joke...")
	client := &http.Client{}
	req, err := http.NewRequest("GET", "https://icanhazdadjoke.com/", nil)
	if err != nil {
		styles.ErrorStyle.Println("Failed to create request.")
		return
	}
	req.Header.Add("Accept", "application/json")

	resp, err := client.Do(req)
	if err != nil {
		styles.ErrorStyle.Println("Failed to fetch joke.")
		return
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		styles.ErrorStyle.Println("Failed to read joke data.")
		return
	}

	var joke JokeData
	err = json.Unmarshal(body, &joke)
	if err != nil {
		styles.ErrorStyle.Println("Failed to parse joke data.")
		return
	}
	if joke.Joke != "" {
		styles.SuccessStyle.Println(joke.Joke)
	} else {
		styles.ErrorStyle.Println("Could not retrieve joke.")
	}
}
