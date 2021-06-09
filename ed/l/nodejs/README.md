Node JS
-
*v8.2.1*
*v6.11.1*
*v0.10.26*

[doc](https://nodejs.org/en/)
[compatibility tables](http://node.green/)
[jsdoc](http://usejsdoc.org/) `node_modules/.bin/jsdoc -r src -d docs`.
[https](https://github.com/cn007b/my/blob/master/ed/nodejs.express/examples/coursera.passport/bin/www#L42)
[testing tool](https://github.com/cucumber/cucumber-js)
[testing tool](https://www.npmjs.com/package/puppeteer)

````sh
sudo ln -s /usr/bin/nodejs /usr/bin/node
````

````sh
node --inspect --inspect-brk x.js
# or
# inject debug in already runned process
kill -SIGUSR1 2128
# and open in browser chrome://inspect
````
````js
console.log(require('util').inspect('ok', false, null));
require('fs').appendFile('/tmp/debug.tmp', JSON.stringify({code: 204}) + "\n"); // tail -f /tmp/debug.tmp
process.exit(); // die;
console.time('test');
console.timeEnd('test');

process.on('exit', () => {});
process.on('uncaughtException', () => {});
process.stdin.resume();

require.main == module // script executed from shell (not required into another script)
````

````js
req.url    // Url string.
req.params // Parsed params from url.
req.params.id

// yargs - alternative to `process.argv[2]`
var argv = require('yargs')
  .usage('Usage: node $0 --l=[num] --b=[num]')
  .demand(['l','b'])
  .argv;
console.log(argv.l,argv.b);
````

#### Common info

Node.js is great for doing asynchronous I/O operations,
but when it comes to real number-crunching, it’s not that great choice.

Currently, by default v8 has a memory limit of 512mb on 32-bit systems,
and 1gb on 64-bit systems.
To increase memory use `--max_old_space_size`
`node --max-old-space-size=8192 server.js`

Boundary for node module - file. File is module.
`module.exports` - for export module.
`exports` alias to `module.exports`.

There are two types of flow control: serial and parallel.

#### Under the hood

Node architecture:
* node code (js)
* node bindings (c++)
* chrome v8 (c++)
* libuv (C)

`libuv` is the open source library that handles the thread-pool,
doing signaling and all other magic that is needed to make the asynchronous tasks work.

Javascript is a single-threaded, event-driven language (even V8 is single-threaded).

We only have **one main thread** and **one call-stack**.

Whenever you call setTimeout, http.get or fs.readFile,
Node.js sends these operations to a different thread (system kernel threads)
allowing V8 to keep executing our code.
Node also calls the callback when the counter has run down
or the IO / http operation has finished.

In case there is another request being served when the said file is read,
its callback will need to wait for the execution stack to become empty.
The limbo where callbacks are waiting for their turn to be executed
is called the **task queue** or **event queue**, or **message queue**.
*Callbacks are being called in an infinite loop whenever the main thread has finished its previous task.*

If this wasn’t enough, we actually have more then one **task queue**.
**One for microtasks** (process.nextTick, promises, Object.observe)
and **another for macrotasks** (setTimeout, setInterval, setImmediate, I/O).
After said macrotask has finished, all of the available microtasks will be processed within the same cycle.

#### Event Loop

````
   ┌───────────────────────┐
┌─>│        timers         │
│  └──────────┬────────────┘
│  ┌──────────┴────────────┐
│  │     I/O callbacks     │
│  └──────────┬────────────┘
│  ┌──────────┴────────────┐
│  │     idle, prepare     │
│  └──────────┬────────────┘      ┌───────────────┐
│  ┌──────────┴────────────┐      │   incoming:   │
│  │         poll          │<─────┤  connections, │
│  └──────────┬────────────┘      │   data, etc.  │
│  ┌──────────┴────────────┐      └───────────────┘
│  │        check          │
│  └──────────┬────────────┘
│  ┌──────────┴────────────┐
└──┤    close callbacks    │
   └───────────────────────┘
````

Each phase has a FIFO queue of callbacks to execute.
When the event loop enters a given phase, it will perform operations specific to that phase,
then execute callbacks in that phase's queue
until the queue has been exhausted or the maximum number of callbacks has executed,
after that will move to the next phase, and so on.

* timers - setTimeout(), setInterval().
* I/O callbacks.
* idle, prepare - only used internally.
* poll - retrieve new I/O events.
* check - setImmediate().
* close callbacks - `socket.on('close', ...)` etc.

Between each run of the event loop,
Node.js checks if it is waiting for any asynchronous I/O or timers and shuts down...

`process.nextTick()` it's a part of the asynchronous API,
but not technically part of the event loop.
The `nextTickQueue` will be processed after the current operation completes,
regardless of the current phase of the event loop.

Blocked Event Loop:
* I/O is blocking.
* Call that does CPU work is blocking.

#### Garbage Collector

Execution nodejs code pushes variables into **execution stack**.
Local variables are popped from the stack when the functions execution finishes.
It happens only when you work with simple values such as numbers, strings and booleans.
Values of objects, arrays and such are stored in the heap and your variable is merely a pointer to them.
If you pass on this variable, you will only pass the said pointer,
making these values mutable in different stack frames.
When the function is popped from the stack,
only the pointer to the Object gets popped with leaving the actual value in the heap.
The garbage collector is the guy who takes care of freeing up space once the objects outlived their usefulness.

Things to Keep in Mind When Using a Garbage Collector:
* performance impact - in order to decide what can be freed up, the GC consumes computing power.
* unpredictable stalls - modern GC implementations try to avoid "stop-the-world" collections.

Unset object link: `Mater = undefined`.

The heap has two main segments, the New Space and the Old Space.
Objects living in the New Space are called Young Generation.
The Old Space where the objects that survived the collector in the New Space are promoted into
- they are called the Old Generation.
Allocation in the Old Space is fast, however collection is expensive so it is infrequently performed.
V8 engine uses two different collection algorithms:

* Scavenge - fast and runs on the Young Generation
* Mark-Sweep collection - slower and runs on the Old Generation.

#### Streams

Streams are objects that let you read data from a source
or write data to a destination in continuous fashion.

* Readable − read operation.
* Writable − write operation.
* Duplex − both read and write operation.
* Transform − duplex stream where the output is computed based on input.

Each type of Stream is an EventEmitter instance and throws several events:
* data − data is available to read.
* end − there is no more data to read.
* error − error receiving or writing data.
* finish − all data has been flushed to underlying system.

Piping is a mechanism where we provide the output of one stream as the input to another stream.

#### Async/await

Is new way to write asynchronous code (before were only callbacks and promises).

The `await` keyword can only be used inside functions defined with `async`.

This approach is:
* Concise and clean.
* Error handling (`try/catch` also works for asynchronous errors).
* Easy to write conditions, easy to debug, better error stack.
