Search
-

#### Best Practices

* Turn negatives into positives:
`‘NOT cuisine:undefined’` -> use two fields, `cuisine`, and `cuisine_known`
and use `‘cuisine_known:yes’`.

* Turn disjunctions into conjunctions:
`‘cuisine:Japanese OR cuisine:Korean’ -> ‘cuisine:Asian’`.

* Avoid large documents.

* Don’t return large result sets.

* Narrow the range before sorting!!! (modre additional filters - less data to sort).

* Do not score matches unless you sort by score.

#### Score

Calculate score manually:
````sh
// formula = score_baseline + (delta_score * boosting_rate) + (score * exp(-decay * days_since_last_scored))
1 + (3 * 1) + (7 * exp(-5.5 * 1)) = 1 + 3 + 7 * 0.0040867
````
