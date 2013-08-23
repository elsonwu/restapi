#Simple API

	package main

	import (
		"github.com/elsonwu/restapi"
	)

	type User struct {
		restapi.Api
	}

	//in this demo, we only allow view method
	func (self *User) View(ctx restapi.IContext) restapi.IOutput {
	    //result, data,        errors
	    //bool,   interface{}, []string
		return restapi.Output(true, ctx.Query(), nil)
	}


	func main() {
		restapi.Add("user", &User{})
		restapi.Run(":8888")
	}
	
	//Visit http://localhost:8888/user/xxx
	{
    	"data": {
        	"id": [
            	"xxx"
        	]
    	},
    	"errors": null,
    	"result": true
	}
	
#API calls another API

    //user api create method calls content api view method.
	func (self *User) Create(ctx restapi.IContext) restapi.IOutput {
		query := ctx.Query()
		query.Add("with_content_comments", "1")
		return restapi.Call("content", restapi.MethodView, ctx)
	}