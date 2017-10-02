const fs = require('fs');
const { parseString } = require('xml2js');

class Reader {
  static read () {
    const xml = fs.readFileSync('./ne.xml', { encoding: 'utf8' });
    const promise = new Promise(resolve => {
      parseString(xml, (er, json) => {
        resolve(json);
      });
    });
    promise.then(data => {
      data = data['Manager']['C'];
    });
    return promise;
  }
}

Reader.read().then(d => {
  console.log(d['Manager']['S']);
});

