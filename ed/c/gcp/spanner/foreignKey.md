Foreign key
-

Interleaved table - child table, albums is interleaved table for singers (no singer -> no album).
For interleaved table can be only 1 parent.
Parent table primary key must be part of child table primary key.

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
  Quantity INT64,
  FOREIGN KEY (OrderID) REFERENCES Orders (OrderID)
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
