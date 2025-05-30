#!/bin/bash

# Kramer Switcher Microservice API Test Script
# This script tests all endpoints from the Kramer Switcher Microservice Postman collection

# Configuration variables - Update these values as needed
MICROSERVICE_URL="localhost:8080"
DEVICE_FQDN="kramer-switcher.local"
VOLUME_CHANNEL="1000"
OUTPUT_CHANNEL="1"
INPUT_CHANNEL_1="1"
INPUT_CHANNEL_2="2"
INPUT_CHANNEL_3="3"

echo "Starting Kramer Switcher Microservice API Tests..."
echo "Microservice URL: $MICROSERVICE_URL"
echo "Device FQDN: $DEVICE_FQDN"
echo "Volume Channel: $VOLUME_CHANNEL"
echo "Output Channel: $OUTPUT_CHANNEL"
echo "Input Channels: $INPUT_CHANNEL_1, $INPUT_CHANNEL_2, $INPUT_CHANNEL_3"
echo "=============================================="

# GET Volume
echo "Testing GET Volume for Channel $VOLUME_CHANNEL..."
curl -X GET "http://$MICROSERVICE_URL/$DEVICE_FQDN/volume/$VOLUME_CHANNEL"
sleep 1

# GET Videoroute
echo "Testing GET Videoroute for Output $OUTPUT_CHANNEL..."
curl -X GET "http://$MICROSERVICE_URL/$DEVICE_FQDN/videoroute/$OUTPUT_CHANNEL"
sleep 1

# GET Audioandvideoroute
echo "Testing GET Audioandvideoroute for Output $OUTPUT_CHANNEL..."
curl -X GET "http://$MICROSERVICE_URL/$DEVICE_FQDN/audioandvideoroute/$OUTPUT_CHANNEL"
sleep 1

# GET Audiomute
echo "Testing GET Audiomute for Channel $OUTPUT_CHANNEL..."
curl -X GET "http://$MICROSERVICE_URL/$DEVICE_FQDN/audiomute/$OUTPUT_CHANNEL"
sleep 1

# GET Videomute
echo "Testing GET Videomute for Channel $OUTPUT_CHANNEL..."
curl -X GET "http://$MICROSERVICE_URL/$DEVICE_FQDN/videomute/$OUTPUT_CHANNEL"
sleep 1

# GET Videoinputstatus
echo "Testing GET Videoinputstatus for Input $INPUT_CHANNEL_1..."
curl -X GET "http://$MICROSERVICE_URL/$DEVICE_FQDN/videoinputstatus/$INPUT_CHANNEL_1"
sleep 1

# GET Occupancystatus
echo "Testing GET Occupancystatus for Channel $OUTPUT_CHANNEL..."
curl -X GET "http://$MICROSERVICE_URL/$DEVICE_FQDN/occupancystatus/$OUTPUT_CHANNEL"
sleep 1

echo "=============================================="
echo "Starting SET/PUT operations..."
echo "=============================================="

# SET Volume - Test different volume levels
echo "Testing SET Volume for Channel $VOLUME_CHANNEL (75)..."
curl -X PUT "http://$MICROSERVICE_URL/$DEVICE_FQDN/volume/$VOLUME_CHANNEL" \
     -H "Content-Type: application/json" \
     -d "\"75\""
sleep 1

echo "Testing SET Volume for Channel $VOLUME_CHANNEL (50)..."
curl -X PUT "http://$MICROSERVICE_URL/$DEVICE_FQDN/volume/$VOLUME_CHANNEL" \
     -H "Content-Type: application/json" \
     -d "\"50\""
sleep 1

echo "Testing SET Volume for Channel $VOLUME_CHANNEL (25)..."
curl -X PUT "http://$MICROSERVICE_URL/$DEVICE_FQDN/volume/$VOLUME_CHANNEL" \
     -H "Content-Type: application/json" \
     -d "\"25\""
sleep 1

# SET Videoroute - Test different input routing
echo "Testing SET Videoroute for Output $OUTPUT_CHANNEL to Input $INPUT_CHANNEL_1..."
curl -X PUT "http://$MICROSERVICE_URL/$DEVICE_FQDN/videoroute/$OUTPUT_CHANNEL" \
     -H "Content-Type: application/json" \
     -d "\"$INPUT_CHANNEL_1\""
sleep 1

echo "Testing SET Videoroute for Output $OUTPUT_CHANNEL to Input $INPUT_CHANNEL_2..."
curl -X PUT "http://$MICROSERVICE_URL/$DEVICE_FQDN/videoroute/$OUTPUT_CHANNEL" \
     -H "Content-Type: application/json" \
     -d "\"$INPUT_CHANNEL_2\""
sleep 1

echo "Testing SET Videoroute for Output $OUTPUT_CHANNEL to Input $INPUT_CHANNEL_3..."
curl -X PUT "http://$MICROSERVICE_URL/$DEVICE_FQDN/videoroute/$OUTPUT_CHANNEL" \
     -H "Content-Type: application/json" \
     -d "\"$INPUT_CHANNEL_3\""
sleep 1

# SET Audioandvideoroute - Test different input routing
echo "Testing SET Audioandvideoroute for Output $OUTPUT_CHANNEL to Input $INPUT_CHANNEL_1..."
curl -X PUT "http://$MICROSERVICE_URL/$DEVICE_FQDN/audioandvideoroute/$OUTPUT_CHANNEL" \
     -H "Content-Type: application/json" \
     -d "\"$INPUT_CHANNEL_1\""
sleep 1

echo "Testing SET Audioandvideoroute for Output $OUTPUT_CHANNEL to Input $INPUT_CHANNEL_2..."
curl -X PUT "http://$MICROSERVICE_URL/$DEVICE_FQDN/audioandvideoroute/$OUTPUT_CHANNEL" \
     -H "Content-Type: application/json" \
     -d "\"$INPUT_CHANNEL_2\""
sleep 1

echo "Testing SET Audioandvideoroute for Output $OUTPUT_CHANNEL to Input $INPUT_CHANNEL_3..."
curl -X PUT "http://$MICROSERVICE_URL/$DEVICE_FQDN/audioandvideoroute/$OUTPUT_CHANNEL" \
     -H "Content-Type: application/json" \
     -d "\"$INPUT_CHANNEL_3\""
sleep 1

# SET Audiomute - Test mute on/off
echo "Testing SET Audiomute for Channel $OUTPUT_CHANNEL (true)..."
curl -X PUT "http://$MICROSERVICE_URL/$DEVICE_FQDN/audiomute/$OUTPUT_CHANNEL" \
     -H "Content-Type: application/json" \
     -d "\"true\""
sleep 1

echo "Testing SET Audiomute for Channel $OUTPUT_CHANNEL (false)..."
curl -X PUT "http://$MICROSERVICE_URL/$DEVICE_FQDN/audiomute/$OUTPUT_CHANNEL" \
     -H "Content-Type: application/json" \
     -d "\"false\""
sleep 1

# SET Videomute - Test mute on/off
echo "Testing SET Videomute for Channel $OUTPUT_CHANNEL (true)..."
curl -X PUT "http://$MICROSERVICE_URL/$DEVICE_FQDN/videomute/$OUTPUT_CHANNEL" \
     -H "Content-Type: application/json" \
     -d "\"true\""
sleep 1

echo "Testing SET Videomute for Channel $OUTPUT_CHANNEL (false)..."
curl -X PUT "http://$MICROSERVICE_URL/$DEVICE_FQDN/videomute/$OUTPUT_CHANNEL" \
     -H "Content-Type: application/json" \
     -d "\"false\""
sleep 1

echo "=============================================="
echo "Testing additional input status checks..."
echo "=============================================="

# GET Videoinputstatus for multiple inputs
echo "Testing GET Videoinputstatus for Input $INPUT_CHANNEL_2..."
curl -X GET "http://$MICROSERVICE_URL/$DEVICE_FQDN/videoinputstatus/$INPUT_CHANNEL_2"
sleep 1

echo "Testing GET Videoinputstatus for Input $INPUT_CHANNEL_3..."
curl -X GET "http://$MICROSERVICE_URL/$DEVICE_FQDN/videoinputstatus/$INPUT_CHANNEL_3"
sleep 1

echo "=============================================="
echo "Final state check - Getting current values..."
echo "=============================================="

# Final state check
echo "Final GET Volume for Channel $VOLUME_CHANNEL..."
curl -X GET "http://$MICROSERVICE_URL/$DEVICE_FQDN/volume/$VOLUME_CHANNEL"
sleep 1

echo "Final GET Videoroute for Output $OUTPUT_CHANNEL..."
curl -X GET "http://$MICROSERVICE_URL/$DEVICE_FQDN/videoroute/$OUTPUT_CHANNEL"
sleep 1

echo "Final GET Audioandvideoroute for Output $OUTPUT_CHANNEL..."
curl -X GET "http://$MICROSERVICE_URL/$DEVICE_FQDN/audioandvideoroute/$OUTPUT_CHANNEL"
sleep 1

echo "Final GET Audiomute for Channel $OUTPUT_CHANNEL..."
curl -X GET "http://$MICROSERVICE_URL/$DEVICE_FQDN/audiomute/$OUTPUT_CHANNEL"
sleep 1

echo "Final GET Videomute for Channel $OUTPUT_CHANNEL..."
curl -X GET "http://$MICROSERVICE_URL/$DEVICE_FQDN/videomute/$OUTPUT_CHANNEL"
sleep 1

echo "Final GET Occupancystatus for Channel $OUTPUT_CHANNEL..."
curl -X GET "http://$MICROSERVICE_URL/$DEVICE_FQDN/occupancystatus/$OUTPUT_CHANNEL"
sleep 1

echo "=============================================="
echo "All Kramer Switcher Microservice API tests completed!"
echo "=============================================="
