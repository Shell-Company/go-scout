package main

import (
	"fmt"
	"log"

	"github.com/aler9/goroslib"
	"github.com/aler9/goroslib/pkg/msg"
)

const (
	Status_PROCESS_OK             int8 = 0
	Status_PROCESS_ERROR          int8 = -1
	Status_OBJ_DETECT_CHARGE      int8 = 1
	Status_RECORD_START           int8 = 1
	Status_RECORD_STOP            int8 = 2
	Status_RECORD_ERROR           int8 = 3
	Status_P2P_AV_PLAYING         int8 = 1
	Status_P2P_AV_STOP            int8 = 2
	Status_P2P_AV_ERROR           int8 = 3
	Status_WIFI_MODE_AP           int8 = 0
	Status_WIFI_MODE_STA          int8 = 1
	Status_WIFI_STATUS_DISCONNECT int8 = 0
	Status_WIFI_STATUS_CONNECTED  int8 = 1
	Status_WIFI_STATUS_CONNECTING int8 = 2
	Status_WIFI_STATUS_WRONG_KEY  int8 = 3
	Status_WIFI_STATUS_CONN_FAIL  int8 = 4
	Status_WIFI_STATUS_STOP       int8 = 5
	Status_BACK_UP_DETECT         int8 = 1
	Status_BACK_UP_ALIGN          int8 = 2
	Status_BACK_UP_BACK           int8 = 3
	Status_BACK_UP_SUCCESS        int8 = 4
	Status_BACK_UP_FAIL           int8 = 5
	Status_BACK_UP_INACTIVE       int8 = 6
	Status_BACK_UP_CANCEL         int8 = 7
	Status_BACK_UP_REDETECT       int8 = 8
	Status_BATTERY_CHARGING       int8 = 0
	Status_BATTERY_UNCHARGE       int8 = 1
	Status_BATTERY_FULL           int8 = 2
	Status_BATTERY_UNKOWN         int8 = 3
)

type Status struct {
	msg.Package     `ros:"roller_eye"`
	msg.Definitions `ros:"int8 PROCESS_OK=0,int8 PROCESS_ERROR=-1,int8 OBJ_DETECT_CHARGE=1,int8 RECORD_START=1,int8 RECORD_STOP=2,int8 RECORD_ERROR=3,int8 P2P_AV_PLAYING=1,int8 P2P_AV_STOP=2,int8 P2P_AV_ERROR=3,int8 WIFI_MODE_AP=0,int8 WIFI_MODE_STA=1,int8 WIFI_STATUS_DISCONNECT=0,int8 WIFI_STATUS_CONNECTED=1,int8 WIFI_STATUS_CONNECTING=2,int8 WIFI_STATUS_WRONG_KEY=3,int8 WIFI_STATUS_CONN_FAIL=4,int8 WIFI_STATUS_STOP=5,int8 BACK_UP_DETECT=1,int8 BACK_UP_ALIGN=2,int8 BACK_UP_BACK=3,int8 BACK_UP_SUCCESS=4,int8 BACK_UP_FAIL=5,int8 BACK_UP_INACTIVE=6,int8 BACK_UP_CANCEL=7,int8 BACK_UP_REDETECT=8,int8 BATTERY_CHARGING=0,int8 BATTERY_UNCHARGE=1,int8 BATTERY_FULL=2,int8 BATTERY_UNKOWN=3"`
	Status          []int32
}

func checkBattery() {
	for {
		// create a subscriber
		sub, err := goroslib.NewSubscriber(goroslib.SubscriberConf{
			Node:  n,
			Topic: "/CoreNode/SensorNode/simple_battery_status",
			Callback: func(msg *Status) {
				fmt.Println(msg.Status)
			},
		})
		if err != nil {
			log.Fatal(err)
		}
		// wait for a message
		defer sub.Close()
	}
}
