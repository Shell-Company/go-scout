package main

import (
	"bytes"
	"flag"
	"fmt"
	_ "image/jpeg"
	_ "image/png"
	"log"
	"math"
	"os"
	"os/signal"
	"time"

	"github.com/aler9/goroslib"
	"github.com/aler9/goroslib/pkg/msg"
	"github.com/aler9/goroslib/pkg/msgs/geometry_msgs"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"github.com/simulatedsimian/joystick"
)

var (
	cameraData         = make(chan []byte)
	controlData        string
	ROSHostAddress     = "192.168.1.224:11311"
	WindowX            = 1920
	WindowY            = 1080
	flagROSHostAddress = flag.String("h", ROSHostAddress, "ROS endpoint such as IP_ADDRESS:PORT")
	flagVerbose        = flag.Bool("v", false, "verbose")
	joystickLeftX      float64
	joystickLeftY      float64
	joystickRightX     float64
	joystickRightY     float64
	pub                *goroslib.Publisher
	forwardSpeed       = .2
	maxForwardSpeed    = 1.0
	minForwardSpeed    = 0.1
)

const (
	Frame_VIDEO_STREAM_H264 int8 = 0
	Frame_VIDEO_STREAM_JPG  int8 = 1
	Frame_AUDIO_STREAM_AAC  int8 = 2
)

func (g *Game) Draw(screen *ebiten.Image) {
	// retrieve the image from the channel
	imageFromCamera := <-cameraData
	// convert to image to an ebiten image
	cameraFeed, _, err := ebitenutil.NewImageFromReader(bytes.NewReader(imageFromCamera))
	if err != nil {
		fmt.Println(err)
	}
	// draw the image on the screen
	screen.DrawImage(cameraFeed, nil)
	// print some HUD data
	ebitenutil.DebugPrint(screen, controlData)

}
func main() {
	flag.Parse()

	// Start up messages
	fmt.Println("Starting go-scout controller for Moorebot Scout")
	fmt.Println("ROS endpoint:", *flagROSHostAddress)

	// listen for ctrl+c
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	go func() {
		sig := <-c
		log.Fatal("Received", sig, "signal")
	}()

	// create a node
	if *flagVerbose {
		log.Println(fmt.Sprintln("creating camera access node using %s", *flagROSHostAddress))
	}
	n, err := goroslib.NewNode(goroslib.NodeConf{
		Name:          "scout-camera-access",
		MasterAddress: *flagROSHostAddress,
	})
	if err != nil {
		panic(err)
	}
	defer n.Close()

	// create a subscriber
	subby := goroslib.SubscriberConf{
		Node:      n,
		Topic:     "/CoreNode/jpg",
		Callback:  onMessageFrame,
		QueueSize: 1,
	}

	sub, err := goroslib.NewSubscriber(subby)
	if err != nil {
		panic(err)
	}
	defer sub.Close()

	// init robo controller
	go robotControl()

	ebiten.SetWindowSize(WindowX, WindowY)
	ebiten.SetWindowTitle("🥷 Scout 🤖")
	ebiten.SetWindowResizable(true)
	if err := ebiten.RunGame(&Game{}); err != nil {
		log.Fatal(err)
	}
}

// onMessageFrame is called when a message is received from the CoreNode/jpg topic (camera)
func onMessageFrame(msg *Frame) {
	// write camera data to channel
	cameraData <- msg.Data

}

// robotControl is a goroutine that will read the joystick and publish the control data to the robot
func robotControl() {
	flag.Parse()
	ROSHostAddress = *flagROSHostAddress

	// create a node and connect to the master
	n, err := goroslib.NewNode(goroslib.NodeConf{
		Name:          "scout-controller",
		MasterAddress: *flagROSHostAddress,
	})
	if err != nil {
		panic(err)
	}
	defer n.Close()

	// Print gamepad status and pos
	jsid := 0
	js, err := joystick.Open(jsid)
	if err != nil {
		log.Fatal("Please connect a joystick	")

	}

	for {
		state, err := js.Read()
		if err != nil {
			panic(err)
		}
		// -32767 to 32768
		joystickLeftX = squashToFloat(state.AxisData[0])
		joystickLeftY = squashToFloat(state.AxisData[1]) * -1
		joystickRightX = squashToFloat(state.AxisData[2])
		joystickRightY = squashToFloat(state.AxisData[3]) * -1
		if joystickLeftY == -0.0 {
			joystickLeftY = 0
		}
		if joystickRightY == -0.0 {
			joystickRightY = 0
		}

		// write message to channel
		controlData = fmt.Sprintf("lX: %f lY: %f rX:%f rY:%f speedModifier:%f", joystickLeftX, joystickLeftY, joystickRightX, joystickRightY, forwardSpeed)
		// forward speed to hud data
		buttonPress := state.Buttons
		if *flagVerbose {
			fmt.Println("Left Stick:", joystickLeftX, joystickLeftY, "\n")
			fmt.Println("Right Stick:", joystickRightX, joystickRightY, "\n")
			fmt.Println("Button Press value", buttonPress, forwardSpeed)
		}

		defer js.Close()

		if buttonPress == 2048 {
			os.Exit(0)
		}
		// right bumper
		if buttonPress == 128 {
			forwardSpeed = forwardSpeed + .1
		}
		// left bumper
		if buttonPress == 64 {
			forwardSpeed = forwardSpeed - .1
			if forwardSpeed < 0 {
				forwardSpeed = 0
			}

		}
		// hold both bumpers to halt the robot
		if buttonPress == 192 {
			forwardSpeed = 0
			joystickRightX = 0
		}
		if forwardSpeed > maxForwardSpeed {
			forwardSpeed = 1.0
		}
		if forwardSpeed < minForwardSpeed {
			forwardSpeed = .1
		}
		msg := &geometry_msgs.Twist{
			Linear: geometry_msgs.Vector3{
				X: joystickRightX * .2,          // strafe l r
				Y: joystickLeftY * forwardSpeed, // move fwd/back
			},
			Angular: geometry_msgs.Vector3{
				Z: joystickLeftX * -2.9, // rotate l r
			},
		}

		// create a publisher
		pub, err = goroslib.NewPublisher(goroslib.PublisherConf{
			Node:  n,
			Topic: "/cmd_vel",
			Msg:   msg,
			Latch: true,
		})
		if err != nil {
			log.Fatal(err)
		}
		// send the message to the robot
		pub.Write(msg)
		// wait a bit
		time.Sleep(time.Millisecond * 600)
		pub.Close()
	}
}

func (g *Game) Update() error {
	return nil
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return WindowX, WindowY
}

// squashToFloat converts a signed 16 bit integer to a float between -1 and 1
func squashToFloat(n int) (f float64) {
	if n < -32767 {
		f = -1.0
	} else if n > 32768 {
		f = 1.0
	} else {
		value := float64(n) / 32768.0
		if math.Abs(value) < 0.015 {
			f = 0.0
		} else {
			f = value
		}
	}
	return f * .5
}

type Game struct{}

// Frame is a struct that holds the image data from the camera topic roller_bot/frame
type Frame struct {
	msg.Package     `ros:"roller_eye"`
	msg.Definitions `ros:"int8 VIDEO_STREAM_H264=0,int8 VIDEO_STREAM_JPG=1,int8 AUDIO_STREAM_AAC=2"`
	Seq             uint32
	Stamp           uint64
	Session         uint32
	Type            int8
	Oseq            uint32
	Par1            int32
	Par2            int32
	Par3            int32
	Par4            int32
	Data            []uint8
}
