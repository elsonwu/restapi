package restapi

import (
	"encoding/json"
	// "errors"
	"net/http"
	// "net/url"
	// "fmt"
	// "reflect"
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

func bind(apiName string) {
	http.HandleFunc("/"+apiName, func(res http.ResponseWriter, req *http.Request) {
		params := Params{}
		params.Query = req.URL.Query()
		handler := new(Handler)
		handler.req = req
		output := handler.Call(apiName, req.Method, params)
		if 0 == output.StatusCode || 200 == output.StatusCode {
			data, _ := json.Marshal(output)
			res.Header().Set("Content-Type", "application/json")
			res.Write([]byte(data))
		} else {
		    //@todo handle error
			http.Error(res, "", output.StatusCode)
		}
	})
}

func Add(apiName string, api IApi) {
	_apis.Add(apiName, api)
	bind(apiName)
}

func Run(bindString string) {
	http.ListenAndServe(bindString, nil)
}
