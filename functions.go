package restapi

import (
	"encoding/json"
	"net/http"
	// "net/url"
	"strings"
)

const (
	MethodGet    string = "GET"
	MethodPost   string = "POST"
	MethodPut    string = "PUT"
	MethodDelete string = "DELETE"
)

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

var apis map[string]IApi

func init() {
	apis = make(map[string]IApi)
}

func bind(apiName string, api IApi) {
	http.HandleFunc("/"+apiName, func(res http.ResponseWriter, req *http.Request) {
		params := Params{}
		params.Query = req.URL.Query()

		handler := new(Handler)
		handler.res = res
		handler.req = req

		rs := new(Runtimes)
		rs.res = res
		rs.req = req

		output := handler.innerCall(api, req.Method, params, rs)
		data, _ := json.Marshal(output)
		res.Header().Set("Content-Type", "application/json")
		res.Write([]byte(data))
	})
}

func Add(apiName string, api IApi) {
	bind(apiName, api)
	apiName = strings.ToLower(apiName)
	apis[apiName] = api
}

func innerCall(api IApi, method string, params Params, rs IRuntimes) Output {

	var output Output
	if MethodGet == method {
		output = api.GET(params, rs)
	} else if MethodPost == method {
		output = api.POST(params, rs)
	} else if MethodPut == method {
		output = api.PUT(params, rs)
	} else if MethodDelete == method {
		output = api.DELETE(params, rs)
	}

	return output
}

func Call(apiName, method string, params Params, rs IRuntimes) Output {
	apiName = strings.ToLower(apiName)
	api, ok := apis[apiName]
	if ok {
		return innerCall(api, method, params, rs)
	}

	return SetupOutput(false, nil, []string{"api - " + apiName + " does not exist"})
}

func Run(bindString string) {
	http.ListenAndServe(bindString, nil)
}
