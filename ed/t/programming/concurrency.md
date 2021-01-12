Concurrency
-

Mutex - (locking mechanism) mutual exclusion object that synchronizes access to a resource.

Semaphore - (signalling mechanism):
* counting semaphore - count is the number of available resources.
* binary semaphore - like counting semaphore but value restricted to 0 and 1.

Starvation - situation where a concurrent process cannot get all the resources
they needs to perform work.

Types:
* Parallel.
* Asynchronous.

#### Locks

Distributed Lock.

Deadlock - concurrent processes are waiting on one another.

Livelock - when two people meet in a narrow corridor,
and each tries to be polite by moving aside to let the other pass,
but they end up swaying from side to side without making any progress
because they both repeatedly move the same way at the same time.

As with deadlock, livelocked threads are unable to make further progress.
However, the threads are not blocked — they are simply too busy responding to each other to resume work.
