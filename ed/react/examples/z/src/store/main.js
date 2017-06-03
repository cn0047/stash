import thunk from 'redux-thunk';
import { createStore, applyMiddleware, compose } from 'redux';

import reducer from '../reducers/main';

const store = defaultState => {
  return createStore(
    reducer,
    defaultState,
    compose(applyMiddleware(thunk))
  );
};

export default store
