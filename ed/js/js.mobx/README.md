mobx
-

[state](https://www.npmjs.com/package/mobx)

````js
import { observable, computed, configure, action } from 'mobx';

@observable name = 'bond';

@observer class SecretAgent extends Agent {}

@computed get UniqueID() {}

// observable object
const myObj = observable({observableProperty: ... }, {increment: action});
````
