Search
-

#### Best Practices

Turn negatives into positives:
`‘NOT cuisine:undefined’` -> use two fields, `cuisine`, and `cuisine_known`
and use `‘cuisine_known:yes’`.

Turn disjunctions into conjunctions:
`‘cuisine:Japanese OR cuisine:Korean’ -> ‘cuisine:Asian’`.

Narrow the range before sorting!!! (modre additional filters - less data to sort).

Do not score matches unless you sort by score.
