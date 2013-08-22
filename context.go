package restapi

import (
	"net/http"
)

type IContext interface {
	Query() IQuery
	Req() *http.Request
	Res() http.ResponseWriter
}

type Context struct {
	query IQuery
	req   *http.Request
	res   http.ResponseWriter
}

func (self *Context) Query() IQuery {
	return self.query
}

func (self *Context) Req() *http.Request {
	return self.req
}

func (self *Context) Res() http.ResponseWriter {
	return self.res
}
