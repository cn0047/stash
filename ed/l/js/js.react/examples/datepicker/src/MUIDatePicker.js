import React, { useState } from "react";

import "date-fns";
import DateFnsUtils from "@date-io/date-fns";
import {DateTimePicker, KeyboardDatePicker, MuiPickersUtilsProvider,} from "@material-ui/pickers";
import TextField from "@material-ui/core/TextField";

function MUIDatePicker() {
  const [startDate, setStartDate] = useState(new Date("2021-03-01T21:11:54"));

  const [open, setOpen] = React.useState(false);

  const onChange1 = (v) => {
    console.log('[onChange1]', v);
    setStartDate(v);
  }

  return (
    <div className="">
      <p>MUIDatePicker</p>
      <div className="App-content">
        <MuiPickersUtilsProvider utils={DateFnsUtils}>
          <br /> dp1:
          <DateTimePicker
            renderInput={props => <TextField {...props} />}
            value={startDate}
            onChange={setStartDate}
            inputFormat="dd/MM/yyyy HH:mm"
          />

          <br /> dp2:
          <KeyboardDatePicker
            open={open}
            disableToolbar
            variant="inline"
            format="MM/dd/yyyy"
            margin="normal"
            id="date-picker-inline"
            label="Date picker inline"
            value={startDate}
            onChange={setStartDate}
            onClose={() => setOpen(false)}
            onOpen={() => setOpen(true)}
          />

          <br /> dp3:
          <TextField
            id="datetime-local"
            label="Next appointment"
            type="datetime-local"
            defaultValue="2017-05-24T10:30"
            sx={{ width: 250 }}
            InputLabelProps={{shrink: true}}
            onChange={onChange1}
          />
        </MuiPickersUtilsProvider>
      </div>
    </div>
  );
}

export default MUIDatePicker;
