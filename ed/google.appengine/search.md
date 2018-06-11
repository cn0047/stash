Search
-

Every document in an index must have a unique document identifier, or `docID` (no longer than 500 characters).
You cannot include the docID in a search.

Two or more fields can have the same name, but different types.

There are three kinds of fields that store character strings:
* Text Field: A string with maximum length `1024**2` characters.
* HTML Field: An HTML-formatted string with maximum length `1024**2` characters.
* Atom Field: A string with maximum length 500 characters.

There are also three field types that store non-textual data:

Number Field: A double precision floating point value between -2,147,483,647 and 2,147,483,647.
Time Field: A `time.Time` value (stored with millisecond precision).
Geopoint Field: A point on earth described by latitude and longitude coordinates.

The maximum length of a query string is 2000 characters.
⚠️ It's recommended to write field values in lower case, because:
`AND, OR, and NOT` are boolean operators and recognized by writing them in upper case.
<br>The boolean operator precedence, from highest to lowest, is: `NOT, OR, AND`.

match: `"rose bud"`.
phrase: `"\"rose bud\""`.

````
index.Search(ctx, "Product = piano AND Price < 5000", nil)

query = "pet = dog"
query = "author = \"Ray Bradbury\""
query = "color:red"
query = "NOT color:red"
query = "price < 500"
query = "birthday>=2011-05-10"
````

Stemming:

`~dog` matches `"dog" or "dogs."`
`~care` will match `"care" and "caring," but not "cares" or "cared."`


#### Atom Fields

Queries on atom fields are case insensitive.
The only valid relational operator for atom fields is the equality operator. 
Stemming is not supported for atom fields.

`"weather=stormy"`, `"weather: stormy"` -  documents with a weather field that equals "stormy".
`"Title: \"Tom&Jerry\""`, `"Couple: \"Fred and Ethel\""`, `"Version = \"1HCP(21.3)\""` - contains whitespace.
`"Color = (red OR blue)"`, `"Color = (\"dark red\" OR \"bright blue\")"` - parentheses.

#### Text and HTML fields

The only valid relational operator for text and HTML fields is equality.
You can use the stemming operator.
You can also use the OR and AND operators.
When searching HTML fields, the text within HTML markup tags is ignored.
Queries on text and HTML fields are case insensitive.

`"Comment = great"`, `"Comment: great"` - documents with a comment field
that contains at least one occurrence of the word "great" in the Comment field.
`"Comment = (great big ball)"`, `"Comment = (great AND big AND ball)"` - two or more words in a field.
`"Comment = \"insanely great\""` - specific string of text.
`"pet = ~dog"` - stemming.
`"Color = (red OR blue)"` - list of alternatives.
`"weather = ((rain OR snow) AND cold)"` - more complex field value.
`"weather = \"rain OR shine\""` - weather field that contains the string "rain or shine", because OR in a quoted string.

#### Number fields

Can be written as an integer, a decimal, or an exponential.
Operators: `<, <=, >, >=`. No inequality `!=` operator.

#### Date fields

Operators: `<, <=, >, >=`. No inequality operator.

````
"start_date: 2012-05-20"
"end_date: 2013-5-1"
"birthday >= 2000-12-31"
"NOT birthday = 2000-12-25"
````

#### Facet

A facet is an attribute/value pair.

#### Sort

You can never sort more than 10,000 docs, the default is 1,000.

#### Best Practices

Use single quotes: `‘field:"some text" some-value’`.

Use atom fields for boolean data.

Turn negatives into positives:
`‘NOT cuisine:undefined’` -> use two fields, `cuisine`, and `cuisine_known`
and use `‘cuisine_known:yes’`.

Turn disjunctions into conjunctions:
`‘cuisine:Japanese OR cuisine:Korean’` -> `‘cuisine:Asian’`.

Narrow the range before sorting!!! (modre additional filters - less data to sort).

Do not score matches unless you sort by score.

#### Examples

Queries on multiple fields:

````
"product=piano manufacturer=steinway"
"product=piano AND manufacturer=steinway"

"product=piano AND NOT manufacturer=steinway"

"product=piano AND price<2000"
````

Mixing global and field searches:

````
"keyboard great price<5000"
"keyboard AND great AND price<5000"

"keyboard OR product=piano"
````

````go
t := index.Search(ctx, "Comment:truth", nil)
t := index.List(ctx, nil)
````
