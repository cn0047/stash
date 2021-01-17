Messaging patterns
-

Messaging is a technology that enables
high-speed, asynchronous, program-to-program communication with reliable delivery.

In essence, a message is transmitted in five steps:
1. Create — The sender creates the message and populates it with data.
2. Send — The sender adds the message to a channel.
3. Deliver — The messaging system moves the message from the sender’s computer
   to the receiver’s computer, making it available to the receiver.
4. Receive — The receiver reads the message from the channel.
5. Process — The receiver extracts the data from the message.

Messaging pattern - is a network-oriented architectural pattern
which describes how two different parts of a Message-Passing-System
communicate with each other.
(You don't have to say what to do, you have to send message and system will decide what to do).

Specific benefits of messaging:
* Remote Communication.
* Platform/Language Integration (independency).
* Asynchronous Communication.
* Variable Timing (external API may respond 1 or 3 or 5 seconds).
* Reliable Communication.
* Disconnected Operation (When offline -> put in queue, when online -> consume).
* Mediation (message broker behaves like brain).

#### Message Exchange Patterns:

* **In-Only** (one-way) - send message receive status.

* **Robust In-Only** (reliable one-way) - send message receive status,
  if status `false` - return status.

* **In-Out** (request–response or standard two-way) - send message receive message.

* **In-Optional-Out** - response is optional.

* **Out-Only** (reverse In-Only) - supports event notification.

* **Robust Out-Only** - like out-only but it can trigger a fault message.

* **Out-In** (reverse In-Out) - consumer initiates the exchange.

* **Out-Optional-In** (reverse In-Optional-Out) - incoming message is optional (Optional-in).

#### In ØMQ:

* **Request–reply** - RPC

* **Publish–subscribe** (data distribution pattern) - connects
  a set of publishers to a set of subscribers.

* Push–pull (**fan-out/fan-in**) - parallel task distribution
  and don't wait for any response.

* **Exclusive pair** - connects two sockets in an exclusive pair.
