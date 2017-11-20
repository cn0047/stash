/*

docker run -it --rm -p 6379:6379 --hostname xredis --name xredis redis:latest

*/

const client = require('redis').createClient();

client.set('some key', 'some val');
client.get('some key', (err, reply) => {
  console.log(reply);
});

// TRANSACTION:
client.multi()
  .set('key A', 'some value A')
  .set('key B', 'some value B')
  .exec((err, data) => {
    if (err) {
      console.log('FAIL');
      throw err;
    }
    console.log(data);
  });

const Redis = require('ioredis');
const redis = new Redis();

redis.set('foo', 'bar');
redis.get('foo', (err, result) => {
  console.log(result);
});

// PIPELINE:
redis.pipeline()
  .set('foo', 'barrr')
  .set('boo', 'baz')
  .exec((err, results) => {
    if (err) throw err;
    console.log(results);
  })
;
// TRANSACTION:
redis.multi()
  .get('foo')
  .get('boo')
  .exec((err, results) => {
    if (err) throw err;
    console.log(results);
  })
;