package pokeapi

import (
	"encoding/json"
	"errors"
	"io"
	"net/http"
)

// We request the PokeAPI for location data and unmarshal it into a locationObj struct.
// Meant to be used by the mpas commands to print out "location-area" data.
func (c *Client) GetLocations(path *string) (locationObj, error) {
	// find current path
	currentURL := api + "location-area/"
	if path != nil {
		currentURL = *path
	}

	// construct request
	req, err := http.NewRequest("GET", currentURL, nil)
	if err != nil {
		return locationObj{}, err
	}

	// "do" GET request
	res, err := c.httpClient.Do(req)
	if err != nil {
		return locationObj{}, err
	}

	// get response body
	// a call to res.Body is streamed on demand and as such needs to be closed -> res.Body.Close()
	body, err := io.ReadAll(res.Body)
	res.Body.Close()
	if res.StatusCode > 299 {
		return locationObj{}, errors.New("Invalid Status Code greater than 299")
	}
	if err != nil {
		return locationObj{}, err
	}

	// unmarshal the raw data into the `locationObj` struct
	data := locationObj{}
	errs := json.Unmarshal(body, &data)
	if errs != nil {
		return locationObj{}, errs
	}

	// return
	return data, nil
}
