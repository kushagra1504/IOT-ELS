#This program is free software: you can redistribute it and/or modify it under the terms of the GNU General Public License as published by the Free Software Foundation, either version 3 of the License, or (at your option) any later version. This program is distributed in the hope that it will be useful, but WITHOUT ANY WARRANTY; without even the implied warranty of MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE. See the GNU General Public License for more details. You should have received a copy of the GNU General Public License along with this program. If not, see <https://www.gnu.org/licenses/>.

import json
import boto3
import time
import datetime

from io import StringIO
from io import BytesIO

dynamodb = boto3.resource('dynamodb')
table = dynamodb.Table('IoTDeviceState')
sns = boto3.client('sns')
client = boto3.client('iot-data')

#print("environment variable: " + os.environ['snsTopic'])
#snsTopic = os.environ['snsTopic']
snsTopic = "arn:aws:sns:xxxx:xxxx:sendEmail"

def isDoorOpenedOutOfSafeTime(timeA,timeB,time2):
			timeA = datetime.datetime.strptime(timeA, "%H:%M")
			timeB = datetime.datetime.strptime(timeB, "%H:%M")
			time2= datetime.datetime.strptime(time2, "%H:%M")
			if timeA <= time2 and timeB >= time2:
				 return True
			return False

def isAuthenticated(doorID):
			response = client.get_thing_shadow(thingName=doorID)
			t = response['payload']
			shadowData = json.loads(t.read())
			if(shadowData['state']['reported']['authenticationStatus'] =="NO" or (shadowData['state']['reported']['authenticationStatus'] =="YES" and ((int(time.time())- shadowData["metadata"]["reported"]["authenticationStatus"]["timestamp"])/60 >=1))):
				return "Unauthenticated"
			return shadowData['state']['reported']['userName']


def lambda_handler(event, context):
			print("Received event: " + json.dumps(event, indent=2))

			authStatus = isAuthenticated(event['doorID'])

			if event['recordType'] == 'Door' and event['doorStatus'] == "CLOSE":
				client.update_thing_shadow(thingName = event['doorID'], payload=BytesIO(b'{"state":{"reported":{"authenticationStatus": "NO","userName": "N/A"}}}'))
			else:
				event['userName'] = authStatus
######### Write to DynamoDB
			table.put_item(
				 Item= event
			)
			print ("Save the information to DynamoDB")
#############
			errorMessageToDynamoDB = {'doorID': event['doorID'],'statusChangeTime':str(int(event['statusChangeTime'])+1),'recordType':'Alarm'}
			errorStatus = False

		# UTC time of the alarm period
			if event['recordType'] == 'Door' and event['doorStatus'] == "OPEN" and (authStatus == "Unauthenticated"):
						response = sns.publish(
								TopicArn=snsTopic,
								Message='{"Message":"'+event['doorID']+' Door have been opened out of hour/unknown person at '+time.strftime('%d/%m/%Y %H:%M:%S %Z', time.gmtime(int(int(event['statusChangeTime'])/1000))) +'"}'
							)
						errorStatus = True
						errorMessageToDynamoDB['errorType'] = "DoorBreakIn"


			elif event['recordType'] == 'Device' and event['deviceStatus'] == "Disconnected" :
						response = sns.publish(
								TopicArn=snsTopic,
								Message='{"Message":"'+event['doorID']+' Device is disconnected at '+ time.strftime('%d/%m/%Y %H:%M:%S %Z', time.gmtime(int(int(event['statusChangeTime'])/1000))) +'"}'
							)
						errorStatus = True
						errorMessageToDynamoDB['errorType'] = "Device"


			elif(event["recordType"] == "Alarm" and event["errorType"]=="DoorKeepOpen"):
						print ("Alarm should be sent ")
						response = sns.publish(
										TopicArn=snsTopic,
										Message='{"Message":"'+ event['doorID']+' Door have been opened for more than 1 minutes '+time.strftime('%d/%m/%Y %H:%M:%S %Z', time.gmtime(int(int(event['statusChangeTime'])/1000)))  +'"}'
									)


			if errorStatus:
					table.put_item(
						 Item= errorMessageToDynamoDB
					)
