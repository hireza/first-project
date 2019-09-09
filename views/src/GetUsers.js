import React from "react";

import axios from "axios";

export default class UserList extends React.Component {
  state = {
    data: {
      users: [],
      count: ""
    }
  };

  componentDidMount() {
    axios.get("http://localhost:1234/").then(res => {
      const obj = res.data;
      const data = JSON.parse(obj);
      this.setState({ data });
    });
  }

  render() {
    return (
      <div>
        <div>
          <h1 id="title">Table users</h1>
          <table id="users">
            <tr>
              <th>ID</th>
              <th>Name</th>
            </tr>
            {this.state.data.users.map(user => (
              <tr>
                <td key={user.id}>{user.id} </td>
                <td key={user.name}>{user.name} </td>
              </tr>
            ))}
          </table>
        </div>
        <div>visit count : {this.state.data.count}</div>
        {/* {console.log(this.state.data.name)} */}
      </div>
    );
  }
}
