package utils

import "encoding/json"

// JSONMarshal - Returns the JSON string representaion of source
func JSONMarshal(source any, pretty bool) (output string, err error) {
	// This is encapsulated so we can swap out the serializer if needed

	var bytes []byte

	if pretty {
		bytes, err = json.MarshalIndent(source, "", "    ")
		if err != nil {
			return "", err
		}
	} else {
		bytes, err = json.Marshal(source)
		if err != nil {
			return "", err
		}
	}

	output = string(bytes)

	return output, nil
}
