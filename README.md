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
	
#Api filter
	package main

	import (
		"errors"
		"github.com/elsonwu/restapi"
	)

	type User struct {
		restapi.Api
	}

	//filter for any call to user API
	func (self *User) BeforeRun(ctx restapi.IContext) error {
		if "" == ctx.Query().Get("key") {
			return errors.New("key is missing")
		}

		return nil
	}

	func (self *User) View(ctx restapi.IContext) restapi.IOutput {
		return restapi.Output(true, "user view data", nil)
	}

	//visit http://www.v3.com:8888/user/xxxx
	//Response
	{
    	"data": null,
    	"errors": [
        	"key is missing"
    	],
    	"result": false
	}
	
#All available filters
	BeforeRun, BeforeView, BeforeList, BeforeUpdate, BeforeCreate, BeforeDelete, BeforeUpdateAll, BeforeDeleteAll

	
#API calls another API

user api create method calls content api view method.

	func (self *User) Create(ctx restapi.IContext) restapi.IOutput {
		query := ctx.Query()
		query.Add("with_content_comments", "1")
		return restapi.Call("content", restapi.MethodView, ctx)
	}
	
#Customized API response
You don't like my defualt response or even don't like to return json?

	//Default
	{
		"result": <bool>,
		"errors": []string,
		"data": interface{}
	}
	
You can use your customized response method
customized responseFunc, do it before restapi.Run

	restapi.Conf.ResponseFunc = func(output restapi.IOutput, ctx restapi.IContext, res http.ResponseWriter, req *http.Request) {
	 	res.Write([]byte("hello elson"))
	}
	
#Customized API RouterFunc
If you want to handle the request yourself, you can replace the default routerFunc.
In this example, all requests will call content view method

	restapi.Conf.RouterFunc = func(ctx restapi.IContext, res http.ResponseWriter, req *http.Request) (apiName, method string, ok bool) {
		return "content", restapi.MethodView, true
	}
		
#Customized API output
If you don't likt to use the default restapi.Output() method, you can return the struct which implements IOutput interface
	
	type IOutput interface {
		Result() bool
		Data() interface{}
		Errors() []string
	}	