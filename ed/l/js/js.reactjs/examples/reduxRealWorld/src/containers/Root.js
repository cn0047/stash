import React from 'react'
import PropTypes from 'prop-types'
import { Provider } from 'react-redux'
import { Router, Route } from 'react-router'
import UserPage from './UserPage'

const Root = ({ store, history }) => (
  <Provider store={store}>
    <div>
      <Router history={history} >
      <Route path="/" component={UserPage}>
        <Route path="/:login" component={UserPage} />
      </Route>
      </Router>
    </div>
  </Provider>
);
Root.propTypes = {
  store: PropTypes.object.isRequired,
  history: PropTypes.object.isRequired
};
export default Root
