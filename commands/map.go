package commands

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type LocationArea struct {
	Count    int          `json:"count"`
	Next     string       `json:"next"`
	Previous string       `json:"previous"`
	Results  []AreaResult `json:"results"`
}

type AreaResult struct {
	Name string `json:"name"`
	URL  string `json:"url"`
}

func commandMap(cfg *Config) error {
	res, err := http.Get(cfg.GetNext())
	defer res.Body.Close()
	if err != nil {
		return err
	}

	locationArea := new(LocationArea)
	err = json.NewDecoder(res.Body).Decode(locationArea)
	if err != nil {
		return err
	}
	cfg.Next = locationArea.Next
	cfg.Previous = locationArea.Previous

	for _, result := range locationArea.Results {
		fmt.Println(result.Name)
	}
	return nil
}

func commandMapb(cfg *Config) error {
	// TODO
	panic("implement me")
	return nil
}
