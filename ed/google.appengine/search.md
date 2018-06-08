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

````
index.Search(ctx, "Product = piano AND Price < 5000", nil)
````

https://cloud.google.com/appengine/docs/standard/go/search/#Go_Special_treatment_of_string_and_date_fields

https://cloud.google.com/appengine/docs/standard/go/search/reference
