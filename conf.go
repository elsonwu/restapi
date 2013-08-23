package restapi

import (
	"encoding/json"
	// "net/http"
)

var Conf *Config

func init() {
	Conf = &Config{
		ResponseFunc: DefaultResponseFunc,
		RouterFunc:   DefaultRouterFunc,
	}
}

type ResponseFunc func(output IOutput, ctx IContext)

type Config struct {
	ResponseFunc ResponseFunc
	RouterFunc   RouterFunc
}

func DefaultResponseFunc(output IOutput, ctx IContext) {
	o := map[string]interface{}{}
	o["result"] = output.Result()
	o["data"] = output.Data()
	o["errors"] = output.Errors()
	data, _ := json.Marshal(o)
	ctx.Res().Header().Set("Content-Type", "application/json")
	ctx.Res().Write([]byte(data))
}
