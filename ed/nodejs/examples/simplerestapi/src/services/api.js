const mongo = require('./../services/mongo');

/**
 * Gets all documents from certain collection.
 *
 * @param {string} collectionName Mongo db collection name, as usual resource name.
 * @param {function} cb Callback function which will receive data from db.
 */
module.exports.get = function(collectionName, cb) {
  mongo.getConnection().collection(collectionName, function(err, collection) {

    if (err) {
      cb({error: 'DB error 1.'});
      return;
    }
    collection.find({}, function(err, docs) {
      if (err) {
        cb({error: 'Documents not found.'});
        return;
      }
      docs.toArray().then(function(docs) {
        cb(docs);
      });
    });

  });

};

/**
 * Gets particular resource by id.
 *
 * @param {string} collectionName Mongo db collection name, as usual resource name.
 * @param {int} documentId Mongo db document id, as usual particular resource id.
 * @param {function} cb Callback function which will receive data from db.
 */
module.exports.getById = function(collectionName, documentId, cb) {

  mongo.getConnection().collection(collectionName).findOne({'id': documentId}).then(function(doc) {
    if (!doc) {
      cb({error: 'Document not found by id.'});
      return;
    }
    cb(doc);
  });

};
