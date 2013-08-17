package restapi

import (
	"net/http"
)

type IRuntimes interface {
	Req() *http.Request
	Res() http.ResponseWriter
}

type Runtimes struct {
	res http.ResponseWriter
	req *http.Request
}

func (self *Runtimes) Req() *http.Request {
	return self.req
}

func (self *Runtimes) Res() http.ResponseWriter {
	return self.res
}
