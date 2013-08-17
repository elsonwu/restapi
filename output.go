package restapi

import ()

func Output(Result bool, Data Map, errors []string) IOutput {
	o := &output{}
	o.result = Result
	o.data = Data
	o.errors = errors
	return o
}

type IOutput interface {
	Result() bool
	Data() map[string]interface{}
	Errors() []string
}

type output struct {
	result bool                   `json:"result"`
	data   map[string]interface{} `json:"data"`
	errors []string               `json:"errors"`
}

func (self *output) Result() bool {
	return self.result
}

func (self *output) Data() map[string]interface{} {
	return self.data
}

func (self *output) Errors() []string {
	return self.errors
}
