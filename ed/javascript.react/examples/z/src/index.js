import React from 'react';
import { browserHistory } from 'react-router';
import { render } from 'react-dom';
import { syncHistoryWithStore } from 'react-router-redux';

import './../public/main.css';
import App from './containers/app';
import configureStore from './store/main';

const store = configureStore();
const history = syncHistoryWithStore(browserHistory, store);
render(
  <App store={store} history={history} />,
  document.getElementById('root')
);
