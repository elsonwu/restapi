package restapi

import (
	"encoding/json"
	// "fmt"
	"net/http"
	"net/url"
	// "reflect"
	"strings"
)

const (
	MethodGet    string = "GET"
	MethodPost   string = "POST"
	MethodPut    string = "PUT"
	MethodDelete string = "DELETE"
)

type Params struct {
	Query url.Values
}

type Map map[string]interface{}

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

type IApi interface {
	GET(Params) Output
	POST(Params) Output
	PUT(Params) Output
	DELETE(Params) Output
}

type Api struct{}

func (self *Api) GET(params Params) Output {
	return SetupOutput(false, Map{}, []string{"Does not support get method"})
}

func (self *Api) POST(params Params) Output {
	return SetupOutput(false, Map{}, []string{"Does not support post method"})
}

func (self *Api) PUT(params Params) Output {
	return SetupOutput(false, Map{}, []string{"Does not support put method"})
}

func (self *Api) DELETE(params Params) Output {
	return SetupOutput(false, Map{}, []string{"Does not support delete method"})
}

var apis map[string]IApi

func init() {
	apis = make(map[string]IApi)
}

func bind(apiName string, api IApi) {
	http.HandleFunc("/"+apiName, func(res http.ResponseWriter, req *http.Request) {
		params := Params{Query: req.URL.Query()}
		output := innerCall(api, req.Method, params)
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

func innerCall(api IApi, method string, params Params) Output {

	var output Output
	if MethodGet == method {
		output = api.GET(params)
	} else if MethodPost == method {
		output = api.POST(params)
	} else if MethodPut == method {
		output = api.PUT(params)
	} else if MethodDelete == method {
		output = api.DELETE(params)
	}

	return output
}

func Call(apiName, method string, params Params) Output {
	apiName = strings.ToLower(apiName)
	api, ok := apis[apiName]
	if ok {
		return innerCall(api, method, params)
	}

	return SetupOutput(false, nil, []string{"api - " + apiName + " does not exist"})
}

func Run(bindString string) {
	http.ListenAndServe(bindString, nil)
}
