package main

import "fmt"

func commandMapF(cfg *config, args ...string) error {
	locationResp, err := cfg.pokeapiClient.ListLocations(cfg.nextLocationUrl)
	if err != nil {
		return err
	}
	cfg.nextLocationUrl = locationResp.Next
	cfg.prevLocationUrl = locationResp.Previous

	for _, location := range locationResp.Results {
		fmt.Println(location.Name)
	}
	return nil
}

func commandMapB(cfg *config, args ...string) error {
	if cfg.prevLocationUrl == nil {
		fmt.Println("You are on the first page")
		return nil
	}
	locationResp, err := cfg.pokeapiClient.ListLocations(cfg.prevLocationUrl)
	if err != nil {
		return err
	}
	cfg.nextLocationUrl = locationResp.Next
	cfg.prevLocationUrl = locationResp.Previous

	for _, location := range locationResp.Results {
		fmt.Println(location.Name)
	}
	return nil
}
