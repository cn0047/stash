Hexagonal Architecture
-

Hexagonal Architecture known as Ports and Adapters.
Adapter - initiate interaction with application through port (REST, WS, GRPC).
Port - technology-agnostic entry point into application.
Primary/Driving port/adapter - on the left side of the hexagon.
Secondary/Driven port/adapter - on the right side.
Application - user interface and infrastructure layer.

Advantages:
* Easy to change BL in application, no need to change infrastructure layer.
* Easy to change infrastructure layer, no need to change BL.

Disadvantages:
* Simple CRUD requires extra effort.
