mocha
-

````json
"scripts": {
  "test": "mocha --reporter spec",
  "cover": "node_modules/istanbul/lib/cli.js cover node_modules/mocha/bin/_mocha -- -R spec test/*"
}
````

````
open coverage/lcov-report/*.html
````

````js
it.skip
it.only
````
