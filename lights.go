package main

import (
	"fmt"

	"github.com/aler9/goroslib"
	"github.com/aler9/goroslib/pkg/msg"
)

type adjustLightRequest struct {
	Cmd int32 `ros:"cmd int32"`
}

type adjustLightResponse struct {
}

type adjustLightService struct {
	msg.Package `ros:"/CoreNode/adjust_light"`
	Request     adjustLightRequest
	Response    adjustLightResponse
}

// call the /CoreNode/adjust_light service with a value of 1 to turn on the light
func turnOnLight(lightValue int32) {

	// create a service client
	cl, err := goroslib.NewServiceClient(goroslib.ServiceClientConf{
		Node: n,
		Name: "/CoreNode/adjust_light",
		Srv: &adjustLightService{
			Request: adjustLightRequest{
				Cmd: lightValue,
			},
		}})
	if err != nil {
		fmt.Println("Light is already on")
	}
	defer cl.Close()

	// call the service
	req := adjustLightRequest{
		Cmd: lightValue,
	}
	res := adjustLightResponse{}

	err = cl.Call(&req, &res)
	if err != nil {
		fmt.Println(err)
	}
}
