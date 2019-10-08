import {listIoTDeviceStates as QueryListIoTDeviceStates}  from "../../graphql/queries";
import {onMockIoTDeviceState as MockSubs} from "../../graphql/subscriptions";
import DeviceList from "./DeviceList.jsx";
import React, { PureComponent } from "react";
import { compose, graphql } from "react-apollo";

var doorClassName;

class Door extends PureComponent {
  subscription;

  componentDidMount = () => {
    console.log("Door mounted");
    this.subscription = this.props.subscribeToIoTDeviceState();
  };

  componentWillUnmount() {
    console.log("Door Unmounted");
    this.subscription();
  }

  render() {
    console.log("Inside Render");
    const { listIotDevices, doorState, recordType, authenticationStatus } = this.props;
    

    console.log('listIotDevices', listIotDevices)
    console.log('props', this.props)
    if (doorState === "closed") {
        doorClassName = "doorbox";
    } else if (doorState === "open") {
        doorClassName = "doorbox open";
    }

    var alarm = recordType === 'Alarm' ? 'Alarm' : ""
    var authenticated;

    console.log("recordType = " + recordType );

    console.log("authenticationStatus = " + authenticationStatus );

    if (recordType === 'Authentication' && authenticationStatus === 'YES') {
        authenticated = 'AUTHENTICATED';
    } else if (recordType === 'Authentication' && authenticationStatus === 'NO') {
        authenticated = 'UNAUTHENTICATED';
    } else {
        authenticated = '';
    }
    
    return (
      <div className="DoorWrapper">
        <div className="container">
            <div>
                <div
                className={doorClassName}
                index="0">
                    <div className="door">
                        <img className="aws-logo" src="https://d0.awsstatic.com/logos/powered-by-aws.png" alt="Powered by AWS Cloud Computing"></img>
                    </div>
                </div>
                <span className={recordType === alarm ? 'blinking' : 'not-blinking'}> {alarm}</span>
                <span className={recordType === 'Authentication' ? 'blinking' : 'not-blinking'}> {authenticated}</span>
                
            </div>
          
          <DeviceList doors={listIotDevices} />
        </div>
        
      </div>
    );
  }
}

const DoorWithData = compose(
  graphql(QueryListIoTDeviceStates, {
    options: () => ({
      fetchPolicy: "network-only"
    }),
    props: props => ({
      listIotDevices: props.data.listIoTDeviceStates
        ? props.data.listIoTDeviceStates.items
        : [],
      doorState:
        (props.data.listIoTDeviceStates &&
        props.data.listIoTDeviceStates.items.length > 0)
          ? ((props.data.listIoTDeviceStates.items[
            props.data.listIoTDeviceStates.items.length - 1
          ].doorStatus) ? props.data.listIoTDeviceStates.items[
            props.data.listIoTDeviceStates.items.length - 1
          ].doorStatus.toLowerCase() : "")
          : "closed",
        recordType:
        props.data.listIoTDeviceStates &&
        props.data.listIoTDeviceStates.items.length > 0
          ? props.data.listIoTDeviceStates.items[
              props.data.listIoTDeviceStates.items.length - 1
            ].recordType
          : "",
          authenticationStatus:props.data.listIoTDeviceStates &&
          props.data.listIoTDeviceStates.items.length > 0
            ? props.data.listIoTDeviceStates.items[
                props.data.listIoTDeviceStates.items.length - 1
              ].authenticationStatus
            : "",
      subscribeToIoTDeviceState: () =>
        props.data.subscribeToMore({
          document: MockSubs,
          variables: null,
          updateQuery: (prev, { subscriptionData }) => {
            console.log(prev)
            console.log('subscriptionData', subscriptionData.data)
            const { onMockIoTDeviceState } = subscriptionData.data;
            const res = {
              ...prev,
              doorState: onMockIoTDeviceState.doorStatus.toLowerCase(),
              recordType: onMockIoTDeviceState.recordType,
              authenticationStatus: onMockIoTDeviceState.authenticationStatus,
              listIoTDeviceStates: {
                ...prev.listIoTDeviceStates,
                items: [...prev.listIoTDeviceStates.items, onMockIoTDeviceState]
              }
            };
            return res;
          }
        })
    })
  })
)(Door);

export default DoorWithData;
