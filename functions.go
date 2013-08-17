package restapi

import (
	"encoding/json"
	"net/http"
	// "net/url"
	// "fmt"
	// "reflect"
	"strings"
)

const (
	MethodGet    string = "GET"
	MethodPost   string = "POST"
	MethodPut    string = "PUT"
	MethodDelete string = "DELETE"
)

var _apis apis

func init() {
	_apis = make(apis)
}

type IQuery interface {
	Add(string, string)
	Del(string)
	Get(string) string
	Set(string, string)
}

type Params struct {
	Query IQuery
}

type Map map[string]interface{}

type apis map[string]IApi

func (self apis) Add(apiName string, api IApi) {
	apiName = strings.ToLower(apiName)
	self[apiName] = api
}

func (self apis) Get(apiName string) IApi {
	apiName = strings.ToLower(apiName)
	api, ok := self[apiName]
	if ok {
		return api
	}

	return nil
}

func bind(apiName string) {
	http.HandleFunc("/"+apiName, func(res http.ResponseWriter, req *http.Request) {
		params := Params{}
		params.Query = req.URL.Query()
		handler := new(Handler)
		handler.req = req
		handler.res = res
		output := handler.Call(apiName, req.Method, params)
		data, _ := json.Marshal(output)
		res.Header().Set("Content-Type", "application/json")
		res.Write([]byte(data))
	})
}

func Add(apiName string, api IApi) {
	_apis.Add(apiName, api)
	bind(apiName)
}

func Run(bindString string) {
	http.ListenAndServe(bindString, nil)
}
