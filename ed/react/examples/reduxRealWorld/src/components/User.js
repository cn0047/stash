import React from 'react'
import PropTypes from 'prop-types'

const User = ({ user }) => {
  return (
      <img src={user.avatarUrl} alt="" width="72" height="72" />
  )
};
User.propTypes = {
  user: PropTypes.shape({
    avatarUrl: PropTypes.string.isRequired,
  }).isRequired
};
export default User
