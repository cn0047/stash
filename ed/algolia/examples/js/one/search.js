const b = require('./bridge');

b.index.search({query: 'leo'}, (err, content) => {
  if (err) throw err;
  console.log(content.hits);
});
