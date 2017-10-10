const fs = require('fs');
const { promisify } = require('util');
const xmlParse = require('xml-parse');

const readFilePromised = promisify(fs.readFile);

class Reader {
  static async getXml () {
    return readFilePromised(__dirname + '/e.xml', { encoding: 'utf8' });
  }
  
  static async get () {
    return xmlParse.parse(await this.getXml())
  }
}

(async () => {
  console.log(require('util').inspect(await Reader.get(), false, null));
})();
