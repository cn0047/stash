import React, { Component } from 'react';

import Cell1 from './Cell1';

class SummaryScreen extends Component {

  constructor(props) {
    super(props);
    this.state = {
      user: this.props.route.user,
      status: {}, paging: {}, users: [], relationships: [],
      favourite: [], flame: [], visit: [],
    };
    this.load(0);
  }

  load(from) {
    let request = new Request('https://api.ziipr.com/v1/users?online=1&pictures=1&from=' + from + '&limit=50', {
      headers: new Headers({'X-AUTH-TOKEN': this.props.route.user.token})
    });
    fetch(request).then((res) => {
      return res.json().then(d => {
        let favourite = [];
        let flame = [];
        let visit = [];
        d.relationships.forEach((r) => {
          if (r.type_id === 200902 && r.user_id === this.state.user.userId) {
            favourite.push(r.to_user_id);
          }
          if (r.type_id === 200903 && r.user_id === this.state.user.userId) {
            flame.push(r.to_user_id);
          }
          if (r.type_id === 200907 && r.user_id === this.state.user.userId) {
            visit.push(r.to_user_id);
          }
        });
        this.setState(prevState => ({
          users: prevState.users.concat(d.users),
          relationships: prevState.relationships.concat(d.relationships),
          favourite: prevState.favourite.concat(favourite),
          flame: prevState.flame.concat(flame),
          visit: prevState.visit.concat(visit)
        }));
        console.log(d.paging);
        if (d.paging.hasOwnProperty('next_from') === true) {
          this.load(d.paging.next_from);
        }
      });
    });
  }

  render() {
    if (this.state.hasOwnProperty('users') === false) {
      return (<div>Loading...</div>);
    }
    let cells = [];
    this.state.users.forEach((user) => {
      let f = this.state.favourite.indexOf(user.user_id);
      let fl = this.state.flame.indexOf(user.user_id);
      let v = this.state.visit.indexOf(user.user_id);
      cells.push(<Cell1 {...user} key={user.user_id} favourited={f} flamed={fl} visited={v} />);
    });
    return (
      <div>{cells}</div>
    );
  }

}

export default SummaryScreen;
