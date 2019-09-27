import React, { Component } from 'react'
import { connect } from 'react-redux'
import { loadUser } from '../actions'
import User from '../components/User'

const loadData = ({ loadUser }) => {
  loadUser()
};
class UserPage extends Component {
  componentWillMount() {
    loadData(this.props)
  }
  render() {
    const { user } = this.props;
    if (!user) {
      return <p>Loading...</p>
    }
    return ( <div><User user={user} /></div> )
  }
}
const mapStateToProps = (state, ownProps) => {
  return {user: state.entities.users['007']}
};
export default connect(mapStateToProps, { loadUser })(UserPage)
