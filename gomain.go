package main

import (
	"fmt"
	"os"

	"github.com/powellquiring/gopkg/fun"
)

var version = "v1"

func Main(params map[string]interface{}) map[string]interface{} {
	fmt.Println(version)

	// define the action response object
	response := make(map[string]interface{})

	// log all action parameters
	fmt.Printf("params: %v\n", params)
	fmt.Printf("environ: %v\n", os.Environ())

	// return statusCode, response headers and body
	response["statusCode"] = 418
	response["headers"] = map[string]interface{}{
		"Content-Type": "application/json;charset=utf-8",
		"Set-Cookie": [...]string{
			"UserID=Anonymous; Max-Age=3600; Version=",
			"SessionID=asdfgh123456; Path = /",
		},
	}
	response["body"] = map[string]interface{}{
		"params":  params,
		"environ": os.Environ(),
		"fun":     fun.FunString(),
	}

	return response
}

func main() {
	Main(nil)
}
