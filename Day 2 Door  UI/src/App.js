import React from 'react';
import './App.css';
import DoorWithData from './components/door/Door'
import { BrowserRouter as Router, Route } from 'react-router-dom';
import appSyncConfig from "./aws-exports";
import { ApolloProvider } from "react-apollo";
import AWSAppSyncClient, { defaultDataIdFromObject } from "aws-appsync";
import { Rehydrated } from "aws-appsync-react";

 const Home = () => (
    <div className="App">
      <DoorWithData />
    </div>
  );

  
  const App = () => (
    <Router>
      <div>
        <Route exact={true} path="/" component={Home} />
      </div>
    </Router>
  );

  

const client = new AWSAppSyncClient({
    url: appSyncConfig.aws_appsync_graphqlEndpoint,
    region: appSyncConfig.aws_appsync_region,
    auth: {
      type: appSyncConfig.aws_appsync_authenticationType,
      apiKey: appSyncConfig.aws_appsync_apiKey,
    },
    cacheOptions: {
      dataIdFromObject: (obj) => {
        let id = defaultDataIdFromObject(obj);
  
        if (!id) {
          const { __typename: typename } = obj;
          switch (typename) {
            case 'Comment':
              return `${typename}:${obj.commentId}`;
            default:
              return id;
          }
        }
  
        return id;
      }
    }
  });
  
  const WithProvider = () => (
    <ApolloProvider client={client}>
      <Rehydrated>
        <App />
      </Rehydrated>
    </ApolloProvider>
  );
  
  export default WithProvider;