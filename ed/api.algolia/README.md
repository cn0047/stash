Algolia
-

[doc](https://www.algolia.com/doc/api-client/go/getting-started/)
[price](https://www.algolia.com/pricing)

Algolia - external 3d party search.

Features: typo-tolerance, highlighting and snippeting, faceting,
synonyms, query expansion, advanced language processing,
geo-awareness, distinct, personalization.

````
Relational DB ⇒ Database ⇒ Table ⇒ Row             ⇒ Column
Algolia       ⇒          ⇒ Index ⇒ Record (object) ⇒
````

Stop words are very common words in a given language.
In English, these would be words like: the, and, at, with and as.

Query Rules allows you to target specific search terms
and alter the way the Algolia engine would normally treat those terms.

`objectID` - record id.

````js
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

````js
index.query({
  query: 'query terms',
  clickAnalytics: true
});
````

## Api Keys

* Admin API Key
* Search-only API Key
* Monitoring API Key

[new key](https://monosnap.com/file/ZDc7JGQwpuOTiKh8HB0LeDUb6WPlqU)

## Search

````js
index.search({
  query: 'query',
  page: 2,
  hitsPerPage: 5
});

index.setSettings({
  paginationLimitedTo: 5000
});

algolia.search("phone", { optionalFilters: ["brand:samsung"] });
````

## Prefix search

````js
index.setSettings({
  queryType: 'prefixLast'
  //queryType: 'prefixAll'
  //queryType: 'prefixNone'
});
````

## Geo search

````js
index.search({
  query: 'query',
  aroundRadius: 1000, // 1km Radius
  minimumAroundRadius: 1000 // 1km Radius
}).then(res => { console.log(res); });
````

## Filters

Filters: facet, numeric, tagged.

`Facets` are used to create categories on a select group of attributes.
For example, on an index of books, useful facets might be: author and genre.

````js
index.search({
  filters: '(productType:book OR productType:dvd) AND genre:"sci-fi" AND price < 10'
});

index.setSettings({
  'attributesForFaceting': ['productType', 'genre', 'price']
});
index.search({
  facetFilters: [["category:Book", "category:Ebook"], "author:JK Rowling"]
});
index.search({
  facets: ['author', 'categories', 'publisher']
});
````

## Scoring

````js
index.search({
  filters: "(company:Google<score=3> OR company:Amazon<score=2> OR company:Facebook<score=1>)",
});
````

## Ranking

````js
index.setSettings({
  customRanking: [
    'desc(popularity)',
    'asc(price)'
  ]
});
````
