package main

import (
	"errors"

	"github.com/Dartmouth-OpenAV/microservice-framework/framework"
)

//deviceStates[socketKey]["videomute" + output] = `"false"`  // unmuted

func setFrameworkGlobals() {
	// Define device specific globals
	framework.DefaultSocketPort = 5000 // Kramer's default socket port
	framework.GlobalDelimiter = 0xA    // Linefeed is Kramer's line delimiter for socket commands
	framework.MicroserviceName = "OpenAV Kramer Video Switcher Microservice"

	// globals that change modes in the microservice framework:
	framework.CheckFunctionAppendBehavior = "Remove older instance"
	framework.KeepAlive = false // Kramer devices don't always fare well if we keep the connection open too long

	// Kramer switches unmute video when switching inputs, so now we need to update the videomute cache value so
	// our system catches up with that change quickly.  We do this by simply setting the cache value to what we know
	// it will be.
	framework.RelatedActions["videoroute"] = make(map[string]string)
	framework.RelatedActions["videoroute"]["endPoint"] = "videomute"
	framework.RelatedActions["videoroute"]["value"] = `"` + "false" + `"`
	framework.RegisterMainGetFunc(doDeviceSpecificGet)
	framework.RegisterMainSetFunc(doDeviceSpecificSet)
}

// Every microservice using this golang microservice framework needs to provide this function to invoke functions to do sets.
// socketKey is the network connection for the framework to use to communicate with the device.
// setting is the first parameter in the URI.
// arg1 are the second and third parameters in the URI.
//
//	  Example PUT URIs that will result in this function being invoked:
//		 ":address/:setting/"
//	  ":address/:setting/:arg1"
//	  ":address/:setting/:arg1/:arg2"
func doDeviceSpecificSet(socketKey string, setting string, arg1 string, arg2 string, arg3 string) (string, error) {
	function := "doDeviceSpecificSet"

	// Add a case statement for each set function your microservice implements.  These calls can use 0, 1, or 2 arguments.
	switch setting {
	case "volume":
		return setVolume(socketKey, arg1, arg2)
	case "videoroute":
		return setVideoRoute(socketKey, arg1, arg2)
	case "audioandvideoroute":
		return setAudioAndVideoRoute(socketKey, arg1, arg2)
	case "audiomute":
		return setAudioMute(socketKey, arg1, arg2)
	case "videomute":
		return setVideoMute(socketKey, arg1, arg2)
	}

	// If we get here, we didn't recognize the setting.  Send an error back to the config writer who had a bad URL.
	errMsg := function + " - unrecognized setting in SET URI: " + setting
	framework.AddToErrors(socketKey, errMsg)
	err := errors.New(errMsg)
	return setting, err
}

// Every microservice using this golang microservice framework needs to provide this function to invoke functions to do gets.
// socketKey is the network connection for the framework to use to communicate with the device.
// setting is the first parameter in the URI.
// arg1 are the second and third parameters in the URI.
//
//	  Example GET URIs that will result in this function being invoked:
//		 ":address/:setting/"
//	  ":address/:setting/:arg1"
//	  ":address/:setting/:arg1/:arg2"
//
// Every microservice using this golang microservice framework needs to provide this function to invoke functions to do gets.
func doDeviceSpecificGet(socketKey string, setting string, arg1 string, arg2 string) (string, error) {
	function := "doDeviceSpecificGet"

	switch setting {
	case "volume":
		return getVolume(socketKey, arg1) // arg1 is name in this case
	case "videoroute":
		return getVideoRoute(socketKey, arg1) // arg1 is name in this case
	case "audioandvideoroute":
		return getAudioAndVideoRoute(socketKey, arg1)
	case "audiomute":
		return getAudioMute(socketKey, arg1) // arg1 is name in this case
	case "videomute":
		return getVideoMute(socketKey, arg1)
	case "videoinputstatus":
		return getVideoInputStatus(socketKey, arg1)
	case "occupancystatus":
		temp, err := getOccupancyStatus(socketKey, arg1)
		return temp, err
		// return getOccupancyStatus(socketKey, arg1)
	case "healthcheck":
		return healthCheck(socketKey)
	}

	// If we get here, we didn't recognize the setting.  Send an error back to the config writer who had a bad URL.
	errMsg := function + " - unrecognized setting in GET URI: " + setting
	framework.AddToErrors(socketKey, errMsg)
	err := errors.New(errMsg)
	return setting, err
}

func main() {
	setFrameworkGlobals()
	framework.Startup()
}
