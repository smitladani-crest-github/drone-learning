package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type JokeResponse struct {
	Joke string `json:"joke"`
}

func main() {
	// get the random jokes from https://icanhazdadjoke.com/api
	req, _ := http.NewRequest("GET", "https://icanhazdadjoke.com/", nil)

	req.Header = http.Header{
		"Accept": []string{"application/json"},
	}

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("Error executing GET request. ", err)
		return
	}
	defer resp.Body.Close()

	joke := &JokeResponse{}
	json.NewDecoder(resp.Body).Decode(joke)

	fmt.Println("Today's joke: ", joke.Joke)
}
