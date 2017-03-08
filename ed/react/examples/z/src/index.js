import React from 'react';
import ReactDOM from 'react-dom';

import '../public/main.css';
import App from './App';
import { MainConfig } from './Config/Main';

let rb = JSON.stringify(MainConfig.auth);
fetch(MainConfig.apiUrl + '/auth', {method: 'post', body: rb}).then((res) => {
  return res.json().then(d => {
    let user = {token: d.token, userId: d.user.user_id};
    ReactDOM.render(<App user={user} />, document.getElementById('root'));
  });
});
