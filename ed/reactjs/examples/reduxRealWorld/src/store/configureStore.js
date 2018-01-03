import { createStore, applyMiddleware, compose } from 'redux'
import thunk from 'redux-thunk'

import rootReducer from '../reducers'

const configureStore = preloadedState => {
  return createStore(
    rootReducer,
    preloadedState,
    compose(applyMiddleware(thunk))
  );
};
export default configureStore
