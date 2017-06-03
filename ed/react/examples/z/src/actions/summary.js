import config from './../config/main';

export const loadFeed = (me) => (dispatch) => {
    let from = 0;
    let limit = 50;
    let url = config.apiUrl + '/users?online=1&pictures=1&from=' + from + '&limit=' + limit;
    let request = new Request(url, {
      headers: new Headers({'X-AUTH-TOKEN': me.token})
    });
    fetch(request)
      .then(response => response.json())
      .then(json => dispatch({type: 'SUMMARY_AFTER_LOAD_FEED', payload: json}))
    ;
};
