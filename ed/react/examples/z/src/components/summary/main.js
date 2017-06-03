import React, { Component } from 'react';

import Cell1 from './cell1';

class Main extends Component {

  render() {
    let favourite = [];
    let flame = [];
    let visit = [];
    this.props.summary.relationships.forEach((r) => {
      if (r.type_id === 200902 && r.user_id === this.props.me.user_id) {
        favourite.push(r.to_user_id);
      }
      if (r.type_id === 200903 && r.user_id === this.props.me.user_id) {
        flame.push(r.to_user_id);
      }
      if (r.type_id === 200907 && r.user_id === this.props.me.user_id) {
        visit.push(r.to_user_id);
      }
    });
    let cells = [];
    this.props.summary.users.forEach((user) => {
      let f = favourite.indexOf(user.user_id);
      let fl = flame.indexOf(user.user_id);
      let v = visit.indexOf(user.user_id);
      cells.push(<Cell1 {...user} key={user.user_id} favourited={f} flamed={fl} visited={v} />);
    });
    return (
      <div>{cells}</div>
    );
  }

}

export default Main;
