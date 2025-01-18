package pokeapi

import (
	"encoding/json"
	"io"
	"net/http"
)

func (c *Client) ListLocations(pageUrl *string) (RespShallowLocation, error) {
	url := baseUrl + "/location-area"
	if pageUrl != nil {
		url = *pageUrl
	}

	if cacheData, ok := c.cache.Get(url); ok {
		var cacheResp RespShallowLocation
		err := json.Unmarshal(cacheData, &cacheResp)
		if err != nil {
			return RespShallowLocation{}, err
		}
		return cacheResp, nil
	}

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return RespShallowLocation{}, err
	}
	resp, err := c.httpClient.Do(req)
	if err != nil {
		return RespShallowLocation{}, err
	}
	defer resp.Body.Close()

	data, err := io.ReadAll(resp.Body)
	if err != nil {
		return RespShallowLocation{}, err
	}

	locationResp := RespShallowLocation{}
	err = json.Unmarshal(data, &locationResp)
	if err != nil {
		return RespShallowLocation{}, err
	}
	c.cache.Add(url, data)
	return locationResp, nil

}
