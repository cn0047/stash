RPC (Remote Procedure Call)
-

RPC - interprocess communication technique.
May work over TCP, HTTP, etc.

Problems:

* How should the client react if there are no servers running or RPC server is down for a long time?
* Should a client have some kind of timeout for the RPC?
* If the server malfunctions and raises an exception, should it be forwarded to the client?
* Protecting against invalid incoming messages (eg checking bounds, type) before processing.
