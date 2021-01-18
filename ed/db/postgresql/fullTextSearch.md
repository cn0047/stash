Full Text Search
-

````sql
@@                   -- match operator
tsvector             -- document
tsvector || tsvector -- concatenation
strip(tsvector)      -- returns tsvector without weight information
tsquery && tsquery   -- and
tsquery || tsquery   -- or
!! tsquery           -- not
numnode(tsquery)     -- number of nodes in a tsquery without stop words

SELECT 'fat cat ate a fat rat'::tsvector @@ 'cat & rat'::tsquery;
SELECT to_tsvector('fat cats ate fat rats') @@ to_tsquery('fat & rat');

-- debug:
SELECT * FROM ts_debug('english', 'fat cat ate fat rats');
SELECT * FROM ts_parse('default', '123 - a number');
````

`to_tsvector` parses a textual document into tokens,
reduces the tokens to lexemes, and returns a tsvector
which lists the lexemes together with their positions in the document.

Because `to_tsvector(NULL)` will return NULL, it is recommended to use `coalesce`
whenever a field might be null.

The function `setweight` can be used to label the entries of a tsvector with a given weight,
where a weight is one of the letters A, B, C, or D
(like title versus body) this can be used for ranking.

Ranking functions: `ts_rank, ts_rank_cd`.

## Example

````sql
DROP TABLE IF EXISTS articles;
CREATE TABLE articles (
  id SERIAL NOT NULL PRIMARY KEY,
  title CHARACTER VARYING(200) DEFAULT '',
  body TEXT
);
INSERT INTO articles (title,body) VALUES
('Use Postgres DataBase','Use Postgres DataBase for full-text search ...'),
('Use Postgres','This is postgres ...'),
('Use Postgres','Postgres DataBase - tutorial about probably best database ...'),
('Use Postgres','In this postgres tutorial we will see IN clause ...'),
('MySQL Tutorial','DBMS stands for DataBase ...'),
('How To Use MySQL Well','After you went through a ...'),
('Optimizing MySQL','In this tutorial we will show ...'),
('1001 MySQL Tricks','1. Never run mysqld as root. 2. ...'),
('MySQL vs. YourSQL','In the following database comparison ...'),
('Optimizing Postgres','In this postgres tutorial we will show ...'),
('MySQL Security','When configured properly, MySQL ...');

-- search in regular table
SELECT * FROM articles
WHERE to_tsvector(body) @@ to_tsquery('database');
-- or
SELECT * FROM articles
WHERE to_tsvector('english', body) @@ to_tsquery('english', 'database');
-- or
SELECT * FROM articles WHERE
  to_tsvector('english', coalesce(title,'') || ' ' || coalesce(body,''))
  @@ to_tsquery('english', 'MySQL');

-- phrase search "postgres tutorial"
SELECT * FROM articles
WHERE to_tsvector(body) @@ to_tsquery('postgres <-> tutorial');
````
`<->` - distance between words.
`'jump <-> quick'` will match: `wizards jump quickly.`
`'sphinx <2> quartz'` will match: `sphinx of quartz.`
`'sphinx <3> quartz'` will match: `Sphinx of black quartz.`

Wildcards:
````sql
SELECT * FROM articles
WHERE to_tsvector(body) @@ to_tsquery('database');
-- result:
 id |       title       |                             body
----+-------------------+---------------------------------------------------------------
  1 | MySQL Tutorial    | DBMS stands for DataBase ...
  5 | MySQL vs. YourSQL | In the following database comparison ...
  9 | Use Postgres      | Postgres DataBase - tutorial about probably best database ...
(3 rows)

SELECT * FROM articles
WHERE to_tsvector(body) @@ to_tsquery('datab:*');
-- result:
 id |       title       |                             body
----+-------------------+---------------------------------------------------------------
  1 | MySQL Tutorial    | DBMS stands for DataBase ...
  5 | MySQL vs. YourSQL | In the following database comparison ...
  9 | Use Postgres      | Postgres DataBase - tutorial about probably best database ...
(3 rows)

SELECT * FROM articles
WHERE to_tsvector(body) @@ to_tsquery('dat*base');
-- result:
 id | title | body
----+-------+------
(0 rows)

SELECT * FROM articles
WHERE to_tsvector(body) @@ to_tsquery(':*base');
-- result:
ERROR:  syntax error in tsquery: ":*base"
````

## Index

GIN indexes are the preferred text search index type.
⚠️ GiST index is lossy, might produce false matches.

````sql
CREATE INDEX articles_i ON articles USING GIN (to_tsvector('english', body));

CREATE INDEX articles_i2 ON articles USING GIN (to_tsvector('english', title || ' ' || body));
````
Index `articles_i` will use queries with:
`to_tsvector('english', body)`
`to_tsvector(body)` wont use this index because 'english' not provide.

## Ranking

Ranking can be expensive since it requires consulting the tsvector of each matching document...

````sql
SELECT body, ts_rank_cd(to_tsvector(body), query) AS rank
FROM articles, to_tsquery('in|(we & see)') query
WHERE query @@ to_tsvector(body)
ORDER BY rank DESC;

SELECT body, ts_rank_cd(to_tsvector(body), query) AS rank
FROM articles, to_tsquery('database') query
WHERE query @@ to_tsvector(body)
ORDER BY rank DESC;
````

## Statistics

````sql
SELECT * FROM ts_stat('SELECT to_tsvector(body) FROM articles');
````
Where:
ndoc — number of documents the word occurred in.
nentry — total number of occurrences of the word.

## Limitations

* The length of each lexeme must be less than 2K bytes.
* The length of a tsvector (lexemes + positions) must be less than 1 megabyte
* The number of lexemes must be less than 2^64
* Position values in tsvector must be greater than 0 and no more than 16,383
* No more than 256 positions per lexeme
* The number of nodes (lexemes + operators) in a tsquery must be less than 32,768

## Speed UP

````sql
ALTER TABLE articles ADD COLUMN tsv tsvector;
CREATE INDEX tsv_i ON articles USING gin(tsv);
UPDATE articles SET tsv =
  setweight(to_tsvector(coalesce(title, '')), 'A') || setweight(to_tsvector(coalesce(body, '')), 'B');

-- check 1:

EXPLAIN SELECT * FROM articles
WHERE to_tsvector(body) @@ to_tsquery('database');
-- cost=0.00..6.12

EXPLAIN SELECT * FROM articles
WHERE tsv @@ to_tsquery('database');
-- cost=0.00..3.62

-- check 2:

EXPLAIN SELECT id, title, body, ts_rank_cd(to_tsvector(title || ' ' || body), query) AS rank
FROM articles, to_tsquery('database') query
WHERE query @@ to_tsvector(title || ' ' || body)
ORDER BY rank DESC;
-- cost=4.30..4.31

EXPLAIN SELECT id, title, body, ts_rank_cd(tsv, query) AS rank
FROM articles, to_tsquery('database') query
WHERE query @@ tsv
ORDER BY rank DESC;
-- cost=1.50..1.50
````

# Trigram (Trigraph) search

Full text search is good for finding words, not substrings.
For substrings you have to use module `pg_trgm`.

Also it's good for misspelled search.

````sql
-- enable module:
CREATE EXTENSION pg_trgm;

SELECT word_similarity('word', 'two words');

ALTER TABLE articles ADD COLUMN trg TEXT;
CREATE INDEX trgm_i ON articles USING GiST (trg gist_trgm_ops);
-- ‼️ This can be implemented quite efficiently by GiST indexes, but not by GIN indexes.
UPDATE articles SET trg = body;
````

````sql
-- distance:
SELECT body, body <-> 'atabas' AS dist
FROM articles;

-- wildcards:

SELECT id, trg, word_similarity('database', trg) AS sml
FROM articles
WHERE 'database' <% trg
;
-- result:
 id |                              trg                              | sml
----+---------------------------------------------------------------+-----
  1 | Use Postgres DataBase for full-text search ...                |   1
  3 | Postgres DataBase - tutorial about probably best database ... |   1
  5 | DBMS stands for DataBase ...                                  |   1
  9 | In the following database comparison ...                      |   1
(4 rows)

SELECT id, trg, word_similarity('datab', trg) AS sml
FROM articles
WHERE 'datab' <% trg
;
-- result:
 id |                              trg                              |   sml
----+---------------------------------------------------------------+----------
  1 | Use Postgres DataBase for full-text search ...                | 0.833333
  3 | Postgres DataBase - tutorial about probably best database ... | 0.833333
  5 | DBMS stands for DataBase ...                                  | 0.833333
  9 | In the following database comparison ...                      | 0.833333
(4 rows)

SELECT id, trg, word_similarity('atabase', trg) AS sml
FROM articles
WHERE 'atabase' <% trg
;
-- result:
 id |                              trg                              | sml
----+---------------------------------------------------------------+------
  1 | Use Postgres DataBase for full-text search ...                | 0.75
  3 | Postgres DataBase - tutorial about probably best database ... | 0.75
  5 | DBMS stands for DataBase ...                                  | 0.75
  9 | In the following database comparison ...                      | 0.75
(4 rows)
````
