package pokescraper

import (
	"encoding/json"
	"fmt"
)

// returns Pokemon data as JSON string
func SerializePokemon(poke Pokemon) (string, error) {
	invalidErr := ValidatePokemon(poke)
	if invalidErr != nil {
		return "", invalidErr
	}

	pokeJSON, err := json.MarshalIndent(poke, "", "    ")
	if err != nil {
		return "", err
	}

	return string(pokeJSON), nil
}

// returns Pokemon data as struct
func DeserializePokemon(pokeJSON string, pokePoint *Pokemon) error {
	err := json.Unmarshal([]byte(pokeJSON), pokePoint)
	if err != nil {
		return fmt.Errorf("error unmarshaling deserialized data: %v", err)
	}

	invalidErr := ValidatePokemon(*pokePoint)
	if invalidErr != nil {
		return invalidErr
	}

	return nil
}
