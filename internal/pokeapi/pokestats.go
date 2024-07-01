package pokeapi

import "encoding/json"

func (c *Client) GetPokemonStats (pokemon string) (PokemonObj, error) {
	// make path
	path := api + "pokemon/" + pokemon
	
	// check to see if the path is in the cache and unmarshal + return early
	if body, ok := c.cache.Get(path); ok {
		data := PokemonObj{}
		err := json.Unmarshal(body, &data)
		if err != nil {
			return PokemonObj{}, err
		}

		return data, nil
	}

	// get response body
	body, err := c.getResponseBody(path)
	if err != nil {
		return PokemonObj{}, err
	}

	// cache the response
	c.cache.Add(path, body)

	// unmarshal and return encounters as an PokemonObj
	data := PokemonObj{}
	err = json.Unmarshal(body, &data)
	if err != nil {
		return PokemonObj{}, err
	}

	return data, nil
}
