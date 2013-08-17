package main

import (
	"fmt"
	"github.com/elsonwu/restapi"
	"net/url"
)

type User struct {
}

func (self *User) GET(params restapi.Params) restapi.Output {
	fmt.Println("User get ", params)
	// output := restapi.Output{}
	// output.Result = true
	// output.Data = map[string]interface{}{"hello": "elson"}
	// return output

	url := url.URL{}
	query := url.Query()
	query.Add("with_content_user", "1")
	query.Add("with_content_comments", "1")
	params2 := restapi.Params{Query: query}
	return restapi.Call("content", "GET", params2)
}

func (self *User) POST(params restapi.Params) restapi.Output {
	fmt.Println("User post ", params)
	output := restapi.Output{}
	output.Result = true
	output.Data = map[string]interface{}{"hello": "elson"}
	return output
}

type Content struct {
}

func (self *Content) GET(params restapi.Params) restapi.Output {
	fmt.Println("Content get ", params)
	output := restapi.Output{}
	output.Result = true
	output.Data = map[string]interface{}{"hello": "elson"}
	return output
}

func (self *Content) POST(params restapi.Params) restapi.Output {
	fmt.Println("Content post ", params)
	output := restapi.Output{}
	output.Result = true
	output.Data = map[string]interface{}{"hello": "elson"}
	return output
}

func main() {
	restapi.Add("user", &User{})
	restapi.Add("content", &Content{})
	restapi.Run()
}
