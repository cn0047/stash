const util = require('util');
const fs = require('fs');

const stat = util.promisify(fs.stat);

stat('.').then((data) => {
  console.log(data);
}).catch((error) => {
  console.error('ERROR:', error);
});
// OR
(async () => {
  console.log(await stat('.'));
})();
