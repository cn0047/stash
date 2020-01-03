// npm install algoliasearch --save
const algoliasearch = require('algoliasearch');

const client = algoliasearch('{appId}', '{apiKey}');
const index = client.initIndex('my_index');

module.exports = { index };
