redux
-

[Examples](http://redux.js.org/docs/introduction/Examples.html),
[ToDo](http://redux.js.org/docs/basics/ExampleTodoList.html),
[RealWorld](https://github.com/reactjs/redux/tree/master/examples/real-world).
[ReactRouter](http://redux.js.org/docs/advanced/UsageWithReactRouter.html).
[ChromeExtension](https://chrome.google.com/webstore/detail/redux-devtools/lmhkpmbekcpmknklioeibfkpmmfibljd?hl=en)

````sh
npm install --save redux react-redux
redux-devtools
redux-thunk
redux-promise-middleware
````

To bind react component state to redux state use `connect` method from `react-redux` package:

````
/**
 * @param {Object} mapStateToProps Object which describes what map from redux to component.
 * @param {Object} mapDispatchToProps Object which describes actions to dispatch.
 */
connect(mapStateToProps, mapDispatchToProps);
````

Thunk: functions, clunky to test, easy to learn.
Saga: generators, easy to test, hard to learn.
