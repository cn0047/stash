import { combineReducers } from 'redux';
import { routerReducer as routing } from 'react-router-redux'

import me from './me';
import summary from './summary';

const reducer = combineReducers({
  me,
  summary,
  routing
});

export default reducer;
