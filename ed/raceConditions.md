Race Conditions
_

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
