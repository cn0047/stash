Race Conditions
-

Race condition - when events do not happen in the order the programmer intended.

Integer overflow (225+1=0; -128-1=127).
Stack overflow.
Heap overflow.

Don't handle `Unix signals`.
Don't specify timeout in microservice SDK HTTP client.

When memcache expires data for heavy query - few threads may run heavy query simultaneously.

#### Dates

* Leap second.
* Ambiguous time.
* Daylight saving time.
* Nepal Time UTC/GMT+05:45.

Always benchmark in nanoseconds it's more accurate that milliseconds and microseconds.

#### MySql

Optimistic Locking.

#### PHP

UniSender threads (crc32)...

#### NODEJS

````
Using fs.access() to check for the accessibility of a file before calling fs.open(), fs.readFile() or fs.writeFile()
is not recommended. Doing so introduces a race condition, since other processes may change the file's state between the two calls.
Instead, user code should open/read/write the file directly and handle the error raised if the file is not accessible.
````
