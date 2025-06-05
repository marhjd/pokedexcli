package commands

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
)

type PokemonEncounter struct {
	EncounterMethodRates []struct {
		EncounterMethod struct {
			Name string `json:"name"`
			URL  string `json:"url"`
		} `json:"encounter_method"`
		VersionDetails []struct {
			Rate    int64 `json:"rate"`
			Version struct {
				Name string `json:"name"`
				URL  string `json:"url"`
			} `json:"version"`
		} `json:"version_details"`
	} `json:"encounter_method_rates"`
	GameIndex int64 `json:"game_index"`
	ID        int64 `json:"id"`
	Location  struct {
		Name string `json:"name"`
		URL  string `json:"url"`
	} `json:"location"`
	Name  string `json:"name"`
	Names []struct {
		Language struct {
			Name string `json:"name"`
			URL  string `json:"url"`
		} `json:"language"`
		Name string `json:"name"`
	} `json:"names"`
	PokemonEncounters []struct {
		Pokemon struct {
			Name string `json:"name"`
			URL  string `json:"url"`
		} `json:"pokemon"`
		VersionDetails []struct {
			EncounterDetails []struct {
				Chance          int64         `json:"chance"`
				ConditionValues []interface{} `json:"condition_values"`
				MaxLevel        int64         `json:"max_level"`
				Method          struct {
					Name string `json:"name"`
					URL  string `json:"url"`
				} `json:"method"`
				MinLevel int64 `json:"min_level"`
			} `json:"encounter_details"`
			MaxChance int64 `json:"max_chance"`
			Version   struct {
				Name string `json:"name"`
				URL  string `json:"url"`
			} `json:"version"`
		} `json:"version_details"`
	} `json:"pokemon_encounters"`
}

func commandExplore(cfg *Config, name string) error {
	if name == "" {
		return errors.New("no name provided")
	}
	url := baseURL + "location-area/" + name

	if shouldReturn, err := checkExploreCache(cfg, url); shouldReturn {
		return err
	}

	res, err := http.Get(url)
	if err != nil {
		return err
	}
	defer res.Body.Close()

	pokemon := new(PokemonEncounter)
	b, err := io.ReadAll(res.Body)
	if err != nil {
		return err
	}
	err = json.Unmarshal(b, pokemon)
	cfg.Cache.Add(url, b)

	printExploreResults(pokemon)

	return nil
}

func checkExploreCache(cfg *Config, url string) (bool, error) {
	if val, ok := cfg.Cache.Get(url); ok {
		pokemon := new(PokemonEncounter)
		err := json.Unmarshal(val, pokemon)
		if err != nil {
			return true, err
		}
		printExploreResults(pokemon)
		return true, nil
	}
	return false, nil
}

func printExploreResults(pokemon *PokemonEncounter) {
	for _, encounter := range pokemon.PokemonEncounters {
		fmt.Println(encounter.Pokemon.Name)
	}
}
