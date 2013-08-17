package restapi

import ()

func SetupOutput(Result bool, Data Map, errors []string, StatusCode int) IOutput {
	output := &Output{}
	output.result = Result
	output.data = Data
	output.errors = errors
	output.statusCode = StatusCode
	return output
}

type IOutput interface {
	Result() bool
	Data() map[string]interface{}
	Errors() []string
	StatusCode() int
}

type Output struct {
	result     bool                   `json:"result"`
	data       map[string]interface{} `json:"data"`
	errors     []string               `json:"errors"`
	statusCode int                    `json:"-"`
}

func (self *Output) Result() bool {
	return self.result
}

func (self *Output) Data() map[string]interface{} {
	return self.data
}

func (self *Output) Errors() []string {
	return self.errors
}

func (self *Output) StatusCode() int {
	return self.statusCode
}
