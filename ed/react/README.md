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

https://facebook.github.io/react/docs/jsx-in-depth.html
