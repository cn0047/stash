const fs = require('fs');
const { promisify } = require('util');
const parser = require('xml2json');

const readFilePromised = promisify(fs.readFile);

class Reader {
  static async getXml () {
    return readFilePromised(__dirname + '/e.xml', { encoding: 'utf8' });
  }
  
  static async get () {
    return parser.toJson(await this.getXml());
  }
}

(async () => {
  console.log(require('util').inspect(await Reader.get(), false, null));
})();
