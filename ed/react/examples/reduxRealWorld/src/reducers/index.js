import { routerReducer as routing } from 'react-router-redux'
import { combineReducers } from 'redux'

const entities = (state = { users: {} }, action) => {
  if (action.u) {
    return action.u;
  }
  return state;
};
const rootReducer = combineReducers({ entities, routing });
export default rootReducer
