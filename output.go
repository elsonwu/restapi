package restapi

import ()

func Output(Result bool, Data interface{}, errors []string) IOutput {
	o := &output{}
	o.result = Result
	o.data = Data
	o.errors = errors
	return o
}

type IOutput interface {
	Result() bool
	Data() interface{}
	Errors() []string
}

type output struct {
	result bool        `json:"result"`
	data   interface{} `json:"data"`
	errors []string    `json:"errors"`
}

func (self *output) Result() bool {
	return self.result
}

func (self *output) Data() interface{} {
	return self.data
}

func (self *output) Errors() []string {
	return self.errors
}
