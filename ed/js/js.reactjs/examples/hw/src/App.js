import React, { Component } from 'react';
import logo from './logo.svg';

class App extends Component {
  render() {
    return (
      <div>
        <img src={logo} className="App-logo" alt="logo" style={{width: '50px', height: '50px'}} />
        <h5>200 OK</h5>
      </div>
    );
  }
}

export default App;
