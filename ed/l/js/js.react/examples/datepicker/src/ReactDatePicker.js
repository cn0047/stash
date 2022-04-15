import React, { useState } from "react";
import DatePicker from "react-datepicker";

import './App.css';
import "react-datepicker/dist/react-datepicker.css";
// import 'react-datepicker/dist/react-datepicker-cssmodules.css';

function ReactDatePicker() {
  const [startDate, setStartDate] = useState(new Date());
  const ExampleCustomTimeInput = ({ date, value, onChange }) => (
    <input
      value={value}
      onChange={(e) => onChange(e.target.value)}
      style={{ border: "solid 1px pink" }}
    />
  );
  return (
    <div className="">
      <p>ReactDatePicker</p>
      <div className="App-content">
        <br /> dp1:
        <DatePicker selected={startDate} onChange={(date:Date) => setStartDate(date)} />

        <br /> dp2:
        <DatePicker
          selected={startDate} onChange={(date) => setStartDate(date)}
          showTimeInput customTimeInput={<ExampleCustomTimeInput/>}
        />

        <br /> dp3:
        <DatePicker
          selected={startDate} onChange={(date) => setStartDate(date)}
          showTimeSelect timeInputLabel="Time:" dateFormat="MM/dd/yyyy h:mm aa"
        />

        <br /> dp4:
        <DatePicker
          selected={startDate} onChange={(date) => setStartDate(date)}
          showTimeInput timeInputLabel="Time:" dateFormat="MM/dd/yyyy h:mm aa"
        />
      </div>
    </div>
  );
}

export default ReactDatePicker;
