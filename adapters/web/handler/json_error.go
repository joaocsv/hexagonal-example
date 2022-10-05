package handler

import "encoding/json"

func errorJson(msg string) []byte {
	error := struct {
		Message string `json:"message"`
	}{
		msg,
	}

	r, err := json.Marshal(error)

	if err != nil {
		return []byte(err.Error())
	}

	return r
}
