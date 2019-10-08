from AWSIoTPythonSDK.MQTTLib import AWSIoTMQTTClient
import logging
import time
import argparse
import json, sys, subprocess, threading
import uuid

def light_up():
    subprocess.call(["challenge2Monitor"])

msg = ('Remember to save your JSON file after changes.'
        'Good Luck !!!'
)

# init actual MQTT client
def client_pub():
    print('------Initializing your MQTT client-----')
    myAWSIoTMQTTClient = AWSIoTMQTTClient(clientId)
    myAWSIoTMQTTClient.configureEndpoint(endpoint, port)
    myAWSIoTMQTTClient.configureCredentials(rootCAPath, privateKeyPath, certificatePath)
    myAWSIoTMQTTClient.connect()

    # init thread to verify challenge status 
    # ------ DO NOT Changed -----
    thread1 = threading.Thread(target=light_up)
    thread1.start()
    time.sleep(2)
    # finish init thread section

    #   ------   TODO    ----------
    #  update the topic to shadow update in file mqttclient_config.json
    #  update message should be "<group-number>":{"challenge2":"on"}
    #  Generate reported statement for shadow message. 
    #  Message update below
    message = {}
    # Message update above using json. 
    print('------Please check your group challenge 2 light status-----')	
    myAWSIoTMQTTClient.publish(topic,json.dumps(message), 1)
	
# Init parser
parser = argparse.ArgumentParser(
    description = 'Please change the JSON file accordingly so the Iot device will connect and publish to the topic.',
    epilog = """Remember to save your JSON file after changes. \nAnd Good Luck !!!"""
)
args=parser.parse_args()

# Configure logging
logger_iot = logging.getLogger("AWSIoTPythonSDK.core")
logger_iot.setLevel(logging.ERROR)
logger_json = logging.getLogger("JSON_File_Error")
logger_json.setLevel(logging.DEBUG)
streamHandler = logging.StreamHandler()
formatter = logging.Formatter('%(asctime)s - %(name)s - %(levelname)s - %(message)s')
streamHandler.setFormatter(formatter)
logger_iot.addHandler(streamHandler)
logger_json.addHandler(streamHandler)

# Load MQTT config JSON
try:
    print('\n')
    print('-------Loading the JSON file------')
    with open ('/home/awstechevent/scripts/mqttclient_config.json') as f:
        mqtt_config = json.load(f)
        print('\n')
        print('Succesfully Loaded JSON File')
        print('\n')


except FileNotFoundError as e:
    logger_json.debug('FileNotFoundError: ' + str(e))
    sys.exit('File Not Found')

except json.JSONDecodeError as e:
    print(type(e).__name__)
    logger_json.debug('FileNotFoundError: ' + str(e))
    sys.exit('File Not Found')

except NameError as e:
    logger_json.debug('NameError: ' + str(e))
    sys.exit('File Not Found')

except KeyError as e:
    logger_json.debug('KeyError: ' + str(e))
    sys.exit('File Not Found')

# Parse the client config, cert, nedpoint, etc....
print('-----Parsing Your JSON File------')
endpoint = mqtt_config['endpoint']
rootCAPath = mqtt_config['rootCAPath']
certificatePath = mqtt_config['certificatePath']
privateKeyPath = mqtt_config['privateKeyPath']
port = mqtt_config['port']
topic = mqtt_config['topic']
clientId = mqtt_config['clientId']

# init start program
if __name__ == '__main__':
    client_pub()
