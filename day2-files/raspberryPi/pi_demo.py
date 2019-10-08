#This program is free software: you can redistribute it and/or modify it under the terms of the GNU General Public License as published by the Free Software Foundation, either version 3 of the License, or (at your option) any later version. This program is distributed in the hope that it will be useful, but WITHOUT ANY WARRANTY; without even the implied warranty of MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE. See the GNU General Public License for more details. You should have received a copy of the GNU General Public License along with this program. If not, see <https://www.gnu.org/licenses/>.

from AWSIoTPythonSDK.MQTTLib import AWSIoTMQTTClient
import json
import time
import threading
import RPi.GPIO as GPIO #import the GPIO library
from datetime import datetime

###########
# Config
###########

thingName = "door1"
awsEndpoint = "xxxx-ats.iot.xxxxx.amazonaws.com"
awsPortNumber = 8883
awsRootCAPath = "./AmazonRootCA1.pem"
awsIoTPrivateKeyPath = "./xxxxx-private.pem.key"
awsIoTCertificatePath = "./xxxx-certificate.pem.crt"
awsTopicPrefix = "SYD37/"


###################
# Sensor Pin config
###################
pinNumber=5

GPIO.setmode(GPIO.BCM)
GPIO.setup(pinNumber, GPIO.IN, pull_up_down=GPIO.PUD_UP)
myAWSIoTMQTTClient = None


#############################
# Left Open Alarm
############################
stopThread = False
def thread_function(epochTime):
    while True:
       time.sleep(1)
       if stopThread :
            break
       if (int(time.time() - epochTime)/30) >=1 :
           JSONPayload = '{"recordType":"Alarm","errorType":"DoorKeepOpen"}'
           myAWSIoTMQTTClient.publish(awsTopicPrefix + thingName,JSONPayload,0)
           break


############################
# Updating Door status
###########################
def updateDoorStatus(doorStatus):
    global stopThread
    JSONPayload = '{"recordType": "Door","doorStatus":"'+doorStatus +'"}'
    myAWSIoTMQTTClient.publish(awsTopicPrefix + thingName,JSONPayload,1)

    # DoorKeptOpen Alarm
    if (doorStatus == "CLOSED"):
        stopThread = True
    elif (doorStatus == "OPEN") :
        stopThread = False
        x = threading.Thread(target=thread_function, args=(int(time.time()),))
        x.start()

###############################################
# Update Device Status Connected / Disconnecred
###############################################
def updateDeviceStatus(deviceStatus):
      JSONPayload = '{"recordType": "Device","deviceStatus":"'+deviceStatus +'"}'
      myAWSIoTMQTTClient.publish(awsTopicPrefix + thingName,JSONPayload,1)



########################################



#======================================
# AWS IoT SDK - MQTT Client connection
#======================================
myAWSIoTMQTTClient = AWSIoTMQTTClient(thingName)
myAWSIoTMQTTClient.configureEndpoint(awsEndpoint, awsPortNumber)
myAWSIoTMQTTClient.configureCredentials(awsRootCAPath, awsIoTPrivateKeyPath,awsIoTCertificatePath)
myAWSIoTMQTTClient.configureMQTTOperationTimeout(5)

# Last will payload for disconnection
lastWillPayload = '{"recordType":"Device","deviceStatus":"Disconnected"}'
myAWSIoTMQTTClient.configureLastWill(awsTopicPrefix + thingName, lastWillPayload, 0)


# Connect

myAWSIoTMQTTClient.connect()
updateDeviceStatus("Connected")
dateTimeObj = datetime.now()
print("Device is connected at "+ str(datetime.now()))

previous_status = 2

# Listen to door sensor and publish

while True:
    if (previous_status != GPIO.input(pinNumber)):
        if(GPIO.input(pinNumber) == 1):
            print("Door is opened at "+ str(datetime.now()))
            updateDoorStatus("OPEN")
        else:
            print("Door is closed at "+ str(datetime.now()))
            updateDoorStatus("CLOSED")
        previous_status = GPIO.input(pinNumber)
    time.sleep(1)

