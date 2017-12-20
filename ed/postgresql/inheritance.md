Inheritance
-

````sql
CREATE TABLE cities (
  name text,
  population real
); 
CREATE TABLE capitals (
  country char(2) 
) INHERITS (cities);
INSERT INTO cities VALUES ('NYC', 8538000);
INSERT INTO cities VALUES ('London', 8788000);
INSERT INTO capitals VALUES ('Tokyo', 9273000, 'JP');

SELECT * FROM cities;
 NYC    |  8.538e+06
 London |  8.788e+06

SELECT * FROM capitals;
 Tokyo |  9.273e+06 | JP
````

Although inheritance is frequently useful, it has not been integrated with
unique constraints or foreign keys, which limits its usefulness.
