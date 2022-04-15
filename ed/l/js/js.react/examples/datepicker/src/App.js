import React from "react";

import './App.css';

import ReactDatePicker from './ReactDatePicker';
import MUIDatePicker from './MUIDatePicker';

function App() {
  return (
    <div className="App">
      <header className="App-header">
        <p>DatePickers</p>
      </header>
      <div>
        {/*<ReactDatePicker />*/}
        <MUIDatePicker />
      </div>
    </div>
  );
}

export default App;
