Full Text Search
-

[see](https://neo4j.com/docs/cypher-manual/current/administration/indexes-for-full-text-search)

Neo4j is using lucene engine.

## Administration

````js
// create index
CALL db.index.fulltext.createNodeIndex('idx_name', ['Person', 'Organization'], ['name']);
CALL db.index.fulltext.createNodeIndex('idx_about', ['Person'], ['about']);
CALL db.index.fulltext.createNodeIndex('idx_all_text', ['Person'], ['name', 'code', 'about']);
CALL db.index.fulltext.createNodeIndex('idx_code', ['Person'], ['code'], {analyzer: 'url_or_email', eventually_consistent: 'true'});

// get indexes
CALL db.index.fulltext.listAvailableAnalyzers;

// delete index
CALL db.index.fulltext.drop('idx_code');
````

## Search

````js
CALL db.index.fulltext.queryNodes('idx_name', 'bond') YIELD node, score RETURN node, score;
CALL db.index.fulltext.queryNodes('idx_name', '*ond') YIELD node, score RETURN node, score;
CALL db.index.fulltext.queryNodes('idx_name', 'bon*') YIELD node, score RETURN node, score;
CALL db.index.fulltext.queryNodes('idx_name', 'b*nd') YIELD node, score RETURN node, score;

CALL db.index.fulltext.queryNodes('idx_about', 'vod*') YIELD node, score RETURN node, score;

// vodka OR martini
CALL db.index.fulltext.queryNodes('idx_about', 'vodka martini') YIELD node, score RETURN node, score;
// phrase search
CALL db.index.fulltext.queryNodes('idx_about', '"vodka martini"') YIELD node, score RETURN node, score;

CALL db.index.fulltext.queryNodes('idx_all_text', '007') YIELD node, score RETURN node, score;
CALL db.index.fulltext.queryNodes('idx_all_text', '007 OR about:vod*') YIELD node, score RETURN node, score;
CALL db.index.fulltext.queryNodes('idx_all_text', 'about:(007 vod*)') YIELD node, score RETURN node, score;

// boosting
CALL db.index.fulltext.queryNodes('idx_all_text', 'about:vodka^3 about:martini') YIELD node, score RETURN node, score;

// full-text + relationship
CALL db.index.fulltext.queryNodes('idx_all_text', '007 OR about:vod*') YIELD node, score
WHERE EXISTS ( (node)-[:WORKS_AT]->(:Organization {name:'MI6'}) )
RETURN node, score;
````
