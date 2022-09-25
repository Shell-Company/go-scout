package main

import (
	"log"

	"github.com/aler9/goroslib"
	"github.com/aler9/goroslib/pkg/msg"
)

type goHomeRequest struct {
}

type goHomeResponse struct {
}

type goHomeService struct {
	msg.Package `ros:"/nav_low_bat"`
	Request     goHomeRequest
	Response    goHomeResponse
}

// call the /CoreNode/adjust_light service with a value of 1 to turn on the light
func scoutGoHome() {
	// create a node and connect to the master
	n, err := goroslib.NewNode(goroslib.NodeConf{
		Name:          "scout-go-home",
		MasterAddress: *flagROSHostAddress,
	})
	if err != nil {
		panic(err)
	}
	defer n.Close()
	// create a service client
	cl, err := goroslib.NewServiceClient(goroslib.ServiceClientConf{
		Node: n,
		Name: "/nav_low_bat",
		Srv: &goHomeService{
			Request: goHomeRequest{},
		}})
	if err != nil {
		log.Println("An error occured while trying to send the bot home", err)
	}
	defer cl.Close()
	// call the service
	req := goHomeRequest{}
	res := goHomeResponse{}

	err = cl.Call(&req, &res)
	if err != nil {
		panic(err)
	}
	log.Println("Command received: Scout returning home in 5 seconds")
}
