Standards
-

* Always Use Asynchronous Methods.

* 1st callback's parameter must be error.

* Always check for errors in callbacks.

`if (err) return cb(err, null);`

* Return on callbacks.

* Validate that Callbacks are Callable.

* Use try-catch in sync code only (it won't work for callbacks!).

* Try to avoid this and new.

Binding to a specific context in Node is not a win...

* Create small modules.

* Use good async patterns.

* Error handling.

Operational errors - timeout, out of memory, failed to connect to a remote service ...<br>
Programmer errors - called an async function without a callback, cannot read property of undefined, etc.<br>
Use a logging library to increase errors visibility.<br>
Monitor your applications.<br>

* Ensure your app automatically restarts.

* Cluster your app to improve performance and reliability.

* Start a new project with npm init.

Specify a start and test script.

* Environment variables.

* Remove `X-Powered-By` header.

* Do not reinvent the wheel.

Always look for existing solutions first.

* Never `require` Modules Inside of Functions.
