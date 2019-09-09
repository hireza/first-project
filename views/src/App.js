import React, { Component } from "react";
import "./App.css";

import GetUsers from "./GetUsers";
import CreateUser from "./CreateUser";

class App extends Component {
  render() {
    return (
      <div width="100px" className="App">
        <GetUsers />

        <br />

        <CreateUser />
      </div>
    );
  }
}

export default App;
