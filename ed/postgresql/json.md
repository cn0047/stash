JSON
-

Operators:

* `->` - int/text
* `->>` - int/text
* `#>` - path
* `#>>` - path as text
* `@>` -
* `<@` -
* `?` -
* `?|` -
* `?&` -
* `||` -
* `-` -
* `#-` -

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
* and else ...

````
CREATE TABLE books (id integer, data json);

INSERT INTO books VALUES (1, '{ "name": "B-First", "tags": ["sql", "one"], "author": { "first_name": "Bob", "last_name": "White" } }');
INSERT INTO books VALUES (2, '{ "name": "B-Second", "tags": ["sql"], "author": { "first_name": "Charles", "last_name": "Xavier" } }');
INSERT INTO books VALUES (3, '{ "name": "B-Third", "tags": ["sql", "fun"], "author": { "first_name": "Jim", "last_name": "Brown" } }');

SELECT * FROM books;
 id |                                            data
----+---------------------------------------------------------------------------------------------
  1 | { "name": "Book the First", "author": { "first_name": "Bob", "last_name": "White" } }
  2 | { "name": "Book the Second", "author": { "first_name": "Charles", "last_name": "Xavier" } }
  3 | { "name": "Book the Third", "author": { "first_name": "Jim", "last_name": "Brown" } }

````

````
SELECT data->'name' FROM books;

SELECT * FROM books WHERE data->>'name' = 'B-First';



-- SELECT * FROM books WHERE jdoc -> 'tags' ? 'qui';
-- SELECT * FROM books WHERE data->'tags' @> '["sql"]';
-- SELECT * FROM books WHERE data->'tags' @> '{"name": "Book the First"}';
````
