import { AWSIoTProvider } from "@aws-amplify/pubsub/lib/Providers";
import Button from "@material-ui/core/Button";
import Card from "@material-ui/core/Card";
import { withStyles } from "@material-ui/core/styles";
import Typography from "@material-ui/core/Typography";

import React from "react";
import "./App.css";

import Amplify, { Auth, PubSub } from "aws-amplify";
import { withAuthenticator } from "aws-amplify-react";
import awsconfig from "./aws-exports";

const styles = theme => ({
  ledContainer: {
    display: "flex",
    alignSelf: "center",
    justifyContent: "center"
  },
  ledBox: {
    margin: "10px"
  },
  container: {
    display: "flex",
    flexWrap: "wrap",
    overflow: "scroll",
    position: "relative",
    width: "100%",
    height: "100vh",
    alignSelf: "center",
    justifyContent: "center",
    padding: 5,
    margin: 5,
    backgroundColor: "#FF9900"
  },
  card: {
    display: "flow-root",
    margin: "20px",
    minWidth: "300px",
    height: "200px",
    textAlign: "center"
  },
  details: {
    display: "flex",
    flexDirection: "column"
  },
  content: {
    flex: "1 0 auto"
  },
  cover: {
    width: 151
  },
  controls: {
    display: "flex",
    alignItems: "center",
    paddingLeft: theme.spacing(1),
    paddingBottom: theme.spacing(1)
  },
  playIcon: {
    height: 38,
    width: 38
  },
  ledGreen: {
    margin: "0 auto",
    width: "32px",
    height: "32px",
    backgroundColor: "#ABFF00",
    borderRadius: "50%"
  },
  ledRed: {
    margin: "0 auto",
    width: "32px",
    height: "32px",
    backgroundColor: "#F00",
    borderRadius: "50%"
  },
  button: {
    margin: theme.spacing(1)
  }
});

Auth.configure(awsconfig);

/*=======================================================================
                    AMPLIFY PUBSUB 
=======================================================================*/

Amplify.addPluggable(
  new AWSIoTProvider({
    aws_pubsub_region: "REGION",
    aws_pubsub_endpoint:
      "wss://XXXXXXXXX-ats.iot.REGION.amazonaws.com/mqtt"
  })
);

/*=======================================================================
                    APP COMPONENT
=======================================================================*/
class App extends React.Component {
  constructor(props) {
    super(props);

    this.subscribe();

    this.state = {
      grp1challenge1Status: "off",
      grp1challenge2Status: "off",
      grp1challenge3Status: "off",
      grp2challenge1Status: "off",
      grp2challenge2Status: "off",
      grp2challenge3Status: "off",
      grp3challenge1Status: "off",
      grp3challenge2Status: "off",
      grp3challenge3Status: "off",
      grp4challenge1Status: "off",
      grp4challenge2Status: "off",
      grp4challenge3Status: "off",
      grp5challenge1Status: "off",
      grp5challenge2Status: "off",
      grp5challenge3Status: "off",
      grp6challenge1Status: "off",
      grp6challenge2Status: "off",
      grp6challenge3Status: "off",
      grp7challenge1Status: "off",
      grp7challenge2Status: "off",
      grp7challenge3Status: "off",
      grp8challenge1Status: "off",
      grp8challenge2Status: "off",
      grp8challenge3Status: "off",
      grp9challenge1Status: "off",
      grp9challenge2Status: "off",
      grp9challenge3Status: "off",
      grp10challenge1Status: "off",
      grp10challenge2Status: "off",
      grp10challenge3Status: "off",
    };
  }

   subscribe() {
    PubSub.subscribe("$aws/things/demo/shadow/get/accepted").subscribe({
        next: data => this.handleGetJson(data),
        error: error => console.error(error),
        close: () => console.log("Done")
      });


     PubSub.subscribe("$aws/things/demo/shadow/update/documents").subscribe({
      next: data => this.handleJson(data),
      error: error => console.error(error),
      close: () => console.log("Done")
    });


    setInterval(async ()=>{
        await PubSub.publish("$aws/things/demo/shadow/get", {});
    }, 2000);
  }

  handleGetJson(data) {
    console.debug("have input json:", data);
    this.setState({grp1challenge1Status: data.value.state.reported.group1.challenge1})
    this.setState({grp1challenge2Status: data.value.state.reported.group1.challenge2})
    this.setState({grp1challenge3Status: data.value.state.reported.group1.challenge3})

    this.setState({grp2challenge1Status: data.value.state.reported.group2.challenge1})
    this.setState({grp2challenge2Status: data.value.state.reported.group2.challenge2})
    this.setState({grp2challenge3Status: data.value.state.reported.group2.challenge3})

    this.setState({grp3challenge1Status: data.value.state.reported.group3.challenge1})
    this.setState({grp3challenge2Status: data.value.state.reported.group3.challenge2})
    this.setState({grp3challenge3Status: data.value.state.reported.group3.challenge3})

    this.setState({grp4challenge1Status: data.value.state.reported.group4.challenge1})
    this.setState({grp4challenge2Status: data.value.state.reported.group4.challenge2})
    this.setState({grp4challenge3Status: data.value.state.reported.group4.challenge3})

    this.setState({grp5challenge1Status: data.value.state.reported.group5.challenge1})
    this.setState({grp5challenge2Status: data.value.state.reported.group5.challenge2})
    this.setState({grp5challenge3Status: data.value.state.reported.group5.challenge3})

    this.setState({grp6challenge1Status: data.value.state.reported.group6.challenge1})
    this.setState({grp6challenge2Status: data.value.state.reported.group6.challenge2})
    this.setState({grp6challenge3Status: data.value.state.reported.group6.challenge3})

    this.setState({grp7challenge1Status: data.value.state.reported.group7.challenge1})
    this.setState({grp7challenge2Status: data.value.state.reported.group7.challenge2})
    this.setState({grp7challenge3Status: data.value.state.reported.group7.challenge3})

    this.setState({grp8challenge1Status: data.value.state.reported.group8.challenge1})
    this.setState({grp8challenge2Status: data.value.state.reported.group8.challenge2})
    this.setState({grp8challenge3Status: data.value.state.reported.group8.challenge3})

    this.setState({grp9challenge1Status: data.value.state.reported.group9.challenge1})
    this.setState({grp9challenge2Status: data.value.state.reported.group9.challenge2})
    this.setState({grp9challenge3Status: data.value.state.reported.group9.challenge3})

    this.setState({grp10challenge1Status: data.value.state.reported.group10.challenge1})
    this.setState({grp10challenge2Status: data.value.state.reported.group10.challenge2})
    this.setState({grp10challenge3Status: data.value.state.reported.group10.challenge3})
  };

  handleJson(data) {
    console.debug("have input json:", data);
    this.setState({grp1challenge1Status: data.value.current.state.reported.group1.challenge1})
    this.setState({grp1challenge2Status: data.value.current.state.reported.group1.challenge2})
    this.setState({grp1challenge3Status: data.value.current.state.reported.group1.challenge3})

    this.setState({grp2challenge1Status: data.value.current.state.reported.group2.challenge1})
    this.setState({grp2challenge2Status: data.value.current.state.reported.group2.challenge2})
    this.setState({grp2challenge3Status: data.value.current.state.reported.group2.challenge3})

    this.setState({grp3challenge1Status: data.value.current.state.reported.group3.challenge1})
    this.setState({grp3challenge2Status: data.value.current.state.reported.group3.challenge2})
    this.setState({grp3challenge3Status: data.value.current.state.reported.group3.challenge3})

    this.setState({grp4challenge1Status: data.value.current.state.reported.group4.challenge1})
    this.setState({grp4challenge2Status: data.value.current.state.reported.group4.challenge2})
    this.setState({grp4challenge3Status: data.value.current.state.reported.group4.challenge3})

    this.setState({grp5challenge1Status: data.value.current.state.reported.group5.challenge1})
    this.setState({grp5challenge2Status: data.value.current.state.reported.group5.challenge2})
    this.setState({grp5challenge3Status: data.value.current.state.reported.group5.challenge3})

    this.setState({grp6challenge1Status: data.value.current.state.reported.group6.challenge1})
    this.setState({grp6challenge2Status: data.value.current.state.reported.group6.challenge2})
    this.setState({grp6challenge3Status: data.value.current.state.reported.group6.challenge3})

    this.setState({grp7challenge1Status: data.value.current.state.reported.group7.challenge1})
    this.setState({grp7challenge2Status: data.value.current.state.reported.group7.challenge2})
    this.setState({grp7challenge3Status: data.value.current.state.reported.group7.challenge3})

    this.setState({grp8challenge1Status: data.value.current.state.reported.group8.challenge1})
    this.setState({grp8challenge2Status: data.value.current.state.reported.group8.challenge2})
    this.setState({grp8challenge3Status: data.value.current.state.reported.group8.challenge3})

    this.setState({grp9challenge1Status: data.value.current.state.reported.group9.challenge1})
    this.setState({grp9challenge2Status: data.value.current.state.reported.group9.challenge2})
    this.setState({grp9challenge3Status: data.value.current.state.reported.group9.challenge3})

    this.setState({grp10challenge1Status: data.value.current.state.reported.group10.challenge1})
    this.setState({grp10challenge2Status: data.value.current.state.reported.group10.challenge2})
    this.setState({grp10challenge3Status: data.value.current.state.reported.group10.challenge3})
  };

  render() {
    const { classes } = this.props;
    console.log(classes);

    return (
      <div className={classes.container}>
        <Card className={classes.card}>
          <Typography variant="h4" gutterBottom>
            Group 1
          </Typography>
          <div className={classes.ledContainer}>
            <div className={classes.ledBox}>
              <div
                className={
                  this.state.grp1challenge1Status === "on"
                    ? classes.ledGreen
                    : classes.ledRed
                }
              ></div>
              <p>Challenge 1</p>
            </div>
            <div className={classes.ledBox}>
              <div
                className={
                  this.state.grp1challenge2Status === "on"
                    ? classes.ledGreen
                    : classes.ledRed
                }
              ></div>
              <p>Challenge 2</p>
            </div>
            <div className={classes.ledBox}>
              <div
                className={
                  this.state.grp1challenge3Status === "on"
                    ? classes.ledGreen
                    : classes.ledRed
                }
              ></div>
              <p>Challenge 3</p>
            </div>
          </div>
        </Card>
        <Card className={classes.card}>
          <Typography variant="h4" gutterBottom>
            Group 2
          </Typography>
          <div className={classes.ledContainer}>
            <div className={classes.ledBox}>
              <div
                className={
                  this.state.grp2challenge1Status === "on"
                    ? classes.ledGreen
                    : classes.ledRed
                }
              ></div>
              <p>Challenge 1</p>
            </div>
            <div className={classes.ledBox}>
              <div
                className={
                  this.state.grp2challenge2Status === "on"
                    ? classes.ledGreen
                    : classes.ledRed
                }
              ></div>
              <p>Challenge 2</p>
            </div>
            <div className={classes.ledBox}>
              <div
                className={
                  this.state.grp2challenge3Status == "on"
                    ? classes.ledGreen
                    : classes.ledRed
                }
              ></div>
              <p>Challenge 3</p>
            </div>
          </div>
        </Card>
        <Card className={classes.card}>
          <Typography variant="h4" gutterBottom>
            Group 3
          </Typography>
          <div className={classes.ledContainer}>
            <div className={classes.ledBox}>
              <div
                className={
                  this.state.grp3challenge1Status == "on"
                    ? classes.ledGreen
                    : classes.ledRed
                }
              ></div>
              <p>Challenge 1</p>
            </div>
            <div className={classes.ledBox}>
              <div
                className={
                  this.state.grp3challenge2Status == "on"
                    ? classes.ledGreen
                    : classes.ledRed
                }
              ></div>
              <p>Challenge 2</p>
            </div>
            <div className={classes.ledBox}>
              <div
                className={
                  this.state.grp3challenge3Status == "on"
                    ? classes.ledGreen
                    : classes.ledRed
                }
              ></div>
              <p>Challenge 3</p>
            </div>
          </div>
        </Card>
        <Card className={classes.card}>
          <Typography variant="h4" gutterBottom>
            Group 4
          </Typography>
          <div className={classes.ledContainer}>
            <div className={classes.ledBox}>
              <div
                className={
                  this.state.grp4challenge1Status == "on"
                    ? classes.ledGreen
                    : classes.ledRed
                }
              ></div>
              <p>Challenge 1</p>
            </div>
            <div className={classes.ledBox}>
              <div
                className={
                  this.state.grp4challenge2Status == "on"
                    ? classes.ledGreen
                    : classes.ledRed
                }
              ></div>
              <p>Challenge 2</p>
            </div>
            <div className={classes.ledBox}>
              <div
                className={
                  this.state.grp4challenge3Status == "on"
                    ? classes.ledGreen
                    : classes.ledRed
                }
              >
              </div>
              <p>Challenge 3</p>
            </div>
          </div>
        </Card>
        <Card className={classes.card}>
          <Typography variant="h4" gutterBottom>
            Group 5
          </Typography>
          <div className={classes.ledContainer}>
            <div className={classes.ledBox}>
              <div
                className={
                  this.state.grp5challenge1Status == "on"
                    ? classes.ledGreen
                    : classes.ledRed
                }
              >
              </div>
              <p>Challenge 1</p>
            </div>
            <div className={classes.ledBox}>
              <div
                className={
                  this.state.grp5challenge2Status == "on"
                    ? classes.ledGreen
                    : classes.ledRed
                }
              ></div>
              <p>Challenge 2</p>
            </div>
            <div className={classes.ledBox}>
              <div
                className={
                  this.state.grp5challenge3Status == "on"
                    ? classes.ledGreen
                    : classes.ledRed
                }
              ></div>
              <p>Challenge 3</p>
            </div>
          </div>
        </Card>
      </div>
    );
  }
}

export default withStyles(styles)(withAuthenticator(App, true));
