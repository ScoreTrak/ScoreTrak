package handler

import "github.com/ogen-go/ogen/json"

type unmarshaller func(data []byte) error

func convertEntityToResponse(data any, um unmarshaller) error {
	res, err := json.Marshal(data)
	if err != nil {
		return err
	}

	err = um(res)
	return err
}

// ConvertStruct accepts source and destination structs, and converts source to destination
func ConvertStruct(source interface{}, destination interface{}) error {
	// Convert the source struct to JSON
	sourceJSON, err := json.Marshal(source)
	if err != nil {
		return err
	}

	// Unmarshal JSON into the destination struct
	err = json.Unmarshal(sourceJSON, destination)
	if err != nil {
		return err
	}

	return nil
}
