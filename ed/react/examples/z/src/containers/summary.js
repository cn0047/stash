import React, { Component } from 'react';
import { browserHistory } from 'react-router';
import { connect } from 'react-redux';

import Main from './../components/summary/main';
import { loadFeed } from './../actions/summary';

const actionLoadFeed = ({ me, loadFeed }) => {
  loadFeed(me);
};

class SummaryScreen extends Component {

  componentWillMount() {
    if (Object.keys(this.props.me).length === 0) {
      return browserHistory.push('/');
    }
    actionLoadFeed(this.props);
  }

  render() {
    return (
      <Main {...this.props} />
    );
  }

}

const mapStateToProps = (state) => {
  return {me: state.me, summary: state.summary};
};

export default connect(mapStateToProps, { loadFeed })(SummaryScreen);
