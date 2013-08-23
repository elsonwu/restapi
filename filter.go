package restapi

import (
	"strings"
)

var Filter *filterHandler

func init() {
	Filter = new(filterHandler)
	Filter.filters = make(map[string][]filterFunc)
}

type filterFunc func(ctx IContext) error

type filterHandler struct {
	filters map[string][]filterFunc
}

func (self *filterHandler) On(filterName string, filterFunc filterFunc) {
	filterName = strings.ToLower(filterName)
	self.filters[filterName] = append(self.filters[filterName], filterFunc)
}

func (self *filterHandler) Emit(filterName string, ctx IContext) error {
	filterName = strings.ToLower(filterName)
	fileters, ok := self.filters[filterName]

	if !ok {
		return nil
	}

	for _, fn := range fileters {
		err := fn(ctx)
		if nil != err {
			return err
		}
	}

	return nil
}
