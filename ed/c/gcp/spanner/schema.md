Schema
-

[best practices](https://cloud.google.com/spanner/docs/schema-design)
[design schema](https://cloud.google.com/spanner/docs/schema-and-data-model)
[view](https://cloud.google.com/spanner/docs/views)
[secondary indexes](https://cloud.google.com/spanner/docs/secondary-indexes)
[FK](https://cloud.google.com/spanner/docs/foreign-keys/overview)
[FK](https://cloud.google.com/spanner/docs/foreign-keys/how-to)
[check constraint](https://cloud.google.com/spanner/docs/check-constraint/how-to)

view - virtual table defined by a SQL query.

Check constraint - `CONSTRAINT start_before_end CHECK(StartTime < EndTime)`.
If expression evaluates to TRUE or NULL - data change is allowed.
Expression can only reference columns in the same table.
No CK on column `allow_commit_timestamp=true`.

Parent-child relationships: interleaving or foreign keys, but not both.

Interleaved table - child table, albums is interleaved table for singers (no singer -> no album).
For interleaved table can be only 1 parent.
Parent table primary key must be part of child table primary key.
Interleaved tables helps with performance, spanner co-locate tables in the same storage layer.
`DELETE CASCADE` works only for interleaved tables.

Foreign keys needed for referential integrity.
No FK on column `allow_commit_timestamp=true` or with array or JSON type.
FKs create own backing indexes.
FKs not co-locate tables in the same storage layer.
FKs permit circular references.

````sql
CREATE TABLE Singers (
  SingerId   INT64 NOT NULL,
  FirstName  STRING(100),
  LastName   STRING(100)
) PRIMARY KEY (SingerId);

CREATE TABLE Albums (
  SingerId     INT64 NOT NULL,
  AlbumId      INT64 NOT NULL,
  AlbumTitle   STRING(MAX),
) PRIMARY KEY (SingerId, AlbumId), INTERLEAVE IN PARENT Singers ON DELETE CASCADE;

-- ok
INSERT INTO Singers (SingerId, LastName) VALUES (1, "Lamar");
INSERT INTO Albums (SingerId, AlbumId, AlbumTitle) VALUES (1, 1, "1st");

-- HTTPError 404: {"code":5,"message":"Insert failed because key was not found in parent table:  Parent Table: Singers  Child Table: Albums  Key: {Int64(0)}"}
INSERT INTO Albums (SingerId, AlbumId, AlbumTitle) VALUES (3, 2, "2nd");

````

````sql
CREATE TABLE Customers (
  ID INT64 NOT NULL,
  FirstName STRING(100),
  LastName STRING(100)
) PRIMARY KEY (ID);

CREATE TABLE Products (
  ID INT64 NOT NULL,
  Name STRING(100),
  Price FLOAT64
) PRIMARY KEY (ID);

CREATE TABLE Orders (
  OrderID INT64 NOT NULL,
  CustomerID INT64 NOT NULL,
  ProductID INT64 NOT NULL,
  Quantity INT64 NOT NULL,
  CONSTRAINT FK_CustomerID FOREIGN KEY (CustomerID) REFERENCES Customers (ID),
  CONSTRAINT FK_ProductID FOREIGN KEY (ProductID) REFERENCES Products (ID)
) PRIMARY KEY (OrderID);

CREATE TABLE OrderItems (
  OrderID INT64 NOT NULL,
  OrderItemID INT64 NOT NULL,
  ProductID INT64,
  Quantity INT64
) PRIMARY KEY (OrderID, OrderItemID), INTERLEAVE IN PARENT Orders ON DELETE CASCADE;

INSERT INTO Customers (ID, FirstName, LastName) VALUES (1, "Bob", "None");
INSERT INTO Customers (ID, FirstName, LastName) VALUES (2, "Tom", "Best");

INSERT INTO Products (ID, Name, Price) VALUES (1, "Apple", 1.13);
INSERT INTO Products (ID, Name, Price) VALUES (2, "Orange", 2.14);

INSERT INTO Orders (OrderID, CustomerID, ProductID, Quantity) VALUES (1, 1, 1, 10);
INSERT INTO Orders (OrderID, CustomerID, ProductID, Quantity) VALUES (1, 3, 1, 1); -- error: no customer 3
INSERT INTO Orders (OrderID, CustomerID, ProductID, Quantity) VALUES (1, 2, 3, 1); -- error: no product 3

INSERT INTO OrderItems (OrderID, OrderItemID, ProductID, Quantity) VALUES (1, 1, 1, 10);
INSERT INTO OrderItems (OrderID, OrderItemID, ProductID, Quantity) VALUES (3, 1, 1, 1); -- error: no order 3

SELECT * FROM Customers;
SELECT * FROM Products;
SELECT * FROM Orders;
SELECT * FROM OrderItems;

DELETE FROM Customers WHERE 1=1; -- error: constraint violation
DELETE FROM Orders WHERE 1=1; -- will delete OrderItems as well

````
