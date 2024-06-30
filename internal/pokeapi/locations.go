package pokeapi

import "encoding/json"

// We request the PokeAPI for location data and unmarshal it into a locationObj struct.
// Meant to be used by the mpas commands to print out "location-area" data.
func (c *Client) GetLocations(path *string) (locationObj, error) {
	// find current path
	currentURL := api + "location-area/"
	if path != nil {
		currentURL = *path
	}

	// check cache for path, unmarshal and return early if so
	if body, ok := c.cache.Get(currentURL); ok {
		data := locationObj{}
		err := json.Unmarshal(body, &data)
		if err != nil {
			return locationObj{}, err
		}

		return data, nil
	}

	// get body
	body, err := c.getResponseBody(currentURL)
	if err != nil {
		return locationObj{}, err
	}

	// add to cache
	c.cache.Add(currentURL, body)

	// unmarshal and return
	data := locationObj{}
	err = json.Unmarshal(body, &data)
	if err != nil {
		return locationObj{}, err
	}

	return data, nil
}
