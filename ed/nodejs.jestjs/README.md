Jest
-

````sh
node_modules/.bin/jest --forceExit -i -t name-of-spec
node_modules/.bin/jest -t test/unit/middlewares/
node_modules/.bin/jest -t xPoweredBy.test.js
node_modules/.bin/jest --coverage=false
````

Jest - Jasmine-based js testing framework.

In jest all modules mocked by default.

`toBe` uses === to test exact equality.
If you want to check the value of an object, use `toEqual` instead.

[Expect](https://facebook.github.io/jest/docs/en/expect.html#content).

````js
const cb = jest.fn();
expect(tpl).toBe('index');
````

`__tests__` - Tests.
`__mocks__` - Manual mocks.

`jest.runAllTimers()` - for normal work of `setTimeout`.

`waitsFor(() => {})` - for async tests.

package.json:

````js
  "jest": {
    "collectCoverage": true,
    "coveragePathIgnorePatterns": [
      "src/some/dir"
    ],
    "coverageThreshold": {
      "global": {
        "lines": 75
      }
    },
    "globals": {
      "app": null
    },
    "setupTestFrameworkScriptFile": "<rootDir>/__tests__/_setup.js",
    "testEnvironment": "node",
    "testPathIgnorePatterns": [
      "<rootDir>/__tests__/_setup.js"
    ],
    "testResultsProcessor": "jest-junit"
  },
  "jest-junit": {
    "suiteName": "My tests",
    "output": "./junit.xml",
    "classNameTemplate": "{classname}/",
    "titleTemplate": "{title}",
    "ancestorSeparator": " › ",
    "usePathForSuiteName": "true"
  }
````
