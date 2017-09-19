const util = require('util');
const fs = require('fs');

const stat = util.promisify(fs.stat);
stat('.').then((stats) => {
  console.log(stats);
}).catch((error) => {
  // Handle the error.
});
