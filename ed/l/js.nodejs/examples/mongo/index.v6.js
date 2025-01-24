// "dependencies": {"mongodb": "^6.12.0"}

const MongoClient = require('mongodb').MongoClient;

const main = async () => {
  const uri = 'mongodb+srv://usr:pwd@host';
  const db = 'db';
  const client = new MongoClient(uri);
  await client.connect();
  const r = await client.db(db).collection('col').findOne({key: 'val'}, {projection: {_id: 1}});
  console.log(`Res: ${JSON.stringify(r)}`);
  return r;
}

main();
