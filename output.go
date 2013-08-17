package restapi

import ()

func SetupOutput(Result bool, Data Map, errors []string) Output {
	output := Output{}
	output.Result = Result
	output.Data = Data
	output.Errors = errors
	return output
}

type Output struct {
	Result bool                   `json:"result"`
	Data   map[string]interface{} `json:"data"`
	Errors []string               `json:"errors"`
}
