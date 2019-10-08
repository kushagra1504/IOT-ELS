import Paper from '@material-ui/core/Paper';
import { withStyles } from '@material-ui/core/styles';
import Table from "@material-ui/core/Table";
import TableBody from "@material-ui/core/TableBody";
import TableCell from "@material-ui/core/TableCell";
import TableHead from "@material-ui/core/TableHead";
import TableRow from "@material-ui/core/TableRow";
import moment from 'moment';
import React, { Component } from "react";

const styles = theme => ({
    root: {
        width: "100%",
        marginTop: 0,
        overflowX: "auto",
        background: "#F8F6F0",
        color: "#000000"
      },
      table: {
      },
  });
  
class DeviceList extends Component {
  render() {
    const { doors, classes } = this.props;
    console.log(doors);
    return (
      <Paper className={classes.root}>
        <Table className={classes.table}>
          <TableHead>
            <TableRow>
              <TableCell>Door ID</TableCell>
              <TableCell>Door Status</TableCell>
              <TableCell>Timestamp</TableCell>
              <TableCell>Username</TableCell>
              <TableCell>Record Type</TableCell>
              <TableCell>Error Type</TableCell>
              <TableCell>Device Status</TableCell>
            </TableRow>
          </TableHead>
          <TableBody>
            {doors.map(door => (
              <TableRow key={door.doorId}>
                <TableCell>{door.doorID}</TableCell>
                <TableCell>{door.doorStatus}</TableCell>
                <TableCell>{
                   moment(parseInt(door.statusChangeTime)).format("DD MMM YYYY hh:mm a")
                    }</TableCell>
                    <TableCell>{door.username}</TableCell>
                    <TableCell>{door.recordType}</TableCell>
                    <TableCell>{door.errorType}</TableCell>
                    <TableCell>{door.deviceStatus}</TableCell>
              </TableRow>
            ))}
          </TableBody>
        </Table>
      </Paper>
    );
  }
}

export default withStyles(styles)(DeviceList);
