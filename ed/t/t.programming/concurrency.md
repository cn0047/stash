Concurrency
-

#### Locks

Deadlock - concurrent processes are waiting on one another.

livelock - when two people meet in a narrow corridor,
and each tries to be polite by moving aside to let the other pass,
but they end up swaying from side to side without making any progress
because they both repeatedly move the same way at the same time.

As with deadlock, livelocked threads are unable to make further progress.
However, the threads are not blocked — they are simply too busy responding to each other to resume work.
