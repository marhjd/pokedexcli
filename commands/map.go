package commands

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
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

func commandMap(cfg *Config, _ string) error {
	url := cfg.GetNext()

	if shouldReturn, err := checkCache(cfg, url); shouldReturn {
		return err
	}

	res, err := http.Get(url)
	if err != nil {
		return err
	}
	defer res.Body.Close()
	locationArea, err := extractLocationArea(res.Body)
	if err != nil {
		return err
	}

	setNextAndPrevLocationArea(cfg, locationArea)
	b := bytes.NewBuffer(make([]byte, 0))
	e := json.NewEncoder(b)
	err = e.Encode(locationArea)
	if err != nil {
		return err
	}
	cfg.Cache.Add(url, b.Bytes())
	printLocationAreaResults(*locationArea)

	return nil
}

func commandMapb(cfg *Config, _ string) error {
	url, err := cfg.GetPrevious()
	if err != nil {
		return err
	}

	if shouldReturn, err := checkCache(cfg, url); shouldReturn {
		return err
	}

	res, err := http.Get(url)
	if err != nil {
		return err
	}
	defer res.Body.Close()

	locationArea, err := extractLocationArea(res.Body)
	if err != nil {
		return err
	}

	setNextAndPrevLocationArea(cfg, locationArea)
	b := bytes.NewBuffer(make([]byte, 0))
	e := json.NewEncoder(b)
	err = e.Encode(locationArea)
	if err != nil {
		return err
	}
	cfg.Cache.Add(url, b.Bytes())
	printLocationAreaResults(*locationArea)

	return nil
}

func extractLocationArea(reader io.Reader) (*LocationArea, error) {
	locationArea := new(LocationArea)
	err := json.NewDecoder(reader).Decode(locationArea)
	if err != nil {
		return nil, err
	}
	return locationArea, nil
}

func setNextAndPrevLocationArea(cfg *Config, locationArea *LocationArea) {
	cfg.Next = locationArea.Next
	cfg.Previous = locationArea.Previous
}

func checkCache(cfg *Config, url string) (bool, error) {
	if val, ok := cfg.Cache.Get(url); ok {
		locationArea := new(LocationArea)
		err := json.Unmarshal(val, locationArea)
		if err != nil {
			return true, err
		}
		setNextAndPrevLocationArea(cfg, locationArea)
		printLocationAreaResults(*locationArea)
		return true, nil
	}
	return false, nil
}

func printLocationAreaResults(locationArea LocationArea) {
	for _, result := range locationArea.Results {
		fmt.Println(result.Name)
	}
}
