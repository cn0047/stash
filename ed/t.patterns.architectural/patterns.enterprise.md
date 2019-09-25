Enterprise Application Architecture Patterns
-

1. Domain Logic Patterns:
    * Transaction Script
    * Domain Model
    * Table Module
    * Service Layer

2. Data Source Architectural Patterns:
    * Table Data Gateway
    * Row Data Gateway
    * Active Record
    * Data Mapper

3. Object-Relational Behavioral Patterns:
    * Unit of Work
    * Identity Map
    * Lazy Load

4. Object-Relational Structural Patterns:
    * Identity Field
    * Foreign Key Mapping
    * Association Table
    * Dependent Mapping
    * Embedded Value
    * Serialized LOB
    * Single Table Inheritance
    * Class Table Inheritance
    * Concrete Table Inheritance
    * Inheritance Mappers

5. Object-Relational Metadata Mapping Patterns:
    * Metadata Mapping
    * Query Object
    * Repository

6. Web Presentation Patterns:
    * Model View Controller
    * Page Controller
    * Front Controller
    * Template View
    * Transform View
    * Two Step View
    * Application Controller

7. Distribution Patterns:
    * Remote Facade
    * Data Transfer Object

8. Offline Concurrency Patterns:
    * Optimistic Offline Lock
    * Pessimistic Offline Lock
    * Coarse-Grained Lock
    * Implicit Lock

9. Session State Patterns:
    * Client Session State
    * Server Session State
    * DB Session State

10. Base Patterns:
    * Gateway
    * Mapper
    * Layer Supertype
    * Separated Interface
    * Registry
    * Money
    * Special Case
    * Plugin
    * Service Stub
    * Record Set
    * Value Object

## Domain Logic Patterns:

**Transaction Script**
Keep business logic as stored procedures, so just use DB.
Good for apps with small amount of logic. Introduce very little DB overhead in performance and in understanding.

**Domain Model**
Requires Data Mapper.
A rich Domain Model (complex logic) can look different from the DB design,
with inheritance, strategies, and other patterns.

**Table Module**
Given many orders, a Domain Model will have one order object per order
while a Table Module will have one object to handle all orders.

**Service Layer**
Layer of services with set of available operations. 

## Data Source Architectural Patterns:

**Table Data Gateway**
Holds all SQL for accessing a single table.
The simplest DB interface pattern to use (maps easy to table).

**Row Data Gateway**
Object that exactly mimics a single record (DB row). 
should contain only DB access logic and no domain logic (Active Record).

**Active Record**
Object that wraps a row in a DB table or view,
encapsulates the DB access, and adds domain logic on that data.
Good choice for domain logic that isn’t too complex.

**Data Mapper**
A layer of mappers that moves data between objects and a DB
while keeping them independent of each other and the mapper itself.
<br>
(Object Relational Mapper (ORM) and the DB Abstraction Layer (DAL)).
Object-relational mapping - is a programming technique
for converting data between incompatible type systems in object-oriented programming languages.

## Object-Relational Behavioral Patterns:

**Unit of Work**
Maintains a list of objects affected by business transaction and coordinates writing changes and the resolve concurrency problems.
If every transaction uses the same sequence of tables to edit, you greatly reduce the risk of dead-locks.

**Identity Map**
Ensures that each object gets loaded only once by keeping every loaded object in a map.

**Lazy Load**
An object that doesn’t contain all of the data you need but knows how to get it.

## Object-Relational Structural Patterns:

**Identity Field**
Saves a DB ID field in object to maintain identity between in-memory object and DB row.

**Foreign Key Mapping**
Maps an object reference to a foreign key in the DB, many-to-many associations.

**Association Table**
Link table with foreign keys to the tables that are linked by the association (many-to-many).

**Dependent Mapping**
When you have an object that’s only referred to by one other object,

**Embedded Value**
Maps an object into several fields of another object’s table.

**Serialized LOB**
Saves a graph of objects by serializing them into a single large object (LOB), which it stores in a DB field.

**Single Table Inheritance**
Represents an inheritance hierarchy of classes as single table that has columns for all fields of various classes.

**Class Table Inheritance**
Represents an inheritance hierarchy of classes with one table for each class.

**Concrete Table Inheritance**
Represents an inheritance hierarchy of classes with one table per concrete class in the hierarchy. 

**Inheritance Mappers**
A structure to organize DB mappers that handle inheritance hierarchies.
This general scheme makes sense for any inheritance-based DB mapping.

## Object-Relational Metadata Mapping Patterns:

**Metadata Mapping**
Holds details of object-relational mapping in metadata.
Much of the code that deals with object-relational mapping describes
how fields in the DB correspond to fields in in-memory objects.
The resulting code tends to be tedious and repetitive to write.
A Metadata Mapping allows developers to define the mappings in a simple tabular form.

**Query Object**
An object that represents a DB query.

**Repository**
Mediates between the domain and data mapping layers using a collection-like interface for accessing domain objects.

## Web Presentation Patterns:

**Model View Controller**
Splits user interface interaction into three distinct roles.
<br>
1. model — contains the core functionality and data.
2. view — displays the information to the user.
3. controller — handles the input from the user.
<br>
Usage: architecture for World Wide Web app.

**Page Controller**
An object that handles a request for a specific page or action on a Web site.

**Front Controller**
A controller that handles all requests for a Web site.

**Template View**
Renders information into HTML by embedding markers in an HTML page.

**Transform View**
A view that processes domain data element by element and transforms it into HTML.

**Two Step View**
Turns domain data into HTML in two steps:
first by forming some kind of logical page, then rendering the logical page into HTML.

**Application Controller**
A centralized point for handling screen navigation and the flow of an application.

## Distribution Patterns:

**Remote Facade**
Provides a coarse-grained facade on fine-grained objects to improve efficiency over a network.

**Data Transfer Object**
An object that carries data between processes in order to reduce the number of method calls.
Use a Data Transfer Object whenever you need to transfer multiple items of data
between two processes in a single method call. 

## Offline Concurrency Patterns:

**Optimistic Offline Lock**
Associate version number with each record in your system.
When a record is loaded that number is maintained by the session along with all other session state
and compare the version stored in your session data to the current version in the record data.

**Pessimistic Offline Lock**
Prevents conflicts between concurrent business transactions by allowing only one business transaction at a time to access data.
Forces a business transaction to acquire a lock on a piece of data before it starts to use it,

**Coarse-Grained Lock**
ACoarse-Grained Lock is a single lock that covers many objects.

**Implicit Lock**
Allows framework or layer supertype code to acquire offline locks.

## Session State Patterns:

**Client Session State**
Stores session state on the client.

**Server Session State**
Keeps the session state on a server system in a serialized form.

**DB Session State**
Stores session data as committed data in the DB.

## Base Patterns:

**Gateway**
An object that encapsulates access to an external system or resource.

**Mapper**
Customer Lease Asset An object that sets up a communication between two independent objects.

**Layer Supertype**
A type that acts as the supertype for all types in its layer.

**Separated Interface**
Defines an interface in a separate package from its implementation.

**Registry**
A well-known object that other objects can use to find common objects and services.

**Money**
Represents a monetary value.

**Special Case**
A subclass that provides special behavior for particular cases.

**Plugin**
Links classes during configuration rather than compilation.

**Service Stub**
Removes dependence upon problematic services during testing.

**Record Set**
Record Set An in-memory representation of tabular data.

**Value Object**
A small simple object, like money or a date range, whose equality isn’t based on identity.
Value Objects are small objects, such as a money object or a date,
while reference objects are large, such as an order or a customer.
<br>
The key difference between reference and value objects lies in how they deal with equality.
A reference object uses identity (primary key in a relational DB) as the basis for equality.
A Value Object bases its notion of equality on field values within the class.
Thus, two date objects may be the same if their day, month, and year values are the same.
<br>
For value objects to work properly in these cases it’s a very good idea to make them immutable.
Value Objects shouldn’t be persisted as complete records.
Treat something as a Value Object when you’re basing equality on something other than an identity.
