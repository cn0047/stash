import config from './../config/main';

export const logIn = () => (dispatch) => {
  let rb = JSON.stringify(config.auth);
  return fetch(config.apiUrl + '/auth', {method: 'post', body: rb})
    .then(response => response.json())
    .then(json => dispatch({type: 'ME_AFTER_LOG_IN', payload: json}))
  ;
};
