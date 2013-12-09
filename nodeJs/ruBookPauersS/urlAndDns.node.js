var dns = require('dns');
dns.lookup('burningbird.net',function(err,ip) {
    if (err) throw err;
    console.log(ip);
});
dns.reverse('173.255.206.103', function(err,domains) {
    domains.forEach(function(domain) {
        console.log(domain);
    });
});
var dns = require('dns');
dns.resolve('burningbird.net', 'NS', function(err,domains) {
    domains.forEach(function(domain) {
        console.log(domain);
    });
});

var url = require('url');
var urlObj = url.parse('http://examples.burningbird.net:8124/?file=main');
/*
Result: {
    protocol: 'http:',
    slashes: true,
    host: 'examples.burningbird.net:8124',
    port: '8124',
    hostname: 'examples.burningbird.net',
    href: 'http://examples.burningbird.net:8124/?file=main',
    search: '?file=main',
    query: 'file=main',
    pathname: '/',
    path: '/?file=main'
}
*/

console.log(url.format(urlObj));
// Result: http://examples.burningbird.net:8124/?file=main

var vals = querystring.parse('file=main&file=secondary&type=html');
console.log(vals);
// Result: { file: [ 'main', 'secondary' ], type: 'html' }