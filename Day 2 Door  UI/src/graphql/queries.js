/* eslint-disable */
// this is an auto generated file. This will be overwritten
import gql from "graphql-tag";

export const listIoTDeviceStates = gql(`
query (
    $filter: TableIoTDeviceStateFilterInput
    $limit: Int
    $nextToken: String
  ) {
    listIoTDeviceStates(filter: $filter, limit: $limit, nextToken: $nextToken) {
      items {
        doorID
        statusChangeTime
        username
        recordType
        doorStatus
        deviceStatus
        errorType
        authenticationStatus
      }
      nextToken
    }
  }`);