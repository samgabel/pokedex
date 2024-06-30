package pokeapi

import "encoding/json"

func (c *Client) GetEncounters(location string) (encounterObj, error) {
	// make path
	path := api + "location-area/" + location

	// check to see if the path is in the cache and unmarshal + return early
	if body, ok := c.cache.Get(path); ok {
		data := encounterObj{}
		err := json.Unmarshal(body, &data)
		if err != nil {
			return encounterObj{}, err
		}

		return data, nil
	}

	// get response body
	body, err := c.getResponseBody(path)
	if err != nil {
		return encounterObj{}, err
	}

	// cache the response
	c.cache.Add(path, body)

	// unmarshal and return encounters as an encounterObj
	data := encounterObj{}
	err = json.Unmarshal(body, &data)
	if err != nil {
		return encounterObj{}, err
	}

	return data, nil
}
