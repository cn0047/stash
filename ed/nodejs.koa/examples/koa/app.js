/**
 * @example node --harmony app.js
 */

const koa = require('koa');
const app = new koa();
const route = require('koa-route');
const parse = require('co-body');

// app.use(function* () {
//   this.body = 'HW';
// });

app.use

app.listen(3000);
console.log('Started on port 3000.');
