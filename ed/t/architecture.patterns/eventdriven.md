EDA (Event-driven Architecture)
-

[event catalog](https://www.eventcatalog.dev/)

Event-driven architecture - architectural style, that async reacts to events triggered in a system.

Event - wrapper for information that something that has happened.
Event processor - handler which reacts to event.

Event notification - service1 instead of sending request to service2 directly, sends event into queue.

Event-carried state transfer - when service2 saves data from other events, and has own state, to not query service1 on next event.
It comes with data lag and consistency problems.

EDA is difficult & expensive to implement, more over it's not that simple to test.
Also events might be processed not in chronological order (due to async nature).
Also all events should be discoverable.

Also following rule "producer should know nothing about consumers" might introduce issues with events updating/extending/deprecating.
