import React, { Component } from 'react';
import { browserHistory, Router, IndexRoute, Route } from 'react-router'

import SummaryScreen from './SummaryScreen/SummaryScreen';

class App extends Component {

  constructor(props) {
    super(props);
    this.state = {user: this.props.user};
  }

  render() {
    return (
      <Router history={browserHistory}>
        <Route path="/">
          <IndexRoute component={SummaryScreen} user={this.state.user} />
        </Route>
      </Router>
    );
  }

}

export default App;
