package restapi

import ()

type IApi interface {
	GET(Params, IRuntimes) Output
	POST(Params, IRuntimes) Output
	PUT(Params, IRuntimes) Output
	DELETE(Params, IRuntimes) Output
}

type Api struct{}

func (self *Api) GET(params Params, rs IRuntimes) Output {
	return SetupOutput(false, Map{}, []string{"Does not support get method"})
}

func (self *Api) POST(params Params, rs IRuntimes) Output {
	return SetupOutput(false, Map{}, []string{"Does not support post method"})
}

func (self *Api) PUT(params Params, rs IRuntimes) Output {
	return SetupOutput(false, Map{}, []string{"Does not support put method"})
}

func (self *Api) DELETE(params Params, rs IRuntimes) Output {
	return SetupOutput(false, Map{}, []string{"Does not support delete method"})
}
