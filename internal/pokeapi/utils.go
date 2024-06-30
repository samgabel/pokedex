package pokeapi

import (
	"errors"
	"fmt"
	"io"
	"net/http"
)

// This method is meant to streamline the process of using the pokeapi.Client for GET requests.
// It returns the response body and is used by pokiapi functions within locations.go and encounters.go.
func (c *Client) getResponseBody(url string) ([]byte, error){
	// construct request
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}

	// "do" GET request
	res, err := c.httpClient.Do(req)
	if err != nil {
		return nil, err
	}
	if res.StatusCode > 299 {
		return nil, errors.New(fmt.Sprint("Returned Invalid Status Code: ", res.Status))
	}

	// get response body
	// a call to res.Body is streamed on demand and as such needs to be closed -> res.Body.Close()
	body, err := io.ReadAll(res.Body)
	res.Body.Close()
	if err != nil {
		return nil, err
	}

	return body, nil
}
