const { Readable, Writable } = require('stream');

class Counter extends Readable {
  constructor(opt) {
    super(opt);
    this.n = 5;
  }

  _read() {
    for (let i = 1; i <= this.n; i++) {
      const str = '# ' + i;
      const buf = Buffer.from(str, 'ascii');
      this.push(buf);
    }
    this.push(null);
  }
}

class Logger extends Writable {
  _write(chunk, encoding, callback) {
    console.log('🔰', chunk.toString());
    callback();
  }
}

const counter = new Counter();
const log = new Logger();
counter.pipe(log);
