React
-

v15.4.0

React is a declarative, efficient, and flexible JavaScript library for building user interfaces.

````
npm install -g create-react-app
create-react-app hello-world
cd hello-world
npm start
````

All React components must act like pure functions with respect to their props.
Functions are called "pure" because they do not attempt to change their inputs, and they always deterministic.

Elements inside the `map()` call need keys.
Keys used within arrays should be unique among their siblings.

Figure out is it state.
Simply ask three questions about each piece of data:
* Is it passed in from a parent via props? If so, it probably isn't state.
* Does it remain unchanged over time? If so, it probably isn't state.
* Can you compute it based on any other state or props in your component? If so, it isn't state.

Treat `this.state` as if it were immutable.

Spread Attributes:
````
function App2() {
  const props = {firstName: 'Ben', lastName: 'Hector'};
  return <Greeting {...props} />;
}
````

To pass data through the component tree without having to pass props down into each component - use `context`.

Uncontrolled Component - (use a `ref`) gets form values from the DOM.

Reconciliation - all path from virtual DOM to the actual DOM.

Conditionally render React elements: `{showHeader && <Header />}`.

Convert false, true, null, undefined to string: `<div>My JavaScript variable is {String(myVariable)}.</div>`.

Most of the time, you can use `React.PureComponent` instead of writing your own `shouldComponentUpdate`.

Not Mutating Data:
````
function updateColorMap(colormap) {
  return Object.assign({}, colormap, {right: 'blue'});
  // OR in ECS6
  return {...colormap, right: 'blue'};
}
````

[The Component Lifecycle](https://facebook.github.io/react/docs/react-component.html#the-component-lifecycle):

Mounting:
`constructor -> componentWillMount -> render -> componentDidMount`

Updating:
`componentWillReceiveProps -> shouldComponentUpdate -> componentWillUpdate -> render -> componentDidUpdate`

Unmounting:

`componentWillUnmount`

Methods prefixed with *will* are called right before something happens,
and methods prefixed with *did* are called right after something happens.

`ReactDOM.findDOMNode(node)`.

All Supported HTML Attributes available
[here](https://facebook.github.io/react/docs/dom-elements.html#all-supported-html-attributes).

[SyntheticEvent](https://facebook.github.io/react/docs/events.html#supported-events).

[Wall](https://github.com/cn007b/wall/blob/master/wall/src/web/js/implementation/react/babel/app.babel)
