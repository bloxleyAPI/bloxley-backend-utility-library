package utils

import "encoding/json"

func MbanqBody(body []byte) (map[string]interface{}, error) {
	var responseBody map[string]interface{}
	if err := json.Unmarshal(body, &responseBody); err != nil {
		return nil, err
	}
	return responseBody, nil
}
