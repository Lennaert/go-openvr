package main

import (
	"fmt"
	"github.com/Lennaert/go-openvr"
)

func main()  {
	vrSystem, err := openvr.Init()
	if err != nil {
		fmt.Printf("")
	}

	if vrSystem == nil {
		panic("vrSystem is nil")
	}

	deviceClass := vrSystem.GetTrackedDeviceClass(1)
	fmt.Printf("DeviceClass: %d", deviceClass)
}
