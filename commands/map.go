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
	if err != nil {
		return err
	}
	defer res.Body.Close()

	locationArea, err := setConfigLocationAreas(cfg, res)
	if err != nil {
		return err
	}
	printLocationAreaResults(*locationArea)

	return nil
}

func commandMapb(cfg *Config) error {
	url, err := cfg.GetPrevious()
	if err != nil {
		return err
	}
	res, err := http.Get(url)
	if err != nil {
		return err
	}
	defer res.Body.Close()

	locationArea, err := setConfigLocationAreas(cfg, res)
	if err != nil {
		return err
	}
	printLocationAreaResults(*locationArea)

	return nil
}

func setConfigLocationAreas(cfg *Config, res *http.Response) (*LocationArea, error) {
	locationArea := new(LocationArea)
	err := json.NewDecoder(res.Body).Decode(locationArea)
	if err != nil {
		return nil, err
	}
	cfg.Next = locationArea.Next
	cfg.Previous = locationArea.Previous
	return locationArea, nil
}

func printLocationAreaResults(locationArea LocationArea) {
	for _, result := range locationArea.Results {
		fmt.Println(result.Name)
	}
}
