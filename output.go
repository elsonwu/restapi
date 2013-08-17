package restapi

import ()

func SetupOutput(Result bool, Data Map, errors []string, StatusCode int) Output {
	output := Output{}
	output.Result = Result
	output.Data = Data
	output.Errors = errors
	output.StatusCode = StatusCode
	return output
}

type Output struct {
	Result     bool                   `json:"result"`
	Data       map[string]interface{} `json:"data"`
	Errors     []string               `json:"errors"`
	StatusCode int                    `json:"-"`
}
