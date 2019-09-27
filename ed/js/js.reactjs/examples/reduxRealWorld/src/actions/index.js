export const loadUser = () => (dispatch, getState) => {
  fetch('https://api.myjson.com/bins/80ij1')
    .then(res => res.json())
    .then(items => dispatch({type: 'LU', u:items}))
  ;
};
