Theory
-

#### Programming:

Imperative programming is a programming paradigm that uses statements that change a program's state.
Imperative program consists of commands for the computer to perform.

Declarative programming focuses on what the program should accomplish
without specifying how the program should achieve the result.

#### Relational database:

A relation is in 1NF if the domain of each attribute contains only atomic values.
1NF enforces these criteria:
* Eliminate repeating groups in individual tables.
* Create a separate table for each set of related data.
* Identify each set of related data with a primary key.
(most basic: each cell in a table must contain only one piece of information,
and there can be no duplicate rows).

Table is in 2NF if it is in 1NF
and every non-prime attribute of the table (in main table) is dependent on the whole of every candidate key.

Table is in 3NF if it is in 2NF
and each attribute must represent a fact about the key, the whole key, and nothing but the key
(every non-prime attribute of table is non-transitively dependent on every key of table).
