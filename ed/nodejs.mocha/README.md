mocha
-

````json
"scripts": {
  "test": "mocha",
  "test": "mocha --reporter spec --compilers js:babel-core/register",
  "cover": "node_modules/istanbul/lib/cli.js cover node_modules/mocha/bin/_mocha -- -R spec test/*",
  "cover": "node_modules/istanbul/lib/cli.js cover node_modules/mocha/bin/_mocha -- --compilers js:babel-core/register -R spec test/*",
}
````

````sh
open coverage/lcov-report/*.html
````

````js
it.skip
it.only
this.retries(2);
this.timeout(500);
````

Reporters: spec, dot, nyan, landing, list, progress, json, min, doc, markdown, html
