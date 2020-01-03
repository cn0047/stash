var elasticsearch = require('elasticsearch');

var es = new elasticsearch.Client({
  host: 'es:9200',
  log: 'error'
});

es.search({
    index: 'megacorp',
    type: 'employee',
    body: {
      query: {ids: {values: [1]}}
    }
}).then(function (resp) {
    var hits = resp.hits.hits;
    console.log(hits[0]._source);
}, function (err) {
    console.trace(err.message);
});
