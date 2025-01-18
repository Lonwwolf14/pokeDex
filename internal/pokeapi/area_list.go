package pokeapi

import (
	"encoding/json"
	"io"
	"net/http"
)

func (c *Client) GetLocationArea(areaName string) (LocationArea, error) {
	url := baseUrl + "/location-area/" + areaName

	if cacheData, ok := c.cache.Get(url); ok {
		var locationArea LocationArea
		err := json.Unmarshal(cacheData, &locationArea)
		if err != nil {
			return LocationArea{}, err
		}
		return locationArea, nil
	}

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return LocationArea{}, err
	}
	resp, err := c.httpClient.Do(req)
	if err != nil {
		return LocationArea{}, err
	}
	defer resp.Body.Close()

	data, err := io.ReadAll(resp.Body)
	if err != nil {
		return LocationArea{}, err
	}

	var locationArea LocationArea
	err = json.Unmarshal(data, &locationArea)
	if err != nil {
		return LocationArea{}, err
	}
	c.cache.Add(url, data)
	return locationArea, nil

}
