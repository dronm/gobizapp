package api

import (
	"bytes"
	"encoding/json"
	"fmt"
)

func UnmarshalParams(payload []byte) ([]string, error) {
	var params []string

	decoder := json.NewDecoder(bytes.NewReader(payload))

	tok, err := decoder.Token()
	if err != nil {
		return nil, fmt.Errorf("decode.Token(): %v", err)
	}

	delim, ok := tok.(json.Delim)
	if !ok || delim != '{' {
		return nil, fmt.Errorf("expected json object")
	}

	for decoder.More() {
		//Do not need key here!!!
		_, err := decoder.Token()
		if err != nil {
			return nil, fmt.Errorf("error reading key: %v", err)
		}

		// key, ok := keyTok.(string)
		// if !ok {
		// 	ServeError(c, http.StatusBadRequest, funcName+" keyTok.(string)", fmt.Errorf("invalid key type"))
		// 	return
		// }
		//
		// Read value token
		var raw json.RawMessage
		if err := decoder.Decode(&raw); err != nil {
			return nil, fmt.Errorf("error decoding json value: %v", err)
		}

		var strVal string
		if err := json.Unmarshal(raw, &strVal); err != nil {
			// If not a string, keep raw JSON text
			strVal = string(raw)
		}

		params = append(params, strVal)
	}

	if _, err := decoder.Token(); err != nil {
		return nil, fmt.Errorf("error reading end of object: %v", err)
	}

	return params, nil
}
