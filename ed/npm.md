npm
-

[Trends](http://www.npmtrends.com/ws-vs-socket.io)
[Counts](https://github.com/npm/registry/blob/master/docs/download-counts.md)
[Stat](https://npm-stat.com)

[Shorthands](https://docs.npmjs.com/misc/config#shorthands-and-other-cli-niceties)

````
npm -g i npm@latest
````

````
^4.13.3 # equals to 4.*.* which is same major version
~4.13.3 # equals to 4.13.* which is same minor version
````

````
npm help npm
npm -g install connect
npm uninstall modulename
npm update modulename

npm ll
npm ls
npm ls -g

npm run-script lint

npm config list
npm config ls -l
npm config delete keyname
npm config set keyname value
npm config set @qsc:registry "http://npm.qrmmlab.qsc.com/repository/npm/"

npm publish

npm prune
````

````
# npm install -g nodemon
nodemon ./server.js localhost 8080

# npm install -g express-generator
express node-express-gen

# npm install -g strongloop
slc loopback
slc loopback:model

# npm i circular-require
node_modules/.bin/circular-require ./src

# npm install morgan --save
var morgan = require('morgan');
app.use(morgan('dev'));

nsp # CLI tool to help identify known vulnerabilities

@nuxtjs/localtunnel # expose your localhost as public server
````

````
console.log(argv.one + " " + argv.two);
./app2.js --one="My" --two="Name"
````

### File

````
{
  "scripts": {
    "prestart": "will run before start",
    "start": "",
  },
  "eslintConfig": {
    "env": {
      "browser": true,
      "jest": true,
      "node": true
    },
    "globals": {
      "window": true,
      "io": true,
      "hljs": true
    },
    "extends": [
      "airbnb-base"
    ]
  },
  "eslintIgnore": [
    "src/public/*"
  ],
  "jest": {
    "collectCoverage": true,
    "coverageDirectory": ".coverage"
  }
}
````

### Colors

````
npm install colors

var colors = require('colors');
console.log('This Node kicks it!'.rainbow.underline);
console.log('We be Nodin'.zebra.bold);
console.log('rainbow'.rainbow, 'zebra'.zebra);
colors.setTheme({
    mod1_warn: 'cyan',
    mod1_error: 'red',
    mod2_note: 'yellow'
});
console.log("This is a helpful message".mod2_note);
console.log("This is a bad message".mod1_error);
````

### History

npm@5

* faster
* lockfiles
