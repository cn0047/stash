import PropTypes from 'prop-types';
import React from 'react';
import { Provider } from 'react-redux';
import { Router, Route, IndexRoute } from 'react-router';

import Me from './me';
import SummaryScreen from './summary';

const App = ({ store, history }) => (
  <Provider store={store}>
    <Router history={history} >
      <Route path="/">
        <IndexRoute component={Me}/>
        <Route path="/summary" component={SummaryScreen}/>
      </Route>
    </Router>
  </Provider>
);

App.propTypes = {
  store: PropTypes.object.isRequired,
  history: PropTypes.object.isRequired
};

export default App
