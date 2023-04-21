package response_models

import "encoding/json"

type ErrorResponse struct {
	Error string `json:"error"`
}

func Error(err error) ErrorResponse {
	return ErrorResponse{
		Error: err.Error(),
	}
}

func ErrorJson(err error) string {
	return Error(err).ToJson()
}

func (e ErrorResponse) ToJson() string {
	bytes, err := json.Marshal(e)
	if err != nil {
		panic(err)
	}

	return string(bytes)
}
