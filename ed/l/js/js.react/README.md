React JS
-

<br>v16.8.6
<br>v15.4.0

[Supported HTML Attributes](https://facebook.github.io/react/docs/dom-elements.html#all-supported-html-attributes)
[SyntheticEvent](https://facebook.github.io/react/docs/events.html#supported-events)
[Wall](https://github.com/cn007b/wall/blob/master/wall/src/web/js/implementation/react/babel/app.babel)

React is a declarative, efficient, and flexible JS library for building user interfaces.

````sh
<script src="https://unpkg.com/react@latest/dist/react.js"></script>
<script src="https://unpkg.com/react-dom@latest/dist/react-dom.js"></script>
# 2019
<script src="https://unpkg.com/react@16.7.0/umd/react.production.min.js"></script>
<script src="https://unpkg.com/react-dom@16.7.0/umd/react-dom.production.min.js"></script>
# 2021
<script crossorigin src="https://unpkg.com/react@17/umd/react.development.js"></script>
<script crossorigin src="https://unpkg.com/react-dom@17/umd/react-dom.development.js"></script>

npm install -g create-react-app
create-react-app hello-world
cd hello-world
npm start
````

Component may be:
* function component (`const MyCmpnt = () => { return someJSX }`).
* class component.

Use class component when you need: to work with state, refs, lifecycle events, child functions.
Everywhere else use function component (stateless component) (`React.FunctionComponent`).

Container (aka controller-view):
* no markup.
* passes data and actions down.
* knows about redux.
* stateful.

Presentation component (aka view):
* no business logic just markup.
* receive data and actions via props.
* doesn't know about redux.
* functional component.

Higher Order Component (HOC) - function that takes component and returns new component.

All React components must act like pure functions with respect to their props.
Functions are called "pure" because they do not attempt to change their inputs,
and they always deterministic.

Elements inside the `map()` call need keys.
Keys used within arrays should be unique among their siblings.

To figure out whether something is state - ask next questions about each piece of data:
* is it passed in from a parent via props? If so - it isn't state.
* does it remain unchanged over time? If so - it isn't state.
* can you compute it based on any other state or props in your component? If so - it isn't state.

Treat `this.state` as if it were immutable.

Spread attributes:
````js
function App2() {
  const props = {firstName: 'Ben', lastName: 'Hector'};
  return <Greeting {...props} />;
}
<div> {props.cards.map(card => <Card {...card} />)} </div>
````

Fragment - group list of children without adding extra nodes to the DOM.
````js
return (
  <dl>
    {props.items.map(item => (
      // Fragments should also have a `key` prop when mapping collections
      <Fragment key={item.id}>
        <dt>{item.term}</dt>
        <dd>{item.description}</dd>
      </Fragment>
    ))}
  </dl>
);
// short syntax
return (
  <>
    <dt>{item.term}</dt>
    <dd>{item.description}</dd>
  </>
);
````

Portal - insert child into different location in DOM
outside the DOM hierarchy of the parent component.
````js
render() {
  return ReactDOM.createPortal(this.props.children, domNode);
}
````

To pass data through the component tree without having to pass props
down into each component - use `context`.
````js
const MyContext = React.createContext(defaultValue);
````

By adding `childContextTypes` and `getChildContext` to component
React passes the information down automatically
and any component in the subtree can access it by defining `contextTypes`.

Changes in context won't cause re-render not in parent nor in child.
`component.forceUpdate()` - call `render` skipping `shouldComponentUpdate`,
but trigger the normal lifecycle methods for child.

Uncontrolled Component - (use a `ref`) gets values from the DOM.
````js
return <input ref="inp" />;
componentDidMount: () => {
  console.log(this.refs.inp.getDOMNode().value);
}
````

Reconciliation - all path from virtual DOM to the actual DOM.

Conditionally render React elements: `{showHeader && <Header />}`.

Convert `false, true, null, undefined` to string: `<div>My JavaScript variable is {String(myVariable)}.</div>`.

Most of the time, you can use `React.PureComponent` instead of writing your own `shouldComponentUpdate`.

Difference between `React.PureComponent` and `React.Component`
`PureComponent` does a shallow comparison on state change.
It means that when comparing scalar values it compares their values,
but when comparing objects it compares only references (helps to improve performance).
<br>Use `React.PureComponent` when: state/props immutable, state/props should not have a hierarchy.
<br>All child components to `React.PureComponent` must also be `React.PureComponent`.

Not mutating data:
````js
this.setState(prevState => Object.assign({}, prevState, {new: "new"}));

function updateColorMap(colormap) {
  return Object.assign({}, colormap, {right: 'blue'});
  // OR in ECS6
  return {...colormap, right: 'blue'};
}
````

[The Component Lifecycle](https://facebook.github.io/react/docs/react-component.html#the-component-lifecycle):

Mounting:
* constructor.
* componentWillMount.
* render.
* componentDidMount.

Updating:
* componentWillReceiveProps.
* shouldComponentUpdate.
* componentWillUpdate.
* render.
* componentDidUpdate.

Unmounting:
* componentWillUnmount.

`componentDidCatch(error, info)`.

Methods prefixed with `will` are called right before something happens,
and methods prefixed with `did` are called right after something happens.

`ReactDOM.findDOMNode(node)`.

Hooks let you use state and other React features without writing a class.

Basic Hooks:
* useState - `[countValue, setCount] = useState(0 /* init count value */);`.
* useEffect - similar to componentDidMount and componentDidUpdate.
* useContext.

Additional Hooks:
* useReducer.
* useCallback - caches function definition between re-renders.
* useMemo - will only recompute memoized value when one of dependencies has changed.
* useRef.
* useImperativeHandle.
* useLayoutEffect.
* useDebugValue.

````js
var Settings = React.createClass({
  statics: {
    willTransitionTo: function (transition, params, query, callback) {
      callback();
    },
    willTransitionFrom: function (transition, params, query, callback) {
      callback();
    },
  }
});

Config.propTypes = {
  data: PropTypes.shape({
    type: PropTypes.string,
    key: PropTypes.string,
    applicationID: PropTypes.string,
    value: PropTypes.string,
    description: PropTypes.string,
  }),
  scenario: PropTypes.oneOf([SCENARIO_CREATE, SCENARIO_UPDATE]).isRequired,
};
````

## Server side rendering

````js
import ReactDOMServer from 'react-dom/server';
const = ssRender = () => {
  return ReactDOMServer.renderToString(
    <App />
  );
};
````
