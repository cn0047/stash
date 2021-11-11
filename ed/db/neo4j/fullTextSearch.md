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
CALL db.indexes();

// get analyzers
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

// wildcard
CALL db.index.fulltext.queryNodes('idx_about', 'vod*') YIELD node, score RETURN node, score;
// single character wildcard
CALL db.index.fulltext.queryNodes('idx_about', 'vod?a') YIELD node, score RETURN node, score;

// simple sort
CALL db.index.fulltext.queryNodes('idx_about', 'vod*') YIELD node, score
RETURN node, node.code, score ORDER BY node.code ASC;

// vodka OR martini
CALL db.index.fulltext.queryNodes('idx_about', 'vodka martini') YIELD node, score RETURN node, score;
// phrase search
CALL db.index.fulltext.queryNodes('idx_about', '"vodka martini"') YIELD node, score RETURN node, score;

// fuzziness
CALL db.index.fulltext.queryNodes('idx_about', 'vod~') YIELD node, score RETURN node, score;

CALL db.index.fulltext.queryNodes('idx_all_text', '007') YIELD node, score RETURN node, score;
CALL db.index.fulltext.queryNodes('idx_all_text', '007 OR about:vod*') YIELD node, score RETURN node, score;
CALL db.index.fulltext.queryNodes('idx_all_text', 'about:(007 vod*)') YIELD node, score RETURN node, score;

// range - string_rating:[50 TO 99]
// CALL db.index.fulltext.queryNodes('idx_all_text', 'code:[005 TO 009]') YIELD node, score RETURN node, score;

// boosting
CALL db.index.fulltext.queryNodes('idx_all_text', 'about:vodka^3 about:martini') YIELD node, score RETURN node, score;

// full-text + relationship
CALL db.index.fulltext.queryNodes('idx_all_text', '007 OR about:vod*') YIELD node, score
WHERE EXISTS ( (node)-[:WORKS_AT]->(:Organization {name:'MI6'}) )
RETURN node, score;

// decay
WITH apoc.text.join(
  [
    x IN range(10, 30, 10)
    | "createdAt:" + "{"+ toString(timestamp() - (x * 3600 * 1000)) + " TO " + toString(timestamp() - ((x-10) * 3600 * 1000)) +"}^" + toInteger(((100-x)/10)+1)
  ],
  " "
) AS decay RETURN decay;
// Result:  "createdAt:{1629261161806 TO 1629297161806}^10 createdAt:{1629225161806 TO 1629261161806}^9 createdAt:{1629189161806 TO 1629225161806}^8"

// query with decay
WITH apoc.text.join(
  [
    x IN range(10, 30, 10)
    | "createdAt:" + "{"+ toString(timestamp() - (x * 3600 * 1000)) + " TO " + toString(timestamp() - ((x-10) * 3600 * 1000)) +"}^" + toInteger(((100-x)/10)+1)
  ],
  " "
) AS decay
CALL db.index.fulltext.queryNodes('idx_all_text', 'about:gadget~ ' + decay) YIELD node, score RETURN node, score;
// without decay, for comparison
CALL db.index.fulltext.queryNodes('idx_all_text', 'about:gadget~ ') YIELD node, score RETURN node, score;

````

## Custom analyzer

[example](https://graphaware.com/neo4j/2019/09/06/custom-fulltext-analyzer.html)
[example](https://github.com/graphaware/custom-fulltext-analyzer-blog)
[maven for analyzer](https://neo4j.com/docs/java-reference/current/extending-neo4j/procedures-and-functions/procedures-setup/#extending-neo4j-procedures-setup)
