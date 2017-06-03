import React, { Component } from 'react';
import { browserHistory } from 'react-router';
import { connect } from 'react-redux';

import Welcome from './../components/welcome';
import { logIn } from './../actions/me'

const actionLogIn = ({ logIn }) => {
  logIn()
};

class Me extends Component {

  componentWillMount() {
    actionLogIn(this.props);
  }

  componentWillUpdate() {
    browserHistory.push('/summary');
  }

  render() {
    return (
      <Welcome />
    );
  }

}

const mapStateToProps = (state) => {
  return {me: state.me};
};

export default connect(mapStateToProps, { logIn })(Me);
