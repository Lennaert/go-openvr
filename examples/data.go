package main

import (
	"fmt"
	openvr "github.com/lennaert/go-openvr"
	"log"
	"time"
)

var vrSystem *openvr.System
var leftController int
var rightController int

func main() {
	var err error
	vrSystem, err = openvr.Init()
	if err != nil {
		fmt.Printf("Cannot init OpenVR: %s", err.Error())
	}

	if vrSystem == nil {
		panic("vrSystem is nil")
	}

	leftController, rightController = FindControllerIds(0)
	log.Printf("[controllers] Found both controllers: left: %d, right: %d", leftController, rightController)

	ListenToEvents()
}

// FindControllerIds returns the device id's of the left and right controller
// The built-in `GetTrackedDeviceIndexForControllerRole` has been proven to be not working properly, so this will
// loop through the TrackedDevices until it finds both left & right, or panics after 10 retries (10 sec)
func FindControllerIds(retries int) (int, int) {
	var foundLeft = false
	var foundRight = false
	var left = 0
	var right = 0

	// check controller states
	for i := openvr.TrackedDeviceIndexHmd + 1; i < openvr.MaxTrackedDeviceCount; i++ {
		deviceClass := vrSystem.GetTrackedDeviceClass(int(i))
		if deviceClass == openvr.TrackedDeviceClassController {
			str, _ := vrSystem.GetStringTrackedDeviceProperty(int(i), openvr.PropModelNumberString)
			log.Printf("[controllers] Found %s as Controller class", str)

			role := vrSystem.GetControllerRoleForTrackedDeviceIndex(int(i))
			if role == openvr.TrackedControllerRoleLeftHand {
				log.Printf("[controllers] Found the left controller: %d", int(i))
				foundLeft = true
				left = int(i)
			}
			if role == openvr.TrackedControllerRoleRightHand {
				log.Printf("[controllers] Found the right controller: %d", int(i))
				foundRight = true
				right = int(i)
			}
		}
	}

	if foundLeft == false && foundRight == false {
		retries++
		if retries >= 10 {
			log.Panicf("[controllers] Cannot find controllers after 10 retries: left: %d, right: %d", left, right)
		}

		time.Sleep(1 * time.Second)
		log.Printf("[controllers] Controllers not found yet: left: %d, right: %d", left, right)
		return FindControllerIds(retries)
	}

	return left, right
}

// ListenToEvents is a blocking process that will continuously poll for events.
func ListenToEvents() {
	var event openvr.VREvent
	for {
		for vrSystem.PollNextEvent(&event) {
			ProccessVREvent(&event)
		}
		time.Sleep(200 * time.Millisecond)
	}
}

// ProccessVREvent This will process the event
func ProccessVREvent(event *openvr.VREvent) {
	switch event.EventType {
	case openvr.VREventTrackedDeviceActivated:
		fmt.Printf("Device %d attached.\n", event.TrackedDeviceIndex)
	case openvr.VREventTrackedDeviceDeactivated:
		fmt.Printf("Device %d detached.\n", event.TrackedDeviceIndex)
	case openvr.VREventTrackedDeviceUpdated:
		fmt.Printf("Device %d updated.\n", event.TrackedDeviceIndex)
	}
}
