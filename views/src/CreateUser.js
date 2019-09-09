import React from "react";

import axios from "axios";
import qs from "qs";

export default class CreateUser extends React.Component {
  state = {
    name: ""
  };

  handleChange = event => {
    this.setState({ name: event.target.value });
  };

  handleSubmit = event => {
    event.preventDefault();

    const name = this.state.name;

    axios.post(`http://localhost:1234/`, { name }).then(res => {
      console.log(res);
      console.log(res.data);
    });
  };

  render() {
    return (
      <div>
        <form onSubmit={this.handleSubmit}>
          <label>User Name: </label>
          <input type="text" name="name" onChange={this.handleChange} />
          <button type="submit">Add</button>
        </form>
      </div>
    );
  }
}
