const mongodb = require('mongodb');

var connection;

module.exports = {

  /**
   * Performs initialization mongo db connection.
   */
  init: function() {
    mongodb.MongoClient.connect(process.env.MONGODB_URI, function (err, db) {
      if (err) {
        console.error(err);
        return;
      }
      connection = db;
    });
  },

  /**
   * Gets mongo db connection.
   *
   * @returns {object} Mongo connection object.
   */
  getConnection: function() {
    return connection;
  }

};
