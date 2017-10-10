const fs = require('fs');
const { promisify } = require('util');
const { Parser } = require('xml2js');

const readFilePromised = promisify(fs.readFile);
const parser = new Parser({ explicitArray: false });
const parseStringPromised = promisify(parser.parseString);

class Reader {
  static getXml () {
    return readFilePromised(__dirname + '/e.xml', { encoding: 'utf8' });
  }
  
  static async get () {
    return await parseStringPromised(await this.getXml());
  }
}

(async () => {
  console.log(require('util').inspect(await Reader.get(), false, null));
})();
