db.createCollection(<name>, {
  capped: <boolean>,
  autoIndexId: <boolean>,
  size: <number>,
  max: <number>,
  storageEngine: <document>,
  validator: <document>,
  validationLevel: <string>,
  validationAction: <string>,
  indexOptionDefaults: <document>,
  viewOn: <string>,
  pipeline: <pipeline>,
  collation: <document>,
  writeConcern: <document>
});
