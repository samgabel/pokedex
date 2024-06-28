package commands

import (
	"fmt"
	"regexp"
	"strconv"
)

func commandMapf(cfg *Config) error {
	// make a call to the pokeapi
	data, err := cfg.pokeapiClient.GetLocations(cfg.nextLocationURL)
	if err != nil {
		return err
	}

	// need to update the appropriate (next and previous) pages
	cfg.nextLocationURL = data.Next
	cfg.previousLocationURL = data.Previous

	// print page num
	if num, err := findPageNum(cfg.nextLocationURL); err == nil {
		fmt.Printf("(%v)\n", num)
	}

	// we need to range over the slice of `Results` structs
	// this way we can access the `.Name` field
	for _, result := range data.Results {
		fmt.Println(result.Name)
	}

	return nil
}

func commandMapb(cfg *Config) error {
	data, err := cfg.pokeapiClient.GetLocations(cfg.previousLocationURL)
	if err != nil {
		return err
	}

	// need to update the appropriate (next and previous) pages
	cfg.nextLocationURL = data.Next
	cfg.previousLocationURL = data.Previous

	// print page num
	if num, err := findPageNum(cfg.nextLocationURL); err == nil {
		fmt.Printf("(%v)\n", num)
	}

	// we need to range over the slice of `Results` structs
	// this way we can access the `.Name` field
	for _, result := range data.Results {
		fmt.Println(result.Name)
	}

	return nil
}

func findPageNum(url *string) (int, error) {
	// compile regexp pattern
	re := regexp.MustCompile(`offset=(\d+)`)

	// find the match
	match := re.FindStringSubmatch(*url)

	// string to int
	num, err := strconv.Atoi(match[1])
	if err != nil {
		return 0, err
	}

	// return int
	return (num / 20), nil
}
