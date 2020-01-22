Best Practices
-

* Disable _source.

* The default dynamic string mappings will index string fields both as text and keyword,
so don't use it.

* Put fields in the same order in documents
(more likely to find if fields always occur in the same order).

* Search as few fields as possible.

* Avoid scripts.
