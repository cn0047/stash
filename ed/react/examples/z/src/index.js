import React from 'react';
import ReactDOM from 'react-dom';

import '../public/main.css';
import App from './App';

let rb = JSON.stringify({
  api_key: "",
  app_type_id: "",
  email: "cnfxlr+1@gmail.com",
  password: "12345"
});
fetch('https://api.ziipr.com/v1/auth', {method: 'post', body: rb}).then((res) => {
  return res.json().then(d => {
    let user = {token: d.token, userId: d.user.user_id};
    ReactDOM.render(<App user={user} />, document.getElementById('root'));
  });
});
