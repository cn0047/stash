import React, { Component } from 'react';
import { render } from 'react-dom';
import PropTypes from 'prop-types';
import { Provider } from 'react-redux';
import { connect } from 'react-redux';
import { createStore, applyMiddleware, compose } from 'redux';
import thunk from 'redux-thunk';
import { combineReducers } from 'redux';



const entities = (state = { status: 'init' }, action) => {
  switch (action.type) {
    case 'ONE':
      return action.payload;
    default:
      return state
  }
};
const rootReducer = combineReducers({ entities });
const configureStore = defaultState => {
  return createStore(
    rootReducer,
    defaultState,
    compose(applyMiddleware(thunk))
  );
};



export const loadDataFromServer = () => (dispatch, getState) => {
  dispatch({type: 'ONE', payload: {status: 'ok', message: 'It works!'}});
};



const loadData = ({ loadDataFromServer }) => {
  loadDataFromServer()
};
class AppComponent extends Component {
  componentWillMount() {
    loadData(this.props)
  }
  render() {
    const data = this.props.data;
    const message = data.status === 'ok' ? data.message : 'Loading...';
    console.log('➡️', message);
    return ( <div>{message}</div> )
  }
}
const mapStateToProps = (state, ownProps) => {
  return {data: state.entities}
};
const App = connect(mapStateToProps, { loadDataFromServer })(AppComponent);



const Root = ({ store }) => (
  <Provider store={store}>
    <App />
  </Provider>
);
Root.propTypes = {
  store: PropTypes.object.isRequired,
};



export default function() {
  const store = configureStore();
  render(
    <Root store={store} />,
    document.getElementById('root')
  );
}
