npm
-

````
npm -g i npm@latest
````

````
npm help npm
npm -g install connect
npm uninstall modulename
npm update modulename

npm ls
npm ls -g

npm config list
npm config ls -l
npm config delete keyname
npm config set keyname value

npm publish

npm prune
````

````
console.log(argv.one + " " + argv.two);
./app2.js --one="My" --two="Name"
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
