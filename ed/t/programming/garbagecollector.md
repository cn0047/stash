Garbage Collector (GC)
-

<img src="https://gist.github.com/cn0047/384d6938ebef985347b29c15476b55c5/raw/7071e67fad3938045037e7ce92db65b2c4dab3f9/memory.jpeg" width="70%" />

#### Mark and sweep (tricolor)

Used in: go.

In this method, the GC scans the program, starting from a point called root
and tries to reach other objects and mark them.
After that performs deletion non-market objects (sweep).

Disadvantages: entire system must be suspended (stop-the-world) during collection.

#### Copying collection

In a Copying Collection collector memory is divided in two; only one half is used at a time.
When used half is full, copy used blocks to the other location, and erase the old one.

Disadvantages:
1. Has to stop the program's execution to move objects from one part of the heap to another.
2. Has to change the addresses.

Biggest advantage: if the size of the memory used by the program is less
than the size of the half of the heap that is being used,
no copying will be necessary, thus avoiding stoping the program and changing objects's addresses.

#### Reference counting

Used in: php.

In a Reference Counting collector each block has a counter of heap links to it.
This counter is incremented when a heap link is copied, decremented when the link is discarded.
When the counter goes to zero, the block is freed.
Reference counting guarantees that objects are destroyed as soon as they become unreachable.

Disadvantages:
1. Extra space is used to store the reference counter.
2. If two or more objects refer to each other, they can create a cycle where neither will be collected.

Biggest advantage: the program does not need to stop in order to perform garbage collection.
