

# go-scout

## Intro

go-scout is a tool that allows you to control a Moorebot Scout robot from your computer (without using the mobile app). The robot is controlled using an XBOX controller, and video is displayed in a new window.

## Demo 

![go-scout demo](demo.gif)
## Installation

To install go-scout, simply download the latest release from the releases page and run the installer, or compile using the steps below.
```
go install github.com/shell-company/go-scout

OR 

git clone https://github.com/shell-company/go-scout && cd go-scout && go build
```


## Usage

To use go-scout, connect an XBOX controller to your computer and launch the application. The left stick is used for forward, reverse, and turning left or right. The right stick is used for strafing. The left bumper lowers the max speed, the right bumper raises max speed. Both bumpers pressed together tell the robot to stop moving. The start button exits the application.

```
Usage of ./scout:
  -h string
    	ROS endpoint such as IP_ADDRESS:PORT (default "192.168.1.224:11311")
  -v	verbose
  


  ╰─$ ./scout -h "192.168.1.225:11311" 
Starting go-scout controller for Moorebot Scout
ROS endpoint: 192.168.1.225:11311
2022/09/02 15:49:54 Please connect a joystick
```

#

## Considerations

Moorebot added a fair amount of bloatware to an otherwise great hardware platform. I took the following steps to increase privacy and reduce resource usage on the robot.

 **Follow the steps below at your own risk** as they may void warranties and potentially violate Moorebot terms of service.

- [ ]  SSH Access via `root:plt` to bot IP
```
ssh root@<SCOUT IP ADDRESS> 
```
- [ ]  Remove /opt/sockproxy/proxy_list.json
```
rm /opt/sockproxy/proxy_list.json
```
- [ ]  Disable sockproxy service
```
systemctl disable sockproxy.service
```
- [ ] Null route what appear to be backdoors in the CloudNode
```
route add 62.210.208.47 gw 127.0.0.1 lo
route add 45.35.33.24 gw 127.0.0.1 lo
route add 118.107.244.35 gw 127.0.0.1 lo
```

# 
## To-Do

- Add support for other controllers
- Create a proper Heads-Up-Display
- Add support for more features of the robot
    - Add battery status to HUD
    - Add compass to HUD
    - Add sensor data to HUD