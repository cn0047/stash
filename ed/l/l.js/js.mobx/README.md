mobx
-

[state](https://www.npmjs.com/package/mobx)

````js
import { observable, computed, configure, action } from 'mobx';

@observable name = 'bond';
// IMPORTANT: use null instead of {}
@observable devices = null;

@observer class SecretAgent extends Agent {}

@computed get UniqueID() {}

// observable object
const myObj = observable({observableProperty: ... }, {increment: action});

@inject('myStore', 'addStore')
class Main extends React.Component {}
````
