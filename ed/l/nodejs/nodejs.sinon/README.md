sinon
-

Sinon - spies, stubs and mocks.

````js
// Perfect for callbacks, to check was cb called and how many times etc.
sinon.spy();

sinon.useFakeXMLHttpRequest();

sinon.stub();

sinon.mock(myAPI);

// Causes Sinon to replace the global:
// setTimeout, clearTimeout, setInterval, clearInterval, setImmediate, clearImmediate and Date
// with a custom implementation which is bound to the returned clock object.
clock = sinon.useFakeTimers();

// Simplifies cleanup.
sinon.sandbox.create();

// Restores all fake methods of supplied object
sinon.restore(obj);
````
