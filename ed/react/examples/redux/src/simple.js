import React from 'react';
import ReactDOM from 'react-dom';
import { createStore } from 'redux'

export default function() {
  console.log('Example 1:');
  ReactDOM.render(<div>Look into console.</div>, document.getElementById('root'));

  function counterReducer(state = 0, action) {
    switch (action.type) {
      case 'INCREMENT':
        return state + 1;
      case 'DECREMENT':
        return state - 1;
      default:
        return state
    }
  }
  let store = createStore(counterReducer);
  store.subscribe(() => console.log(store.getState()));
  store.dispatch({type: 'INCREMENT'});
  store.dispatch({type: 'INCREMENT'});
  store.dispatch({type: 'INCREMENT'});
  store.dispatch({type: 'DECREMENT'});
}
