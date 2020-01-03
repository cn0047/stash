import React from 'react';
import ReactDOM from 'react-dom';
import { combineReducers, createStore } from 'redux'

export default function() {
  console.log('Example 2:');
  ReactDOM.render(<div>Look into console.</div>, document.getElementById('root'));

  function counterReducer(state = 0, action) {
    switch (action.type) {
      case 'INCREMENT':
        return state + 1;
      default:
        return state
    }
  }
  function blackholeReducer(state = [], action) {
    switch (action.type) {
      case 'ADD':
        return [...state, action.text];
      default:
        return state;
    }
  }

  let reducer = combineReducers({counterReducer, blackholeReducer});
  let store = createStore(reducer);
  store.subscribe(() => {
    console.log('counter = %s', store.getState().counterReducer);
    console.log('blackhole:', store.getState().blackholeReducer);
  });
  store.dispatch({type: 'INCREMENT'});
  store.dispatch({type: 'ADD', text: 'New text.'});
}
