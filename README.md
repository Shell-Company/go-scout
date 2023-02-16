

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


Note: This version will not try to automatically determine your local IP address. You must specify it using the `-l` flag. This was to get around a restriction  where the robot will not connect to the computer if the IP address is not in the same subnet as the robot. 

Using the -l flag you are now able to connect to a remotely hosted ROS endpoint. This is useful if you want to run the robot from a remote location and control it from your computer.


```

## Controls
-c joystick

| Button | Action |
| --- | --- |
| Left Stick | Forward, Reverse, Left, Right |
| Right Stick | Strafe |
| Left Bumper | Lower Max Speed |
| Right Bumper | Raise Max Speed |
| Both Bumpers | Stop |
| Start | Exit |
| Keyboard 0 | -Night Vision brightness |
| Keyboard 9 | +Night Vision brightness |
| Keyboard Space | Save image to disk |
| Keyboard h | Return to charging station |

-c keyboard

| Button | Action |
| --- | --- |
| W | Forward |
| S | Reverse |
| A | Left |
| D | Right |
| Q | Strafe Left |
| E | Strafe Right |
| H | Return to charging station |
| Space | Screenshot |
|  0 | -Night Vision brightness |
|  9 | +Night Vision brightness |
| Left Shift | Lower Max Speed |
| Left Ctrl | Raise Max Speed |
| Esc | Exit |

```
Usage of ./scout:
  -c string
    	control scheme, keyboard or joystick (default "keyboard")
  -h string
    	ROS endpoint such as IP_ADDRESS:PORT (default "192.168.1.224:11311")
  -l string
    	localhost address (default "192.168.1.211")
  -v	verbose
  -windowX int
    	window width (default 1920)
  -windowY int
    	window height (default 1080)


  ╰─$ ./scout -h "192.168.1.225:11311" -l 192.168.1.211 -c joystick
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


- Create a proper Heads-Up-Display
- Add support for more features of the robot
    - Add battery status to HUD
    - Add compass to HUD
    - Add sensor data to HUD