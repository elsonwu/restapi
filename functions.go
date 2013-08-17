package restapi

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"reflect"
	"strings"
)

type Params struct {
	Query url.Values
}

type Map map[string]interface{}

func SetupOutput(Result bool, Data Map) Output {
	output := Output{}
	output.Result = Result
	output.Data = Data
	return output
}

type Output struct {
	Result bool                   `json:"result"`
	Data   map[string]interface{} `json:"data"`
}

type IApi interface{}

var apis map[string]IApi

func init() {
	apis = make(map[string]IApi)
}

func Add(apiName string, api IApi) {
	apiName = strings.ToLower(apiName)
	apis[apiName] = api
}

func Call(apiName, method string, params Params) Output {
	apiName = strings.ToLower(apiName)
	api, ok := apis[apiName]
	if ok {
		refValue := reflect.ValueOf(api)
		refValueParams := reflect.ValueOf(params)
		refReturn := refValue.MethodByName(method).Call([]reflect.Value{refValueParams})
		if v, ok := refReturn[0].Interface().(Output); ok {
			return v
		}
	} else {
		fmt.Println("api - " + apiName + " does not exist")
	}

	return Output{}
}

func Run() {
	http.HandleFunc("/", func(res http.ResponseWriter, req *http.Request) {
		apiName := strings.Trim(req.URL.Path, "/")
		params := Params{Query: req.URL.Query()}
		output := Call(apiName, req.Method, params)

		data, _ := json.Marshal(output)
		res.Header().Set("Content-Type", "application/json")
		res.Write([]byte(data))
	})

	http.ListenAndServe(":8888", nil)
}
