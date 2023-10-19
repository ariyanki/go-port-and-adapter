package automap

import (
	"encoding/json"
)

func AutoMap(source interface{}, destination interface{}) error {
	sourceJSON, err := json.Marshal(source)
	if err != nil {
		return err
	}
	err = json.Unmarshal(sourceJSON, destination)
	if err != nil {
		return err
	}
	return nil
}
