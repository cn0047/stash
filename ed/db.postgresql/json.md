JSON
-

Operators:

* `->` - for int/text values.
* `->>` - for int/text values.
* `#>` - path.
* `#>>` - path as text.
* `@>`, `<@` - for JSONB.
* `||` - concatenate two JSONB values.

Functions:

* to_json(string)
* to_jsonb
* array_to_json
* row_to_json(fieldFromDB)
* json_build_array
* jsonb_build_array
* json_build_objectB
* jsonb_build_objectB
* json_object
* jsonb_object
* json_object('{a, 1, b, "def", c, 3.5}')
* jsonb_object
* json_agg(fieldFromDB)

````sql
-- sum
select sum((data->>'productsCount')::integer)
````

````sql
CREATE TABLE books (id integer, data json);

INSERT INTO books VALUES
     (1, '{ "name": "B-First",  "tags": ["sql", "one"], "author": { "first_name": "Bob", "last_name": "White" } }')
    ,(2, '{ "name": "B-Second", "tags": ["sql"],        "author": { "first_name": "Charles", "last_name": "Xavier" } }')
    ,(3, '{ "name": "B-Third",  "tags": ["sql", "fun"], "author": { "first_name": "Jim", "last_name": "Brown" } }')
;
````

````sql
-- SELECT:

SELECT * FROM books;
 id |                                            data
----+---------------------------------------------------------------------------------------------
  1 | { "name": "Book the First", "author": { "first_name": "Bob", "last_name": "White" } }
  2 | { "name": "Book the Second", "author": { "first_name": "Charles", "last_name": "Xavier" } }
  3 | { "name": "Book the Third", "author": { "first_name": "Jim", "last_name": "Brown" } }

SELECT data->'name' FROM books;
  ?column?
------------
 "B-First"
 "B-Second"
 "B-Third"

SELECT * FROM books WHERE data->>'name' = 'B-First';
 id |                                                  data
----+---------------------------------------------------------------------------------------------------------
  1 | { "name": "B-First",  "tags": ["sql", "one"], "author": { "first_name": "Bob", "last_name": "White" } }

SELECT data::json#>'{author, first_name}' FROM books WHERE data->>'name' = 'B-First';
 ?column?
----------
 "Bob"

SELECT data::jsonb FROM books WHERE data::jsonb @> '{"name": "B-Third"}';
----------------------------------------------------------------------------------------------------
 {"name": "B-Third", "tags": ["sql", "fun"], "author": {"last_name": "Brown", "first_name": "Jim"}}

SELECT data->'name'FROM books WHERE (data->'tags')::jsonb ? 'fun';
 ?column?
-----------
 "B-Third"
````

````sql
# INDEX:
# 
CREATE INDEX ON books ((data->>'author'));
````
