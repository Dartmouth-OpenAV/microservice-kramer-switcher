package main

import (
	"errors"
	"fmt"
	"strconv"
	"strings"

	"github.com/Dartmouth-OpenAV/microservice-framework/framework"
)

func vsDoSetVideoRoute(socketKey string, output string, input string) (string, error) {
	function := "vsDoSetVideoRoute"
	baseStr := "X-ROUTE"
	noQuotesInput := strings.ReplaceAll(input, `"`, "")

	var outputPortType string
	var inputPortType string
	switch output {
	case "1", "2":
		outputPortType = "HDMI"
	case "3", "4":
		outputPortType = "HDBT"
	default:
		errMsg := fmt.Sprintf(function + " - unrecognized output value: " + output)
		framework.AddToErrors(socketKey, errMsg)
		return output, errors.New(errMsg)
	}
	switch noQuotesInput {
	case "1", "2", "3", "4", "5", "6":
		inputPortType = "HDMI"
	case "7", "8":
		inputPortType = "HDBT"
	default:
		errMsg := fmt.Sprintf(function + " - unrecognized input value: " + noQuotesInput)
		framework.AddToErrors(socketKey, errMsg)
		return noQuotesInput, errors.New(errMsg)
	}
	// Ex Command: #X-ROUTE OUT.HDMI.3.VIDEO.1,IN.HDMI.2.VIDEO.1
	outputCommandStr := "OUT." + outputPortType + "." + output + ".VIDEO.1"
	inputCommandStr := "IN." + inputPortType + "." + noQuotesInput + ".VIDEO.1"
	commandStr := []byte("#" + baseStr + " " + outputCommandStr + "," + inputCommandStr)

	resp, err := sendCommand(socketKey, commandStr, baseStr)
	if err != nil {
		errMsg := fmt.Sprintf(function+" - error sending command %v", err.Error())
		framework.AddToErrors(socketKey, errMsg)
		return string(resp), errors.New(errMsg)
	}

	// Now we need to parse the response from the Kramer
	// Ex Response: #X-ROUTE OUT.HDMI.3.VIDEO.1,IN.HDMI.2.VIDEO.1
	// Parse and verify the value we got back
	respVals := strings.Split(string(resp), " ")    // Kramer delimits responses with spaces
	respSections := strings.Split(respVals[1], ",") // Kramer delimits the video route part of the response with commas
	if len(respSections) != 2 {
		errMsg := fmt.Sprintf(function+" - 445dfc wrong number of tokens in response: %s", string(resp))
		framework.AddToErrors(socketKey, errMsg)
		return string(resp), errors.New(errMsg)
	}
	// Ex: OUT.HDMI.3.VIDEO.1
	outputResp := strings.Split(respSections[0], ".")
	if len(outputResp) != 5 {
		errMsg := fmt.Sprintf(function+" - 5n4dm4 wrong number of tokens in output response: %s", string(respSections[0]))
		framework.AddToErrors(socketKey, errMsg)
		return string(resp), errors.New(errMsg)
	}
	outputInt := outputResp[2]

	// Ex: IN.HDMI.2.VIDEO.1
	inputResp := strings.Split(respSections[1], ".")
	if len(inputResp) != 5 {
		errMsg := fmt.Sprintf(function+" - 0di4do wrong number of tokens in input response: %s", string(respSections[1]))
		framework.AddToErrors(socketKey, errMsg)
		return string(resp), errors.New(errMsg)
	}
	inputInt := inputResp[2]

	if outputInt != output {
		errMsg := fmt.Sprintf(function+" - output returned: %s doesn't match output specified: %s\n", outputInt, output)
		framework.AddToErrors(socketKey, errMsg)
		return string(resp), errors.New(errMsg)
	}
	if `"`+inputInt+`"` != input { // add quotes because the input and the cache store values in JSON
		errMsg := fmt.Sprintf(function+" - input returned: %s doesn't match input set: %s\n", `"`+inputInt+`"`, input)
		framework.AddToErrors(socketKey, errMsg)
		return string(resp), errors.New(errMsg)
	}

	returnStr := `"` + inputInt + `"`
	return returnStr, nil
}

func vsDoSetVideoMute(socketKey, output string, value string) (string, error) {
	function := "vsDoSetVideoMute"
	onoff := "not set"

	switch value {
	case `"true"`:
		onoff = "ON" // Disable video in Kramer VMUTE command
	case `"false"`:
		onoff = "OFF" // Enable video
	default:
		errMsg := fmt.Sprintf(function + " - unrecognized videomute value: " + value)
		framework.AddToErrors(socketKey, errMsg)
		return value, errors.New(errMsg)
	}

	baseStr := "X-MUTE"
	var portType string
	switch output {
	case "1", "2":
		portType = "HDMI"
	case "3", "4":
		portType = "HDBT"
	default:
		errMsg := fmt.Sprintf(function + " - unrecognized output value: " + output)
		framework.AddToErrors(socketKey, errMsg)
		return output, errors.New(errMsg)
	}
	// Ex Command: #X-MUTE OUT.HDMI.4.VIDEO.1,ON
	commandStr := []byte("#" + baseStr + " OUT." + portType + "." + output + ".VIDEO.1," + onoff)
	resp, err := sendCommand(socketKey, commandStr, baseStr)

	if err != nil {
		errMsg := fmt.Sprintf(function+" - error sending command %v", err.Error())
		framework.AddToErrors(socketKey, errMsg)
		return string(resp), errors.New(errMsg)
	}
	// Parse the value we got back
	// Ex Response: #X-MUTE OUT.HDMI.4.VIDEO.1,ON
	respVals := strings.Split(string(resp), " ") // Kramer delimits responses with spaces
	respMute := strings.Split(respVals[1], ",")  // Kramer delimits the mute part of the response with commas
	if len(respMute) != 2 {
		errMsg := fmt.Sprintf(function+" - bbfsd4 wrong number of tokens in response: %s", string(resp))
		framework.AddToErrors(socketKey, errMsg)
		return string(resp), errors.New(errMsg)
	}
	state := respMute[1]
	switch state {
	case "OFF":
		value = "false"
	case "ON":
		value = "true"
	default: // not a legal value
		errMsg := fmt.Sprintf(function + " - unrecognized response to mute command: " + state)
		framework.AddToErrors(socketKey, errMsg)
		return string(resp), errors.New(errMsg)
	}

	// If we got here, the response was good, so successful return with the state indication

	return value, nil
}

func vsDoGetVideoRoute(socketKey string, output string) (string, error) {
	function := "vsDoGetVideoRoute"
	baseStr := "X-ROUTE"

	var portType string
	switch output {
	case "1", "2":
		portType = "HDMI"
	case "3", "4":
		portType = "HDBT"
	default:
		errMsg := fmt.Sprintf(function + " - unrecognized output value: " + output)
		framework.AddToErrors(socketKey, errMsg)
		return output, errors.New(errMsg)
	}

	// #X-ROUTE? OUT.HDMI.3.VIDEO.1
	commandStr := []byte("#" + baseStr + "? OUT." + portType + "." + output + ".VIDEO.1")
	resp, err := sendCommand(socketKey, commandStr, baseStr)
	if err != nil {
		errMsg := fmt.Sprintf(function+" - error sending command %v", err.Error())
		framework.AddToErrors(socketKey, errMsg)
		return string(resp), errors.New(errMsg)
	}
	// Ex Response: ~01@X-ROUTE OUT.HDMI.1.VIDEO.1,IN.HDMI.1.VIDEO.1
	respVals := strings.Split(string(resp), " ") // Kramer delimits responses with spaces
	respRoute := strings.Split(respVals[1], ",") // Kramer delimits the volume part of the response with commas
	if len(respRoute) != 2 {
		errMsg := fmt.Sprintf(function+" - sd23q4gd wrong number of tokens in response: %s", string(resp))
		framework.AddToErrors(socketKey, errMsg)
		return string(resp), errors.New(errMsg)
	}
	// respRoute = [ OUT.HDMI.1.VIDEO.1 IN.HDMI.1.VIDEO.1 ]
	outputString := respRoute[0]
	inputString := respRoute[1]

	outputStringParts := strings.Split(outputString, ".")
	inputStringParts := strings.Split(inputString, ".")

	_, err = strconv.Atoi(string(outputStringParts[2]))
	intIn, err2 := strconv.Atoi(string(inputStringParts[2]))

	if err != nil || err2 != nil { // Illegal input or output
		errMsg := function + " - Invalid input or output value in device response received: " +
			string(outputStringParts[2]) + " " + string(inputStringParts[2])
		framework.AddToErrors(socketKey, errMsg)
		return string(resp), errors.New(errMsg)
	}

	returnStr := `"` + strconv.Itoa(intIn) + `"`
	return returnStr, nil
}

func vsDoGetVideoMute(socketKey string, output string) (string, error) {
	function := "vsDoGetVideoMute"

	baseStr := "X-MUTE"
	var portType string
	switch output {
	case "1", "2":
		portType = "HDMI"
	case "3", "4":
		portType = "HDBT"
	default:
		errMsg := fmt.Sprintf(function + " - unrecognized output value: " + output)
		framework.AddToErrors(socketKey, errMsg)
		return output, errors.New(errMsg)
	}
	// #X-MUTE? OUT.HDMI.4.VIDEO.1
	commandStr := []byte("#" + baseStr + "? OUT." + portType + "." + output + ".VIDEO.1")
	resp, err := sendCommand(socketKey, commandStr, baseStr)
	framework.Log(fmt.Sprintf(function + " - got resp: " + string(resp) + "\n"))
	if err != nil {
		errMsg := fmt.Sprintf(function+" - error sending command %v", err.Error())
		framework.AddToErrors(socketKey, errMsg)
		return string(resp), errors.New(errMsg)
	}
	// Parse the value we got back
	respVals := strings.Split(string(resp), " ") // Kramer delimits responses with spaces
	respMute := strings.Split(respVals[1], ",")  // Kramer delimits the mute part of the response with commas
	if len(respMute) != 2 {
		errMsg := fmt.Sprintf(function+" - 54fsav wrong number of tokens in response: %s", string(resp))
		framework.AddToErrors(socketKey, errMsg)
		return string(resp), errors.New(errMsg)
	}

	value := "unknown"
	state := respMute[1]
	switch state {
	case "OFF":
		value = "false"
	case "ON":
		value = "true"
	default: // not a legal value
		errMsg := function + " - get video mute returned: " + state + " which is not a legal value\n"
		framework.AddToErrors(socketKey, errMsg)
		return errMsg, errors.New(errMsg)
	}

	return `"` + value + `"`, nil
}
