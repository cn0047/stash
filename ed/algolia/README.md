Algolia
-

[doc](https://www.algolia.com/doc/api-client/go/getting-started/)
[price](https://www.algolia.com/pricing)

Algolia - external 3d party search.

Features: typo-tolerance, highlighting and snippeting, faceting,
synonyms, query expansion, advanced language processing,
geo-awareness, distinct, personalization.

````
Relational DB ⇒ Databases ⇒ Table ⇒ Rows            ⇒ Columns
Algolia       ⇒           ⇒ Index ⇒ Record (object) ⇒
````

`Facets` are used to create categories on a select group of attributes.
For example, on an index of books, useful facets might be author and genre.

Stop words are very common words in a given language.
In English, these would be words like the, and, at, with and as.

Query Rules allows you to target specific search terms
and alter the way the Algolia engine would normally treat those terms.

`objectID` - record id.

````
const objects = [{country: 'France', objectID: 42}];
const ids = [42];

index.addObjects(objects, function(err, content) {});
index.saveObjects(objects, function(err, content) {});
index.partialUpdateObject(objects, function(err, content) {});

index.deleteObjects(ids, function(err, content) {});
index.deleteBy(
  {filters: 'year < 2017'},
  function(err) {
    if (!err) { console.log('success'); }
  }
);
````

## Api Keys

* Admin API Key
* Search-only API Key
* Monitoring API Key

[new key](https://monosnap.com/file/ZDc7JGQwpuOTiKh8HB0LeDUb6WPlqU)
