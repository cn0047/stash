const fs = require('fs');
const { parseString } = require('xml2js');

class Reader {
  static readSemiSync () {
    const xml = fs.readFileSync('./e.xml', { encoding: 'utf8' });
    return new Promise(resolve => {
      parseString(xml, (er, json) => {
        resolve(json);
      });
    });
  }

  static read () {
    return new Promise(resolve => {
      fs.readFile('./e.xml', (err, xml) => {
        if (err) throw err;
        parseString(xml, (er, json) => {
          if (er) throw er;
          resolve(json);
        });
      });
    });
  }
}

(async () => {
  const r = await Reader.read();
  console.log(r['Manager']['C']);
})();
