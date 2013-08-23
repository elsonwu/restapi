package restapi

import (
	"encoding/json"
	"net/http"
)

var Conf *Config

func init() {
	Conf = &Config{
		ResponseFunc: DefaultResponseFunc,
		RouterFunc:   DefaultRouterFunc,
	}
}

type ResponseFunc func(output IOutput, ctx IContext, res http.ResponseWriter, req *http.Request)

type Config struct {
	ResponseFunc ResponseFunc
	RouterFunc   RouterFunc
}

func DefaultResponseFunc(output IOutput, ctx IContext, res http.ResponseWriter, req *http.Request) {
	o := map[string]interface{}{}
	o["result"] = output.Result()
	o["data"] = output.Data()
	o["errors"] = output.Errors()
	data, _ := json.Marshal(o)
	res.Header().Set("Content-Type", "application/json")
	res.Write([]byte(data))
}
