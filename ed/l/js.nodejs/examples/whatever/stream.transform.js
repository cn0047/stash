const { PassThrough, Transform } = require('stream');
const { request } = require('http');

class MyTransform extends Transform {
    constructor(header = '') {
        super({ objectMode: true });

        this.header = header;
        this.isHeaderSent = false;
    }

    _transform (data, enc, callback) {
        if (!this.isHeaderSent) {
            this.push(this.header+'\n');
            this.isHeaderSent = true;
        }
        this.push(data.value.toString() + '\n');

        callback();
    }
}

function getStream() {
    const stream = new PassThrough({ objectMode: true });
    const processor = async () => {
        for (let i = 0; i < 10; i++) {
            stream.write({ value: i });
        }
        stream.end();
    };

    processor();

    return stream;
}

async function processStream(stream) {
    const headers = {};
    const options = {host: '', path: '/v1/upload', method: 'PUT', headers: headers, form: {'file': 'fn.csv'}}
    const req = request(options, () => {});
    let body = '';
    stream.on('data', (chunk) => body += chunk);
    stream.on('end', () => req.end(body));
}

async function main() {
    const stream = getStream();
    const wrapper = new MyTransform('Number');
    stream.pipe(wrapper);
    await processStream(wrapper);
}

main();
