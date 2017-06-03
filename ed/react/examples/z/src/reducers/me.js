const me = (state = {}, action) => {
  switch (action.type) {
    case 'ME_AFTER_LOG_IN':
      let d = action.payload;
      return {
        ...d.user,
        token: d.token,
        qb: d.quickblox
      };
    default:
      return state;
  }
};

export default me;
