import React from 'react';
import ReactDOM from 'react-dom';
import { combineReducers, createStore } from 'redux'

ReactDOM.render(<div>Look into console.</div>, document.getElementById('root'));

function counter(state = 0, action) {
  switch (action.type) {
  case 'INCREMENT':
    return state + 1
  case 'DECREMENT':
    return state - 1
  default:
    return state
  }
}
let store = createStore(counter);
store.subscribe(() => console.log(store.getState()));
store.dispatch({type: 'INCREMENT'});
store.dispatch({type: 'INCREMENT'});
store.dispatch({type: 'INCREMENT'});
store.dispatch({type: 'DECREMENT'});

console.log('Example 2:');

function blackhole(state = [], action) {
  switch (action.type) {
  case 'ADD':
    return [...state, action.text];
  default:
    return state;
  }
}
let reducer = combineReducers({counter, blackhole});
let store2 = createStore(reducer)
store2.subscribe(() => console.log(store2.getState().blackhole));
store2.dispatch({type: 'ADD', text: 'New text.'});
store2.dispatch({type: 'ADD', text: 'One more new text.'});
