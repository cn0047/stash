Elasticsearch
-
2.3.3
2.2
1.6.0

In Elasticsearch, a document belongs to a type, and those types live inside an index.
You can draw some (rough) parallels to a traditional relational database:
````
Relational DB ⇒ Databases ⇒ Tables ⇒ Rows      ⇒ Columns
Elasticsearch ⇒ Indices   ⇒ Types  ⇒ Documents ⇒ Fields
````
We could use the DELETE verb to delete the document.
And the HEAD verb to check whether the document exists.
To replace an existing document - just PUT it again.

Use `delete-by-query` plugin to delete all documents matching a specific query.

`_score` field in the search results - the higher the score,
the more relevant the document is, the lower the score, the less relevant the document is.
`filter` clauses which allow to use a query
to restrict the documents that will be matched by other clauses, without changing how scores are computed.

[Meta-Fields](https://www.elastic.co/guide/en/elasticsearch/reference/current/mapping-fields.html).

[Groovy](https://www.elastic.co/guide/en/elasticsearch/reference/current/modules-scripting.html#_document_fields).

[Glossary](https://www.elastic.co/guide/en/elasticsearch/reference/current/glossary.html).

A shard is a single Lucene instance.
An index is a logical namespace which points to primary and replica shards.

https://www.elastic.co/guide/en/elasticsearch/reference/2.2/release-notes-2.2.0.html

Pay attention:

* Messure marvel & JVM.
* Not use open JDK but latest Oracle JDK.

In mapping `_source: {"enabled": "false"}` - don't save _source only indexed data.
It can save space on prod.

#### Versioning

Internally, Elasticsearch has marked the old document as deleted and added an entirely new document.
The old version of the document doesn’t disappear immediately,
although you won’t be able to access it.
Elasticsearch cleans up deleted documents in the background as you continue to index more data.

#### Analyzers

Analyzers are used - when we index a document.

* `Standard` - divides text into terms on word boundaries.
* `Simple` - divides whenever it encounters a character which is not a letter.
* `Whitespace` - divides whenever it encounters any whitespace character.
* `Stop` - like simpleanalyzer, but also supports removal of stop words.
* `Keyword` - is a “noop” analyzer that accepts whatever text it is given and outputs the exact same text as a single term.
* `Pattern` - uses a regular expression to split the text into terms.
* `Language` - language-specific analyzers like english or french.
* `Fingerprint` - is a specialist analyzer which creates a fingerprint which can be used for duplicate detection.
