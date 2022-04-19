const fs = require('fs');
const { promisify } = require('util');
const toJSON = require('xmljson').to_json;

const readFilePromised = promisify(fs.readFile);
const toJSONPromised = promisify(toJSON);

class Reader {
  static async getXml () {
    return readFilePromised(__dirname + '/e.xml', { encoding: 'utf8' });
  }
  
  static async get () {
    return toJSONPromised(await this.getXml());
  }
}

(async () => {
  console.log(require('util').inspect(await Reader.get(), false, null));
})();
